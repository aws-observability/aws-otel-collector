package ecs

import (
	"encoding/json"
	"fmt"
	common "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	"github.com/aws-observability/aws-otel-collector/internal/aws/httpclient"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"
	"strings"
	"sync"
	"time"
)

type EcsInfo struct {
	logger              *zap.Logger
	hostIP              string
	clusterName         string
	containerInstanceId string
	cgroup              *cgroupScanner
	runningTaskCount    int64
	cpuReserved         int64
	memReserved         int64
	refreshInterval     time.Duration
	shutdownC           chan bool
	httpClient          *httpclient.HttpClient
	sync.RWMutex
}

const (
	ecsAgentEndpoint         = "http://%s:51678/v1/metadata"
	ecsAgentTaskInfoEndpoint = "http://%s:51678/v1/tasks"
	taskStatusRunning        = "RUNNING"
)

type ContainerInstance struct {
	Cluster              string
	ContainerInstanceArn string
}

type ECSContainer struct {
	DockerId string
}
type ECSTask struct {
	KnownStatus string
	ARN         string
	Containers  []ECSContainer
}

type ECSTasksInfo struct {
	Tasks []ECSTask
}

func (e *EcsInfo) updateRunningTaskCount() {
	ecsTasksInfo := e.getTasksInfo()
	runningTaskCount := int64(0)
	cpuReserved := int64(0)
	memReserved := int64(0)
	for _, task := range ecsTasksInfo.Tasks {
		if task.KnownStatus != taskStatusRunning {
			continue
		}
		taskId, err := getTaskCgroupPathFromARN(task.ARN)
		if err != nil {
			e.logger.Warn("Failed to get ecs taskid from arn: ", zap.Error(err))
			continue
		}

		// ignore the one only consume 2 shares which is the default value in cgroup
		if cr := e.cgroup.getCPUReserved(taskId, e.clusterName, e.logger); cr > 2 {
			cpuReserved += cr
		}
		memReserved += e.cgroup.getMEMReserved(taskId, e.clusterName, task.Containers, e.logger)

		runningTaskCount += 1
	}

	e.Lock()
	defer e.Unlock()
	e.runningTaskCount = runningTaskCount
	e.cpuReserved = cpuReserved
	e.memReserved = memReserved
}

func (e *EcsInfo) getRunningTaskCount() int64 {
	e.RLock()
	defer e.RUnlock()
	return e.runningTaskCount
}

func (e *EcsInfo) getCpuReserved() int64 {
	e.RLock()
	defer e.RUnlock()
	return e.cpuReserved
}

func (e *EcsInfo) getMemReserved() int64 {
	e.RLock()
	defer e.RUnlock()
	return e.memReserved
}

func (e *EcsInfo) DecorateMetrics(mds []pdata.Metrics) {
	if len(mds) == 0 {
		e.logger.Warn("Empty metrics for ECS info")
		return
	}
	e.logger.Debug("collect data for ecs...")
	for _, md := range mds {
		rms := md.ResourceMetrics()
		for i := 0; i < rms.Len(); i++ {
			rm := rms.At(i)
			resource := rm.Resource()
			attributes := resource.Attributes()
			metricTypeValue, exited := attributes.Get(common.MetricType)
			if !exited {
				e.logger.Warn(common.MetricType + "is not exited in the metric")
			} else {
				if metricTypeValue.StringVal() == common.TypeInstance {
					e.addECSMetric(rm, common.TypeInstance)
				}
				e.addECSResources(metricTypeValue.StringVal(), resource)
			}
		}
	}

}

func (e *EcsInfo) addECSMetric(rm pdata.ResourceMetrics, metricType string) {
	cpuExist, cupLimits, cpuTimestamp := getMetricValue(common.InstancePrefix+common.CpuLimit, rm)
	memExist, memLimits, memTimestamp := getMetricValue(common.InstancePrefix+common.MemLimit, rm)
	ilms := rm.InstrumentationLibraryMetrics()
	e.logger.Warn("Insert metrics begin!")
	if cpuExist {
		cpuReservedPercentage := float64(e.cpuReserved) / (float64(cupLimits.(int64))) * 100
		insertMetrics(common.InstancePrefix+common.CpuReservedCapacity, cpuReservedPercentage, ilms, metricType, cpuTimestamp, e.logger)
	} else {
		e.logger.Warn("Can not get " + common.CpuLimit + " !")
	}
	if memExist {
		memReservedPercentage := float64(e.memReserved) / (float64(memLimits.(int64))) * 100
		insertMetrics(common.InstancePrefix+common.MemReservedCapacity, memReservedPercentage, ilms, metricType, memTimestamp, e.logger)
	} else {
		e.logger.Warn("Can not get " + common.MemLimit + " !")
	}
	if memTimestamp != cpuTimestamp {
		e.logger.Warn("The time of " + common.MemLimit + " is not equal to the time of " + common.CpuLimit)
	}
	insertMetrics(common.InstancePrefix+common.RunningTaskCount, e.runningTaskCount, ilms, metricType, memTimestamp, e.logger)
}

func (e *EcsInfo) addECSResources(metricType string, rs pdata.Resource) {
	var sources []string
	switch metricType {
	case common.TypeInstance:
		sources = append(sources, []string{"cadvisor", "/proc", "ecsagent", "calculated"}...)
	case common.TypeInstanceFS:
		sources = append(sources, []string{"cadvisor", "calculated"}...)
	case common.TypeInstanceNet:
		sources = append(sources, []string{"cadvisor", "calculated"}...)
	case common.TypeInstanceDiskIO:
		sources = append(sources, []string{"cadvisor"}...)
	}
	if len(sources) > 0 {
		sourcesInfo, err := json.Marshal(sources)
		if err != nil {
			return
		}
		AddAttribute(common.SourcesKey, string(sourcesInfo), rs.Attributes())
	}
	AddAttribute(common.ContainerInstanceIdKey, e.containerInstanceId, rs.Attributes())
	AddAttribute(common.ClusterNameKey, e.clusterName, rs.Attributes())
}

func (e *EcsInfo) getECSInfo() {
	containerInstance := e.getContainerInstanceInfo()
	e.clusterName = containerInstance.Cluster
	e.containerInstanceId = e.getContainerInstanceIdFromArn(containerInstance.ContainerInstanceArn)
	e.cgroup = newCGroupScannerForContainer(e.logger)
	e.updateRunningTaskCount()
	go func() {
		refreshTicker := time.NewTicker(e.refreshInterval)
		defer refreshTicker.Stop()
		for {
			select {
			case <-refreshTicker.C:
				e.updateRunningTaskCount()
			case <-e.shutdownC:
				refreshTicker.Stop()
				return
			}
		}
	}()
	return
}

func (e *EcsInfo) Shutdown() {
	close(e.shutdownC)
}

func (e *EcsInfo) getECSAgentEndpoint() string {
	return fmt.Sprintf(ecsAgentEndpoint, e.hostIP)
}

func (e *EcsInfo) getECSAgentTaskInfoEndpoint() string {
	return fmt.Sprintf(ecsAgentTaskInfoEndpoint, e.hostIP)

}

// There are two formats of ContainerInstance ARN (https://docs.aws.amazon.com/AmazonECS/latest/userguide/ecs-account-settings.html#ecs-resource-ids)
// arn:aws:ecs:region:aws_account_id:container-instance/container-instance-id
// arn:aws:ecs:region:aws_account_id:container-instance/cluster-name/container-instance-id
// This function will return "container-instance-id" for both ARN format
func (e *EcsInfo) getContainerInstanceIdFromArn(arn string) (containerInstanceId string) {
	// When splitting the ARN with ":", the 6th segments could be either:
	// container-instance/47c0ab6e-2c2c-475e-9c30-b878fa7a8c3d or
	// container-instance/cluster-name/47c0ab6e-2c2c-475e-9c30-b878fa7a8c3d
	if splitedList := strings.Split(arn, ":"); len(splitedList) >= 6 {
		// Further splitting tmpResult with "/", it could be splitted into either 2 or 3
		// Characters of "cluster-name" is only allowed to be letters, numbers and hyphens
		tmpResult := strings.Split(splitedList[5], "/")
		if len(tmpResult) == 2 {
			containerInstanceId = tmpResult[1]
			return
		} else if len(tmpResult) == 3 {
			containerInstanceId = tmpResult[2]
			return
		}
	}

	e.logger.Warn("Can't get ecs container instance id from ContainerInstance arn: " + arn)
	return

}

func (e *EcsInfo) getContainerInstanceInfo() (containerInstance *ContainerInstance) {
	containerInstance = &ContainerInstance{}
	resp, err := e.httpClient.Request(e.getECSAgentEndpoint())
	if err != nil {
		e.logger.Warn("Failed to call ecsagent endpoint, error: ", zap.Error(err))
		return containerInstance
	}

	err = json.Unmarshal(resp, containerInstance)
	if err != nil {
		e.logger.Warn("Unable to parse resp from ecsagent endpoint, error: ", zap.Error(err))
		e.logger.Warn("Resp content is " + string(resp))
	}
	return
}

func (e *EcsInfo) getTasksInfo() (ecsTasksInfo *ECSTasksInfo) {
	ecsTasksInfo = &ECSTasksInfo{}
	resp, err := e.httpClient.Request(e.getECSAgentTaskInfoEndpoint())
	if err != nil {
		e.logger.Warn("Failed to call ecsagent taskinfo endpoint, error: ", zap.Error(err))
		return ecsTasksInfo
	}

	err = json.Unmarshal(resp, ecsTasksInfo)
	if err != nil {
		e.logger.Warn("Unable to parse resp from ecsagent taskinfo endpoint, error:", zap.Error(err))
		e.logger.Warn("D! resp content is %s" + string(resp))
	}
	return
}

// There are two formats of Task ARN (https://docs.aws.amazon.com/AmazonECS/latest/userguide/ecs-account-settings.html#ecs-resource-ids)
// arn:aws:ecs:region:aws_account_id:task/task-id
// arn:aws:ecs:region:aws_account_id:task/cluster-name/task-id
// we should get "task-id" as result no matter what format the ARN is.
func getTaskCgroupPathFromARN(arn string) (string, error) {
	result := strings.Split(arn, ":")
	if len(result) < 6 {
		return "", fmt.Errorf("invalid ecs task arn: %v", result)
	}

	result = strings.Split(result[5], "/")
	if len(result) == 2 {
		return result[1], nil
	} else if len(result) == 3 {
		return result[2], nil
	} else {
		return "", fmt.Errorf("invalid ecs task arn: %v", result)
	}
}

func New(logger *zap.Logger, hostIP string) *EcsInfo {
	return newECSInfo(logger, hostIP)
}

func newECSInfo(logger *zap.Logger, hostIP string) (e *EcsInfo) {
	e = &EcsInfo{logger: logger, hostIP: hostIP, refreshInterval: 1 * time.Minute, shutdownC: make(chan bool), httpClient: httpclient.New()}
	e.getECSInfo()
	return
}
