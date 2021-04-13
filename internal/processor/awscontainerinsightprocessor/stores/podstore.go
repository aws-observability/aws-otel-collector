package stores

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"

	aws "github.com/aws-observability/aws-otel-collector/internal/aws"
	. "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	"go.opentelemetry.io/collector/consumer/pdata"

	"github.com/aws-observability/aws-otel-collector/internal/aws/k8scommon/k8sclient"
	"github.com/aws-observability/aws-otel-collector/internal/aws/k8scommon/kubeletutil"
	corev1 "k8s.io/api/core/v1"
)

const (
	refreshInterval    = 30 * time.Second
	MeasurementsExpiry = 10 * time.Minute
	PodsExpiry         = 2 * time.Minute
	memoryKey          = "memory"
	cpuKey             = "cpu"
	splitRegexStr      = "\\.|-"
	kubeProxy          = "kube-proxy"
)

var (
	re = regexp.MustCompile(splitRegexStr)
)

type cachedEntry struct {
	pod      corev1.Pod
	creation time.Time
}

type Owner struct {
	OwnerKind string `json:"owner_kind"`
	OwnerName string `json:"owner_name"`
}

type prevPodMeasurement struct {
	containersRestarts int
}

type prevContainerMeasurement struct {
	restarts int
}

type mapWithExpiry struct {
	*aws.MapWithExpiry
}

func (m *mapWithExpiry) Get(key string) (interface{}, bool) {
	if val, ok := m.MapWithExpiry.Get(aws.NewKey(key, nil)); ok {
		return val.RawValue, ok
	}

	return nil, false
}

func (m *mapWithExpiry) Set(key string, content interface{}) {
	val := aws.MetricValue{
		RawValue:  content,
		Timestamp: time.Now(),
	}
	m.MapWithExpiry.Set(aws.NewKey(key, nil), val)
}

func newMapWithExpiry(ttl time.Duration) *mapWithExpiry {
	return &mapWithExpiry{
		MapWithExpiry: aws.NewMapWithExpiry(ttl),
	}
}

type PodStore struct {
	cache            *mapWithExpiry
	prevMeasurements map[string]*mapWithExpiry //preMeasurements per each Type (Pod, Container, etc)
	kubeClient       *kubeletutil.KubeClient
	lastRefreshed    time.Time
	nodeInfo         *nodeInfo
	prefFullPodName  bool
	sync.Mutex
}

func NewPodStore(hostIP string, prefFullPodName bool) *PodStore {
	podStore := &PodStore{
		cache:            newMapWithExpiry(PodsExpiry),
		prevMeasurements: make(map[string]*mapWithExpiry),
		kubeClient:       &kubeletutil.KubeClient{Port: KubeSecurePort, BearerToken: BearerToken, KubeIP: hostIP},
		nodeInfo:         newNodeInfo(),
		prefFullPodName:  prefFullPodName,
	}

	// Try to detect kubelet permission issue here
	if _, err := podStore.kubeClient.ListPods(); err != nil {
		panic(fmt.Sprintf("Cannot get pod from kubelet, err: %v", err))
	}

	return podStore
}

func (p *PodStore) getPrevMeasurement(metricType, metricKey string) (interface{}, bool) {
	prevMeasurement, ok := p.prevMeasurements[metricType]
	if !ok {
		return nil, false
	}

	content, ok := prevMeasurement.Get(metricKey)

	if !ok {
		return nil, false
	}

	return content, true
}

func (p *PodStore) setPrevMeasurement(metricType, metricKey string, content interface{}) {
	prevMeasurement, ok := p.prevMeasurements[metricType]
	if !ok {
		prevMeasurement = newMapWithExpiry(MeasurementsExpiry)
		p.prevMeasurements[metricType] = prevMeasurement
	}
	prevMeasurement.Set(metricKey, content)
}

func (p *PodStore) RefreshTick() {
	now := time.Now()
	if now.Sub(p.lastRefreshed) >= refreshInterval {
		p.refresh(now)
		// call cleanup every refresh cycle
		p.cleanup(now)
		p.lastRefreshed = now
	}
}

func (p *PodStore) Decorate(rm pdata.ResourceMetrics, kubernetesBlob map[string]interface{}) bool {
	attributes := rm.Resource().Attributes()

	if val, ok := attributes.Get(MetricType); ok && val.StringVal() == TypeNode {
		p.decorateNode(rm)
	} else if _, ok := attributes.Get(K8sPodNameKey); ok {
		podKey := createPodKeyFromMetric(attributes)
		if podKey == "" {
			log.Printf("E! podKey is unavailable when decorating pod.")
			return false
		}

		entry := p.getCachedEntry(podKey)
		if entry == nil {
			log.Printf("I! no pod is found for %s, refresh the cache now...", podKey)
			p.refresh(time.Now())
			entry = p.getCachedEntry(podKey)
		}

		// If pod is still not found, insert a placeholder to avoid too many refresh
		if entry == nil {
			log.Printf("W! no pod is found after reading through kubelet, add a placeholder for %s", podKey)
			p.setCachedEntry(podKey, &cachedEntry{creation: time.Now()})
			return false
		}

		// If the entry is not a placeholder, decorate the pod
		if entry.pod.Name != "" {
			p.decorateCpu(rm, attributes, &entry.pod)
			p.decorateMem(rm, attributes, &entry.pod)
			p.addStatus(rm, attributes, &entry.pod)
			addContainerCount(rm, attributes, &entry.pod)
			addContainerId(&entry.pod, attributes, rm, kubernetesBlob)
			p.addPodOwnersAndPodName(attributes, &entry.pod, kubernetesBlob)
			addLabels(&entry.pod, kubernetesBlob)
		} else {
			log.Printf("W! no pod information is found in podstore for pod %s", podKey)
			return false
		}
	}
	return true
}

func (p *PodStore) getCachedEntry(podKey string) *cachedEntry {
	p.Lock()
	defer p.Unlock()
	if content, ok := p.cache.Get(podKey); ok {
		return content.(*cachedEntry)
	}
	return nil
}

func (p *PodStore) setCachedEntry(podKey string, entry *cachedEntry) {
	p.Lock()
	defer p.Unlock()
	p.cache.Set(podKey, entry)
}

func (p *PodStore) setNodeStats(stats nodeStats) {
	p.Lock()
	defer p.Unlock()
	p.nodeInfo.nodeStats = stats
}

func (p *PodStore) getNodeStats() nodeStats {
	p.Lock()
	defer p.Unlock()
	return p.nodeInfo.nodeStats
}

func (p *PodStore) refresh(now time.Time) {
	podList, _ := p.kubeClient.ListPods()
	p.refreshInternal(now, podList)
}

func (p *PodStore) cleanup(now time.Time) {
	for _, prevMeasurement := range p.prevMeasurements {
		prevMeasurement.CleanUp(now)
	}

	p.Lock()
	defer p.Unlock()
	p.cache.CleanUp(now)
}

func (p *PodStore) refreshInternal(now time.Time, podList []corev1.Pod) {
	var podCount int
	var containerCount int
	var cpuRequest int64
	var memRequest int64

	for _, pod := range podList {
		podKey := createPodKeyFromMetaData(&pod)
		if podKey == "" {
			log.Printf("W! podKey is unavailable refresh pod store for pod %s", pod.Name)
			continue
		}
		tmpCpuReq, _ := getResourceSettingForPod(&pod, p.nodeInfo.getCPUCapacity(), cpuKey, getRequestForContainer)
		cpuRequest += tmpCpuReq
		tmpMemReq, _ := getResourceSettingForPod(&pod, p.nodeInfo.getMemCapacity(), memoryKey, getRequestForContainer)
		memRequest += tmpMemReq
		if pod.Status.Phase == corev1.PodRunning {
			podCount += 1
		}

		for _, containerStatus := range pod.Status.ContainerStatuses {
			if containerStatus.State.Running != nil {
				containerCount += 1
			}
		}

		p.setCachedEntry(podKey, &cachedEntry{
			pod:      pod,
			creation: now})
	}

	p.setNodeStats(nodeStats{podCnt: podCount, containerCnt: containerCount, memReq: memRequest, cpuReq: cpuRequest})
}

func (p *PodStore) decorateNode(metric pdata.ResourceMetrics) {
	nodeStats := p.getNodeStats()

	if containsMetric(MetricName(TypeNode, CpuTotal), metric) {
		if val := getMetricValue(MetricName(TypeNode, CpuLimit), metric); val != nil {
			p.nodeInfo.setCPUCapacity(val)
		}
		addMetric(MetricName(TypeNode, CpuRequest), nodeStats.cpuReq, UnitCount, metric)
		if p.nodeInfo.getCPUCapacity() != 0 {
			addMetric(MetricName(TypeNode, CpuReservedCapacity),
				float64(nodeStats.cpuReq)/float64(p.nodeInfo.getCPUCapacity())*100, UnitPercent, metric)
		}
	}

	if containsMetric(MetricName(TypeNode, MemWorkingset), metric) {
		if val := getMetricValue(MetricName(TypeNode, MemLimit), metric); val != nil {
			p.nodeInfo.setMemCapacity(val)
		}
		addMetric(MetricName(TypeNode, MemRequest), nodeStats.memReq, UnitBytes, metric)
		if p.nodeInfo.getMemCapacity() != 0 {
			addMetric(MetricName(TypeNode, MemReservedCapacity),
				float64(nodeStats.memReq)/float64(p.nodeInfo.getMemCapacity())*100, UnitPercent, metric)
		}
	}

	addMetric(MetricName(TypeNode, RunningPodCount), nodeStats.podCnt, UnitCount, metric)
	addMetric(MetricName(TypeNode, RunningContainerCount), nodeStats.containerCnt, UnitCount, metric)
}

func (p *PodStore) decorateCpu(metric pdata.ResourceMetrics, attributes pdata.AttributeMap, pod *corev1.Pod) {
	if GetStringValFromAttributes(MetricType, attributes) == TypePod {
		// add cpu limit and request for pod cpu
		if podCpuTotal := getMetricValue(MetricName(TypePod, CpuTotal), metric); podCpuTotal != nil {
			podCpuReq, _ := getResourceSettingForPod(pod, p.nodeInfo.getCPUCapacity(), cpuKey, getRequestForContainer)
			// set podReq to the sum of containerReq which has req
			if podCpuReq != 0 {
				addMetric(MetricName(TypePod, CpuRequest), podCpuReq, UnitCount, metric)
			}

			if p.nodeInfo.getCPUCapacity() != 0 {
				if podCpuReq != 0 {
					addMetric(MetricName(TypePod, CpuReservedCapacity), float64(podCpuReq)/float64(p.nodeInfo.getCPUCapacity())*100, UnitPercent, metric)
				}
			}

			podCpuLimit, ok := getResourceSettingForPod(pod, p.nodeInfo.getCPUCapacity(), cpuKey, getLimitForContainer)
			// only set podLimit when all the containers has limit
			if ok && podCpuLimit != 0 {
				addMetric(MetricName(TypePod, CpuLimit), podCpuLimit, UnitCount, metric)
				addMetric(MetricName(TypePod, CpuUtilizationOverPodLimit), podCpuTotal.(float64)/float64(podCpuLimit)*100, UnitPercent, metric)
			}
		}
	} else if GetStringValFromAttributes(MetricType, attributes) == TypeContainer {
		// add cpu limit and request for container
		if containsMetric(MetricName(TypeContainer, CpuTotal), metric) {
			if containerName := GetStringValFromAttributes(ContainerNamekey, attributes); containerName != "" {
				for _, containerSpec := range pod.Spec.Containers {
					if containerSpec.Name == containerName {
						if cpuLimit, ok := getLimitForContainer(cpuKey, containerSpec); ok {
							addMetric(MetricName(TypeContainer, CpuLimit), cpuLimit, UnitCount, metric)
						}
						if cpuReq, ok := getRequestForContainer(cpuKey, containerSpec); ok {
							addMetric(MetricName(TypeContainer, CpuRequest), cpuReq, UnitCount, metric)
						}
					}
				}
			}
		}
	}
}

func (p *PodStore) decorateMem(metric pdata.ResourceMetrics, attributes pdata.AttributeMap, pod *corev1.Pod) {
	if GetStringValFromAttributes(MetricType, attributes) == TypePod {
		if podMemWorkingset := getMetricValue(MetricName(TypePod, MemWorkingset), metric); podMemWorkingset != nil {
			// add mem limit and request for pod mem
			podMemReq, _ := getResourceSettingForPod(pod, p.nodeInfo.getMemCapacity(), memoryKey, getRequestForContainer)
			// set podReq to the sum of containerReq which has req
			if podMemReq != 0 {
				addMetric(MetricName(TypePod, MemRequest), podMemReq, UnitBytes, metric)
			}

			if p.nodeInfo.getMemCapacity() != 0 {
				if podMemReq != 0 {
					addMetric(MetricName(TypePod, MemReservedCapacity), float64(podMemReq)/float64(p.nodeInfo.getMemCapacity())*100, UnitBytes, metric)
				}
			}

			podMemLimit, ok := getResourceSettingForPod(pod, p.nodeInfo.getMemCapacity(), memoryKey, getLimitForContainer)
			// only set podLimit when all the containers has limit
			if ok && podMemLimit != 0 {
				addMetric(MetricName(TypePod, MemLimit), podMemLimit, UnitBytes, metric)
				addMetric(MetricName(TypePod, MemUtilizationOverPodLimit), podMemWorkingset.(float64)/float64(podMemLimit)*100, UnitPercent, metric)
			}
		}
	} else if GetStringValFromAttributes(MetricType, attributes) == TypeContainer {
		// add mem limit and request for container
		if containsMetric(MetricName(TypeContainer, MemWorkingset), metric) {
			if containerName := GetStringValFromAttributes(ContainerNamekey, attributes); containerName != "" {
				for _, containerSpec := range pod.Spec.Containers {
					if containerSpec.Name == containerName {
						if memLimit, ok := getLimitForContainer(memoryKey, containerSpec); ok {
							addMetric(MetricName(TypeContainer, MemLimit), memLimit, UnitBytes, metric)
						}
						if memReq, ok := getRequestForContainer(memoryKey, containerSpec); ok {
							addMetric(MetricName(TypeContainer, MemRequest), memReq, UnitBytes, metric)
						}
					}
				}
			}
		}
	}
}

func (p *PodStore) addStatus(metric pdata.ResourceMetrics, attributes pdata.AttributeMap, pod *corev1.Pod) {
	if GetStringValFromAttributes(MetricType, attributes) == TypePod {
		AddAttribute(PodStatus, string(pod.Status.Phase), attributes)
		var curContainerRestarts int
		for _, containerStatus := range pod.Status.ContainerStatuses {
			curContainerRestarts += int(containerStatus.RestartCount)
		}
		podKey := createPodKeyFromMetric(attributes)
		if podKey != "" {
			content, ok := p.getPrevMeasurement(TypePod, podKey)
			if ok {
				prevMeasurement := content.(prevPodMeasurement)
				result := 0
				if curContainerRestarts > prevMeasurement.containersRestarts {
					result = curContainerRestarts - prevMeasurement.containersRestarts
				}
				addMetric(MetricName(TypePod, ContainerRestartCount), result, UnitCount, metric)
			}
			p.setPrevMeasurement(TypePod, podKey, prevPodMeasurement{containersRestarts: curContainerRestarts})
		}
	} else if GetStringValFromAttributes(MetricType, attributes) == TypeContainer {
		if containerName := GetStringValFromAttributes(ContainerNamekey, attributes); containerName != "" {
			for _, containerStatus := range pod.Status.ContainerStatuses {
				if containerStatus.Name == containerName {
					if containerStatus.State.Running != nil {
						AddAttribute(ContainerStatus, "Running", attributes)
					} else if containerStatus.State.Waiting != nil {
						AddAttribute(ContainerStatus, "Waiting", attributes)
						if containerStatus.State.Waiting.Reason != "" {
							AddAttribute(ContainerStatusReason, containerStatus.State.Waiting.Reason, attributes)
						}
					} else if containerStatus.State.Terminated != nil {
						AddAttribute(ContainerStatus, "Terminated", attributes)
						if containerStatus.State.Terminated.Reason != "" {
							AddAttribute(ContainerStatusReason, containerStatus.State.Terminated.Reason, attributes)
						}
					}
					if containerStatus.LastTerminationState.Terminated != nil && containerStatus.LastTerminationState.Terminated.Reason != "" {
						AddAttribute(ContainerLastTerminationReason, containerStatus.LastTerminationState.Terminated.Reason, attributes)
					}
					containerKey := createContainerKeyFromMetric(attributes)
					if containerKey != "" {
						content, ok := p.getPrevMeasurement(TypeContainer, containerKey)
						if ok {
							prevMeasurement := content.(prevContainerMeasurement)
							result := 0
							if int(containerStatus.RestartCount) > prevMeasurement.restarts {
								result = int(containerStatus.RestartCount) - prevMeasurement.restarts
							}
							addMetric(ContainerRestartCount, result, UnitCount, metric)
						}
						p.setPrevMeasurement(TypeContainer, containerKey, prevContainerMeasurement{restarts: int(containerStatus.RestartCount)})
					}
				}
			}
		}
	}
}

// It could be used to get limit/request(depend on the passed-in fn) per pod
// return the sum of ResourceSetting and a bool which indicate whether all container set Resource
func getResourceSettingForPod(pod *corev1.Pod, bound int64, resource corev1.ResourceName, fn func(resource corev1.ResourceName, spec corev1.Container) (int64, bool)) (int64, bool) {
	var result int64
	allSet := true
	for _, containerSpec := range pod.Spec.Containers {
		val, ok := fn(resource, containerSpec)
		if ok {
			result += val
		} else {
			allSet = false
		}
	}
	if bound != 0 && result > bound {
		result = bound
	}
	return result, allSet
}

func getLimitForContainer(resource corev1.ResourceName, spec corev1.Container) (int64, bool) {
	if v, ok := spec.Resources.Limits[resource]; ok {
		var limit int64
		if resource == cpuKey {
			limit = v.MilliValue()
		} else {
			limit = v.Value()
		}
		return limit, true
	}
	return 0, false
}

func getRequestForContainer(resource corev1.ResourceName, spec corev1.Container) (int64, bool) {
	if v, ok := spec.Resources.Requests[resource]; ok {
		var req int64
		if resource == cpuKey {
			req = v.MilliValue()
		} else {
			req = v.Value()
		}
		return req, true
	}
	return 0, false
}

func addContainerId(pod *corev1.Pod, attributes pdata.AttributeMap, metric pdata.ResourceMetrics, kubernetesBlob map[string]interface{}) {
	if containerName := GetStringValFromAttributes(ContainerNamekey, attributes); containerName != "" {
		rawId := ""
		for _, container := range pod.Status.ContainerStatuses {
			if GetStringValFromAttributes(ContainerNamekey, attributes) == container.Name {
				rawId = container.ContainerID
				if rawId != "" {
					ids := strings.Split(rawId, "://")
					if len(ids) == 2 {
						kubernetesBlob[ids[0]] = map[string]string{"container_id": ids[1]}
					} else {
						log.Printf("W! Cannot parse container id from %s for container %s", rawId, container.Name)
						kubernetesBlob["container_id"] = rawId
					}
				}
				break
			}
		}
		if rawId == "" {
			kubernetesBlob["container_id"] = GetStringValFromAttributes(ContainerIdkey, attributes)
		}
		RemoveAttribute(ContainerIdkey, attributes)
	}
}

func addLabels(pod *corev1.Pod, kubernetesBlob map[string]interface{}) {
	labels := make(map[string]string)
	for k, v := range pod.Labels {
		labels[k] = v
	}
	if len(labels) > 0 {
		kubernetesBlob["labels"] = labels
	}
}

func getJobNamePrefix(podName string) string {
	return re.Split(podName, 2)[0]
}

func (p *PodStore) addPodOwnersAndPodName(attributes pdata.AttributeMap, pod *corev1.Pod, kubernetesBlob map[string]interface{}) {
	var owners []interface{}
	podName := ""
	for _, owner := range pod.OwnerReferences {
		if owner.Kind != "" && owner.Name != "" {
			kind := owner.Kind
			name := owner.Name
			if owner.Kind == ReplicaSet {
				rsToDeployment := k8sclient.Get().ReplicaSet.ReplicaSetToDeployment()
				if parent := rsToDeployment[owner.Name]; parent != "" {
					kind = Deployment
					name = parent
				} else if parent := parseDeploymentFromReplicaSet(owner.Name); parent != "" {
					kind = Deployment
					name = parent
				}
			} else if owner.Kind == Job {
				if parent := parseCronJobFromJob(owner.Name); parent != "" {
					kind = CronJob
					name = parent
				} else if !p.prefFullPodName {
					name = getJobNamePrefix(name)
				}
			}
			owners = append(owners, map[string]string{"owner_kind": kind, "owner_name": name})

			if podName == "" {
				if owner.Kind == StatefulSet {
					podName = pod.Name
				} else if owner.Kind == DaemonSet || owner.Kind == Job || owner.Kind == ReplicaSet || owner.Kind == ReplicationController {
					podName = name
				}
			}
		}
	}

	if len(owners) > 0 {
		kubernetesBlob["pod_owners"] = owners
	}

	// if podName is not set according to a well-known controllers, then set it to its own name
	if podName == "" {
		if strings.HasPrefix(pod.Name, kubeProxy) && !p.prefFullPodName {
			podName = kubeProxy
		} else {
			podName = pod.Name
		}
	}

	AddAttribute(PodNameKey, podName, attributes)
}

func addContainerCount(metric pdata.ResourceMetrics, attributes pdata.AttributeMap, pod *corev1.Pod) {
	runningContainerCount := 0
	for _, containerStatus := range pod.Status.ContainerStatuses {
		if containerStatus.State.Running != nil {
			runningContainerCount += 1
		}
	}
	if GetStringValFromAttributes(MetricType, attributes) == TypePod {
		addMetric(MetricName(TypePod, RunningContainerCount), runningContainerCount, UnitCount, metric)
		addMetric(MetricName(TypePod, ContainerCount), len(pod.Status.ContainerStatuses), UnitCount, metric)
	}
}
