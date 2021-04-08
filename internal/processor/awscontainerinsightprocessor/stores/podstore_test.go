package stores

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"

	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"

	common "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	"github.com/aws-observability/aws-otel-collector/internal/aws/k8scommon/k8sclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	corev1 "k8s.io/api/core/v1"
)

func getBaseTestPodInfo() *corev1.Pod {
	podJson := `
{
  "kind": "PodList",
  "apiVersion": "v1",
  "metadata": {

  },
  "items": [
    {
      "metadata": {
        "name": "cpu-limit",
        "namespace": "default",
        "ownerReferences": [
            {
                "apiVersion": "apps/v1",
                "blockOwnerDeletion": true,
                "controller": true,
                "kind": "DaemonSet",
                "name": "DaemonSetTest",
                "uid": "36779a62-4aca-11e9-977b-0672b6c6fc94"
            }
        ],
        "selfLink": "/api/v1/namespaces/default/pods/cpu-limit",
        "uid": "764d01e1-2a2f-11e9-95ea-0a695d7ce286",
        "resourceVersion": "5671573",
        "creationTimestamp": "2019-02-06T16:51:34Z",
        "labels": {
          "app": "hello_test"
        },
        "annotations": {
          "kubernetes.io/config.seen": "2019-02-19T00:06:56.109155665Z",
          "kubernetes.io/config.source": "api"
        }
      },
      "spec": {
        "volumes": [
          {
            "name": "default-token-tlgw7",
            "secret": {
              "secretName": "default-token-tlgw7",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "ubuntu",
            "image": "ubuntu",
            "command": [
              "/bin/bash"
            ],
            "args": [
              "-c",
              "sleep 300000000"
            ],
            "resources": {
              "limits": {
                "cpu": "10m",
                "memory": "50Mi"
              },
              "requests": {
                "cpu": "10m",
                "memory": "50Mi"
              }
            },
            "volumeMounts": [
              {
                "name": "default-token-tlgw7",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "Always"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "default",
        "serviceAccount": "default",
        "nodeName": "ip-192-168-67-127.us-west-2.compute.internal",
        "securityContext": {

        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-02-06T16:51:34Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-02-06T16:51:43Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": null
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-02-06T16:51:34Z"
          }
        ],
        "hostIP": "192.168.67.127",
        "podIP": "192.168.76.93",
        "startTime": "2019-02-06T16:51:34Z",
        "containerStatuses": [
          {
            "name": "ubuntu",
            "state": {
              "running": {
                "startedAt": "2019-02-06T16:51:42Z"
              }
            },
            "lastState": {

            },
            "ready": true,
            "restartCount": 0,
            "image": "ubuntu:latest",
            "imageID": "docker-pullable://ubuntu@sha256:7a47ccc3bbe8a451b500d2b53104868b46d60ee8f5b35a24b41a86077c650210",
            "containerID": "docker://637631e2634ea92c0c1aa5d24734cfe794f09c57933026592c12acafbaf6972c"
          }
        ],
        "qosClass": "Guaranteed"
      }
    }
  ]
}`
	pods := corev1.PodList{}
	err := json.Unmarshal([]byte(podJson), &pods)
	if err != nil {
		panic(fmt.Sprintf("unmarshal pod err %v", err))
	}

	return &pods.Items[0]
}

func getPodStore() *PodStore {
	nodeInfo := newNodeInfo()
	nodeInfo.setCPUCapacity(4000)
	nodeInfo.setMemCapacity(400 * 1024 * 1024)
	return &PodStore{
		cache:            newMapWithExpiry(time.Minute),
		nodeInfo:         nodeInfo,
		prevMeasurements: make(map[string]*mapWithExpiry),
	}
}

func getResourceMetricsAndAttributes(fields map[string]interface{}, tags map[string]string, t time.Time) (pdata.ResourceMetrics, pdata.AttributeMap) {
	tags[common.Timestamp] = strconv.FormatInt(t.UnixNano(), 10)
	md := common.ConvertToOTLPMetrics(fields, tags, zap.NewNop())
	rm := md.ResourceMetrics().At(0)
	attributes := rm.Resource().Attributes()
	return rm, attributes
}

func TestPodStore_decorateCpu(t *testing.T) {
	podStore := getPodStore()

	pod := getBaseTestPodInfo()

	tags := map[string]string{common.MetricType: common.TypePod}
	fields := map[string]interface{}{common.MetricName(common.TypePod, common.CpuTotal): float64(1)}

	rm, attributes := getResourceMetricsAndAttributes(fields, tags, time.Now())
	podStore.decorateCpu(rm, attributes, pod)

	assert.Equal(t, int64(10), getMetricValue("pod_cpu_request", rm))
	assert.Equal(t, int64(10), getMetricValue("pod_cpu_limit", rm))
	assert.Equal(t, float64(0.25), getMetricValue("pod_cpu_reserved_capacity", rm))
	assert.Equal(t, float64(10), getMetricValue("pod_cpu_utilization_over_pod_limit", rm))
	assert.Equal(t, float64(1), getMetricValue("pod_cpu_usage_total", rm))
}

func TestPodStore_decorateMem(t *testing.T) {
	podStore := getPodStore()
	pod := getBaseTestPodInfo()

	tags := map[string]string{common.MetricType: common.TypePod}
	fields := map[string]interface{}{common.MetricName(common.TypePod, common.MemWorkingset): float64(10 * 1024 * 1024)}

	rm, attributes := getResourceMetricsAndAttributes(fields, tags, time.Now())
	podStore.decorateMem(rm, attributes, pod)

	assert.Equal(t, int64(52428800), getMetricValue("pod_memory_request", rm))
	assert.Equal(t, int64(52428800), getMetricValue("pod_memory_limit", rm))
	assert.Equal(t, float64(12.5), getMetricValue("pod_memory_reserved_capacity", rm))
	assert.Equal(t, float64(20), getMetricValue("pod_memory_utilization_over_pod_limit", rm))
	assert.Equal(t, float64(10*1024*1024), getMetricValue("pod_memory_working_set", rm))
}

func TestPodStore_addContainerCount(t *testing.T) {
	pod := getBaseTestPodInfo()

	tags := map[string]string{common.MetricType: common.TypePod}
	fields := map[string]interface{}{common.MetricName(common.TypePod, common.CpuTotal): float64(1)}

	rm, attributes := getResourceMetricsAndAttributes(fields, tags, time.Now())

	addContainerCount(rm, attributes, pod)
	assert.Equal(t, int64(1), getMetricValue(common.MetricName(common.TypePod, common.RunningContainerCount), rm))
	assert.Equal(t, int64(1), getMetricValue(common.MetricName(common.TypePod, common.ContainerCount), rm))

	rm, attributes = getResourceMetricsAndAttributes(fields, tags, time.Now())

	pod.Status.ContainerStatuses[0].State.Running = nil
	addContainerCount(rm, attributes, pod)
	assert.Equal(t, int64(0), getMetricValue(common.MetricName(common.TypePod, common.RunningContainerCount), rm))
	assert.Equal(t, int64(1), getMetricValue(common.MetricName(common.TypePod, common.ContainerCount), rm))
}

func TestPodStore_addStatus(t *testing.T) {
	pod := getBaseTestPodInfo()
	tags := map[string]string{common.MetricType: common.TypePod, common.K8sNamespace: "default", common.K8sPodNameKey: "cpu-limit"}
	fields := map[string]interface{}{common.MetricName(common.TypePod, common.CpuTotal): float64(1)}
	podStore := getPodStore()
	rm, attributes := getResourceMetricsAndAttributes(fields, tags, time.Now())

	podStore.addStatus(rm, attributes, pod)
	assert.Equal(t, "Running", GetStringValFromAttributes(common.PodStatus, attributes))
	val := getMetricValue(common.MetricName(common.TypePod, common.ContainerRestartCount), rm)
	assert.Nil(t, val)

	tags = map[string]string{common.MetricType: common.TypeContainer, common.K8sNamespace: "default", common.K8sPodNameKey: "cpu-limit", common.ContainerNamekey: "ubuntu"}
	rm, attributes = getResourceMetricsAndAttributes(fields, tags, time.Now())

	podStore.addStatus(rm, attributes, pod)
	assert.Equal(t, "Running", GetStringValFromAttributes(common.ContainerStatus, attributes))
	val = getMetricValue(common.ContainerRestartCount, rm)
	assert.Nil(t, val)

	pod.Status.ContainerStatuses[0].State.Running = nil
	pod.Status.ContainerStatuses[0].State.Terminated = &corev1.ContainerStateTerminated{}
	pod.Status.ContainerStatuses[0].LastTerminationState.Terminated = &corev1.ContainerStateTerminated{Reason: "OOMKilled"}
	pod.Status.ContainerStatuses[0].RestartCount = 1
	pod.Status.Phase = "Succeeded"

	tags = map[string]string{common.MetricType: common.TypePod, common.K8sNamespace: "default", common.K8sPodNameKey: "cpu-limit"}
	rm, attributes = getResourceMetricsAndAttributes(fields, tags, time.Now())

	podStore.addStatus(rm, attributes, pod)
	assert.Equal(t, "Succeeded", GetStringValFromAttributes(common.PodStatus, attributes))
	assert.Equal(t, int64(1), getMetricValue(common.MetricName(common.TypePod, common.ContainerRestartCount), rm))

	tags = map[string]string{common.MetricType: common.TypeContainer, common.K8sNamespace: "default", common.K8sPodNameKey: "cpu-limit", common.ContainerNamekey: "ubuntu"}
	rm, attributes = getResourceMetricsAndAttributes(fields, tags, time.Now())

	podStore.addStatus(rm, attributes, pod)
	assert.Equal(t, "Terminated", GetStringValFromAttributes(common.ContainerStatus, attributes))
	assert.Equal(t, "OOMKilled", GetStringValFromAttributes(common.ContainerLastTerminationReason, attributes))
	assert.Equal(t, int64(1), getMetricValue(common.ContainerRestartCount, rm))

	// test delta of restartCount
	pod.Status.ContainerStatuses[0].RestartCount = 3
	tags = map[string]string{common.MetricType: common.TypePod, common.K8sNamespace: "default", common.K8sPodNameKey: "cpu-limit"}
	rm, attributes = getResourceMetricsAndAttributes(fields, tags, time.Now())

	podStore.addStatus(rm, attributes, pod)
	assert.Equal(t, int64(2), getMetricValue(common.MetricName(common.TypePod, common.ContainerRestartCount), rm))

	tags = map[string]string{common.MetricType: common.TypeContainer, common.K8sNamespace: "default", common.K8sPodNameKey: "cpu-limit", common.ContainerNamekey: "ubuntu"}
	rm, attributes = getResourceMetricsAndAttributes(fields, tags, time.Now())

	podStore.addStatus(rm, attributes, pod)
	assert.Equal(t, int64(2), getMetricValue(common.ContainerRestartCount, rm))
}

func TestPodStore_addContainerId(t *testing.T) {
	pod := getBaseTestPodInfo()
	tags := map[string]string{common.ContainerNamekey: "ubuntu", common.ContainerIdkey: "123"}
	fields := map[string]interface{}{common.MetricName(common.TypePod, common.CpuTotal): float64(1)}
	kubernetesBlob := map[string]interface{}{}
	rm, attributes := getResourceMetricsAndAttributes(fields, tags, time.Now())
	addContainerId(pod, attributes, rm, kubernetesBlob)

	expected := map[string]interface{}{}
	expected["docker"] = map[string]string{"container_id": "637631e2634ea92c0c1aa5d24734cfe794f09c57933026592c12acafbaf6972c"}
	assert.Equal(t, expected, kubernetesBlob)
	assert.Equal(t, GetStringValFromAttributes(common.ContainerNamekey, attributes), "ubuntu")

	tags = map[string]string{common.ContainerNamekey: "notUbuntu", common.ContainerIdkey: "123"}
	kubernetesBlob = map[string]interface{}{}
	rm, attributes = getResourceMetricsAndAttributes(fields, tags, time.Now())
	addContainerId(pod, attributes, rm, kubernetesBlob)

	expected = map[string]interface{}{}
	expected["container_id"] = "123"
	assert.Equal(t, expected, kubernetesBlob)
	assert.Equal(t, GetStringValFromAttributes(common.ContainerNamekey, attributes), "notUbuntu")
}

func TestPodStore_addLabel(t *testing.T) {
	pod := getBaseTestPodInfo()
	kubernetesBlob := map[string]interface{}{}
	addLabels(pod, kubernetesBlob)
	expected := map[string]interface{}{}
	expected["labels"] = map[string]string{"app": "hello_test"}
	assert.Equal(t, expected, kubernetesBlob)
}

//
// Mock client start
//
var mockClient = new(MockClient)

var mockK8sClient = &k8sclient.K8sClient{
	Job:        mockClient,
	ReplicaSet: mockClient,
}

func mockGet() *k8sclient.K8sClient {
	return mockK8sClient
}

type MockClient struct {
	k8sclient.JobClient
	k8sclient.ReplicaSetClient

	mock.Mock
}

// k8sclient.JobClient
func (client *MockClient) JobToCronJob() map[string]string {
	args := client.Called()
	return args.Get(0).(map[string]string)
}

// k8sclient.ReplicaSetClient
func (client *MockClient) ReplicaSetToDeployment() map[string]string {
	args := client.Called()
	return args.Get(0).(map[string]string)
}

func (client *MockClient) Init() {
}

func (client *MockClient) Shutdown() {
}

//
// Mock client end
//

//
// Mock client 2 start
//
var mockClient2 = new(MockClient2)

var mockK8sClient2 = &k8sclient.K8sClient{
	Job:        mockClient2,
	ReplicaSet: mockClient2,
}

func mockGet2() *k8sclient.K8sClient {
	return mockK8sClient2
}

type MockClient2 struct {
	k8sclient.JobClient
	k8sclient.ReplicaSetClient

	mock.Mock
}

// k8sclient.JobClient
func (client *MockClient2) JobToCronJob() map[string]string {
	args := client.Called()
	return args.Get(0).(map[string]string)
}

// k8sclient.ReplicaSetClient
func (client *MockClient2) ReplicaSetToDeployment() map[string]string {
	args := client.Called()
	return args.Get(0).(map[string]string)
}

func (client *MockClient2) Init() {
}

func (client *MockClient2) Shutdown() {
}

//
// Mock client 2 end
//

func TestGetJobNamePrefix(t *testing.T) {
	assert.Equal(t, "abcd", getJobNamePrefix("abcd-efg"))
	assert.Equal(t, "abcd", getJobNamePrefix("abcd.efg"))
	assert.Equal(t, "abcd", getJobNamePrefix("abcd-e.fg"))
	assert.Equal(t, "abc", getJobNamePrefix("abc.d-efg"))
	assert.Equal(t, "abcd", getJobNamePrefix("abcd-.efg"))
	assert.Equal(t, "abcd", getJobNamePrefix("abcd.-efg"))
	assert.Equal(t, "abcdefg", getJobNamePrefix("abcdefg"))
	assert.Equal(t, "abcdefg", getJobNamePrefix("abcdefg-"))
	assert.Equal(t, "", getJobNamePrefix(".abcd-efg"))
	assert.Equal(t, "", getJobNamePrefix(""))
}

func TestPodStore_addPodOwnersAndPodNameFallback(t *testing.T) {
	k8sclient.Get = mockGet2
	mockClient2.On("JobToCronJob").Return(map[string]string{})
	mockClient2.On("ReplicaSetToDeployment").Return(map[string]string{})

	podStore := &PodStore{}
	pod := getBaseTestPodInfo()
	tags := map[string]string{common.MetricType: common.TypePod, common.ContainerNamekey: "ubuntu"}
	fields := map[string]interface{}{common.MetricName(common.TypePod, common.CpuTotal): float64(1)}
	_, attributes := getResourceMetricsAndAttributes(fields, tags, time.Now())

	// Test ReplicaSet
	rsName := "ReplicaSetTest"
	suffix := "-42kcz"
	pod.OwnerReferences[0].Kind = common.ReplicaSet
	pod.OwnerReferences[0].Name = rsName + suffix
	kubernetesBlob := map[string]interface{}{}
	podStore.addPodOwnersAndPodName(attributes, pod, kubernetesBlob)
	expectedOwner := map[string]interface{}{}
	expectedOwner["pod_owners"] = []interface{}{map[string]string{"owner_kind": common.Deployment, "owner_name": rsName}}
	expectedOwnerName := rsName
	assert.Equal(t, expectedOwnerName, GetStringValFromAttributes(common.PodNameKey, attributes))
	assert.Equal(t, expectedOwner, kubernetesBlob)

	// Test Job
	_, attributes = getResourceMetricsAndAttributes(fields, tags, time.Now())
	jobName := "Job"
	suffix = "-0123456789"
	pod.OwnerReferences[0].Kind = common.Job
	pod.OwnerReferences[0].Name = jobName + suffix
	kubernetesBlob = map[string]interface{}{}
	podStore.addPodOwnersAndPodName(attributes, pod, kubernetesBlob)
	expectedOwner["pod_owners"] = []interface{}{map[string]string{"owner_kind": common.CronJob, "owner_name": jobName}}
	expectedOwnerName = jobName
	assert.Equal(t, expectedOwnerName, GetStringValFromAttributes(common.PodNameKey, attributes))
	assert.Equal(t, expectedOwner, kubernetesBlob)
}

func TestPodStore_addPodOwnersAndPodName(t *testing.T) {
	k8sclient.Get = mockGet
	mockClient.On("JobToCronJob").Return(map[string]string{"CronJobTest-1556582405": "CronJobTest"})
	mockClient.On("ReplicaSetToDeployment").Return(map[string]string{"DeploymentTest-sftrz2785": "DeploymentTest"})

	podStore := &PodStore{}

	pod := getBaseTestPodInfo()
	tags := map[string]string{common.MetricType: common.TypePod, common.ContainerNamekey: "ubuntu"}
	fields := map[string]interface{}{common.MetricName(common.TypePod, common.CpuTotal): float64(1)}

	// Test DaemonSet
	_, attributes := getResourceMetricsAndAttributes(fields, tags, time.Now())
	kubernetesBlob := map[string]interface{}{}
	podStore.addPodOwnersAndPodName(attributes, pod, kubernetesBlob)

	expectedOwner := map[string]interface{}{}
	expectedOwner["pod_owners"] = []interface{}{map[string]string{"owner_kind": common.DaemonSet, "owner_name": "DaemonSetTest"}}
	expectedOwnerName := "DaemonSetTest"
	assert.Equal(t, expectedOwnerName, GetStringValFromAttributes(common.PodNameKey, attributes))
	assert.Equal(t, expectedOwner, kubernetesBlob)

	// Test ReplicaSet
	_, attributes = getResourceMetricsAndAttributes(fields, tags, time.Now())
	rsName := "ReplicaSetTest"
	pod.OwnerReferences[0].Kind = common.ReplicaSet
	pod.OwnerReferences[0].Name = rsName
	kubernetesBlob = map[string]interface{}{}
	podStore.addPodOwnersAndPodName(attributes, pod, kubernetesBlob)
	expectedOwner["pod_owners"] = []interface{}{map[string]string{"owner_kind": common.ReplicaSet, "owner_name": rsName}}
	expectedOwnerName = rsName
	assert.Equal(t, expectedOwnerName, GetStringValFromAttributes(common.PodNameKey, attributes))
	assert.Equal(t, expectedOwner, kubernetesBlob)

	// Test StatefulSet
	_, attributes = getResourceMetricsAndAttributes(fields, tags, time.Now())
	ssName := "StatefulSetTest"
	pod.OwnerReferences[0].Kind = common.StatefulSet
	pod.OwnerReferences[0].Name = ssName
	kubernetesBlob = map[string]interface{}{}
	podStore.addPodOwnersAndPodName(attributes, pod, kubernetesBlob)
	expectedOwner["pod_owners"] = []interface{}{map[string]string{"owner_kind": common.StatefulSet, "owner_name": ssName}}
	expectedOwnerName = "cpu-limit"
	assert.Equal(t, expectedOwnerName, GetStringValFromAttributes(common.PodNameKey, attributes))
	assert.Equal(t, expectedOwner, kubernetesBlob)

	// Test ReplicationController
	rcName := "ReplicationControllerTest"
	pod.OwnerReferences[0].Kind = common.ReplicationController
	pod.OwnerReferences[0].Name = rcName
	kubernetesBlob = map[string]interface{}{}
	podStore.addPodOwnersAndPodName(attributes, pod, kubernetesBlob)
	expectedOwner["pod_owners"] = []interface{}{map[string]string{"owner_kind": common.ReplicationController, "owner_name": rcName}}
	expectedOwnerName = rcName
	assert.Equal(t, expectedOwnerName, GetStringValFromAttributes(common.PodNameKey, attributes))
	assert.Equal(t, expectedOwner, kubernetesBlob)

	// Test Job
	podStore.prefFullPodName = true
	_, attributes = getResourceMetricsAndAttributes(fields, tags, time.Now())
	jobName := "JobTest"
	pod.OwnerReferences[0].Kind = common.Job
	surfixHash := ".088123x12"
	pod.OwnerReferences[0].Name = jobName + surfixHash
	kubernetesBlob = map[string]interface{}{}
	podStore.addPodOwnersAndPodName(attributes, pod, kubernetesBlob)
	expectedOwner["pod_owners"] = []interface{}{map[string]string{"owner_kind": common.Job, "owner_name": jobName + surfixHash}}
	expectedOwnerName = jobName + surfixHash
	assert.Equal(t, expectedOwnerName, GetStringValFromAttributes(common.PodNameKey, attributes))
	assert.Equal(t, expectedOwner, kubernetesBlob)

	podStore.prefFullPodName = false
	kubernetesBlob = map[string]interface{}{}
	podStore.addPodOwnersAndPodName(attributes, pod, kubernetesBlob)
	expectedOwner["pod_owners"] = []interface{}{map[string]string{"owner_kind": common.Job, "owner_name": jobName}}
	expectedOwnerName = jobName
	assert.Equal(t, expectedOwnerName, GetStringValFromAttributes(common.PodNameKey, attributes))
	assert.Equal(t, expectedOwner, kubernetesBlob)

	// Test Deployment
	_, attributes = getResourceMetricsAndAttributes(fields, tags, time.Now())
	dpName := "DeploymentTest"
	pod.OwnerReferences[0].Kind = common.ReplicaSet
	pod.OwnerReferences[0].Name = dpName + "-sftrz2785"
	kubernetesBlob = map[string]interface{}{}
	podStore.addPodOwnersAndPodName(attributes, pod, kubernetesBlob)
	expectedOwner["pod_owners"] = []interface{}{map[string]string{"owner_kind": common.Deployment, "owner_name": dpName}}
	expectedOwnerName = dpName
	assert.Equal(t, expectedOwnerName, GetStringValFromAttributes(common.PodNameKey, attributes))
	assert.Equal(t, expectedOwner, kubernetesBlob)

	// Test CronJob
	_, attributes = getResourceMetricsAndAttributes(fields, tags, time.Now())
	cjName := "CronJobTest"
	pod.OwnerReferences[0].Kind = common.Job
	pod.OwnerReferences[0].Name = cjName + "-1556582405"
	kubernetesBlob = map[string]interface{}{}
	podStore.addPodOwnersAndPodName(attributes, pod, kubernetesBlob)
	expectedOwner["pod_owners"] = []interface{}{map[string]string{"owner_kind": common.CronJob, "owner_name": cjName}}
	expectedOwnerName = cjName
	assert.Equal(t, expectedOwnerName, GetStringValFromAttributes(common.PodNameKey, attributes))
	assert.Equal(t, expectedOwner, kubernetesBlob)

	// Test kube-proxy created in kops
	podStore.prefFullPodName = true
	_, attributes = getResourceMetricsAndAttributes(fields, tags, time.Now())
	kpName := kubeProxy + "-xyz1"
	pod.OwnerReferences = nil
	pod.Name = kpName
	kubernetesBlob = map[string]interface{}{}
	podStore.addPodOwnersAndPodName(attributes, pod, kubernetesBlob)
	assert.Equal(t, kpName, GetStringValFromAttributes(common.PodNameKey, attributes))
	assert.True(t, len(kubernetesBlob) == 0)

	podStore.prefFullPodName = false
	_, attributes = getResourceMetricsAndAttributes(fields, tags, time.Now())
	pod.OwnerReferences = nil
	pod.Name = kpName
	kubernetesBlob = map[string]interface{}{}
	podStore.addPodOwnersAndPodName(attributes, pod, kubernetesBlob)
	assert.Equal(t, kubeProxy, GetStringValFromAttributes(common.PodNameKey, attributes))
	assert.True(t, len(kubernetesBlob) == 0)
}

func TestPodStore_refreshInternal(t *testing.T) {
	pod := getBaseTestPodInfo()
	podList := []corev1.Pod{*pod}

	podStore := getPodStore()
	podStore.refreshInternal(time.Now(), podList)

	assert.Equal(t, int64(10), podStore.nodeInfo.nodeStats.cpuReq)
	assert.Equal(t, int64(50*1024*1024), podStore.nodeInfo.nodeStats.memReq)
	assert.Equal(t, 1, podStore.nodeInfo.nodeStats.podCnt)
	assert.Equal(t, 1, podStore.nodeInfo.nodeStats.containerCnt)
	assert.Equal(t, 1, podStore.cache.Size())
}

func TestPodStore_decorateNode(t *testing.T) {
	pod := getBaseTestPodInfo()
	podList := []corev1.Pod{*pod}

	podStore := getPodStore()
	podStore.refreshInternal(time.Now(), podList)

	tags := map[string]string{common.MetricType: common.TypeNode}
	fields := map[string]interface{}{
		common.MetricName(common.TypeNode, common.CpuTotal):      float64(100),
		common.MetricName(common.TypeNode, common.CpuLimit):      int64(4000),
		common.MetricName(common.TypeNode, common.MemWorkingset): float64(100 * 1024 * 1024),
		common.MetricName(common.TypeNode, common.MemLimit):      int64(400 * 1024 * 1024),
	}

	rm, _ := getResourceMetricsAndAttributes(fields, tags, time.Now())
	podStore.decorateNode(rm)

	assert.Equal(t, int64(10), getMetricValue("node_cpu_request", rm))
	assert.Equal(t, int64(4000), getMetricValue("node_cpu_limit", rm))
	assert.Equal(t, float64(0.25), getMetricValue("node_cpu_reserved_capacity", rm))
	assert.Equal(t, float64(100), getMetricValue("node_cpu_usage_total", rm))

	assert.Equal(t, int64(50*1024*1024), getMetricValue("node_memory_request", rm))
	assert.Equal(t, int64(400*1024*1024), getMetricValue("node_memory_limit", rm))
	assert.Equal(t, float64(12.5), getMetricValue("node_memory_reserved_capacity", rm))
	assert.Equal(t, float64(100*1024*1024), getMetricValue("node_memory_working_set", rm))

	assert.Equal(t, int64(1), getMetricValue("node_number_of_running_containers", rm))
	assert.Equal(t, int64(1), getMetricValue("node_number_of_running_pods", rm))
}
