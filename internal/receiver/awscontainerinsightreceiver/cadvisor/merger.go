package cadvisor

import (
	"fmt"

	common "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	"github.com/aws-observability/aws-otel-collector/internal/receiver/awscontainerinsightreceiver/cadvisor/extractors"
)

func mergeMetrics(metrics []*extractors.CAdvisorMetric) []*extractors.CAdvisorMetric {
	var result []*extractors.CAdvisorMetric
	metricMap := make(map[string]*extractors.CAdvisorMetric)
	for _, metric := range metrics {
		if metricKey := getMetricKey(metric); metricKey != "" {
			if mergedMetric, ok := metricMap[metricKey]; ok {
				mergedMetric.Merge(metric)
			} else {
				metricMap[metricKey] = metric
			}
		} else {
			// this metric cannot be merged
			result = append(result, metric)
		}
	}
	for _, metric := range metricMap {
		result = append(result, metric)
	}
	return result
}

// return MetricKey for merge-able metrics
func getMetricKey(metric *extractors.CAdvisorMetric) string {
	metricType := metric.GetMetricType()
	metricKey := ""
	switch metricType {
	case common.TypeInstance:
		// merge cpu, memory, net metric for type Instance
		metricKey = fmt.Sprintf("metricType:%s", common.TypeInstance)
	case common.TypeNode:
		// merge cpu, memory, net metric for type Node
		metricKey = fmt.Sprintf("metricType:%s", common.TypeNode)
	case common.TypePod:
		// merge cpu, memory, net metric for type Pod
		metricKey = fmt.Sprintf("metricType:%s,podId:%s", common.TypePod, metric.GetTags()[common.PodIdKey])
	case common.TypeContainer:
		// merge cpu, memory metric for type Container
		metricKey = fmt.Sprintf("metricType:%s,podId:%s,containerName:%s", common.TypeContainer, metric.GetTags()[common.PodIdKey], metric.GetTags()[common.ContainerNamekey])
	case common.TypeInstanceDiskIO:
		// merge io_serviced, io_service_bytes for type InstanceDiskIO
		metricKey = fmt.Sprintf("metricType:%s,device:%s", common.TypeInstanceDiskIO, metric.GetTags()[common.DiskDev])
	case common.TypeNodeDiskIO:
		// merge io_serviced, io_service_bytes for type NodeDiskIO
		metricKey = fmt.Sprintf("metricType:%s,device:%s", common.TypeNodeDiskIO, metric.GetTags()[common.DiskDev])
	default:
		metricKey = ""
	}
	return metricKey
}
