package extractors

import (
	"time"

	aws "github.com/aws-observability/aws-otel-collector/internal/aws"
	. "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	cinfo "github.com/google/cadvisor/info/v1"
	"go.uber.org/zap"
)

type MemMetricExtractor struct {
	logger         *zap.Logger
	rateCalculator aws.MetricCalculator
}

func (m *MemMetricExtractor) HasValue(info *cinfo.ContainerInfo) bool {
	return info.Spec.HasMemory
}

func (m *MemMetricExtractor) GetValue(info *cinfo.ContainerInfo, mInfo MachineInfoProvider, containerType string) []*CAdvisorMetric {
	var metrics []*CAdvisorMetric
	if info.Spec.Labels[containerNameLable] == infraContainerName {
		return metrics
	}

	metric := newCadvisorMetric(containerType)
	curStats := GetStats(info)

	metric.fields[MetricName(containerType, MemUsage)] = curStats.Memory.Usage
	metric.fields[MetricName(containerType, MemCache)] = curStats.Memory.Cache
	metric.fields[MetricName(containerType, MemRss)] = curStats.Memory.RSS
	metric.fields[MetricName(containerType, MemMaxusage)] = curStats.Memory.MaxUsage
	metric.fields[MetricName(containerType, MemSwap)] = curStats.Memory.Swap
	metric.fields[MetricName(containerType, MemFailcnt)] = curStats.Memory.Failcnt
	metric.fields[MetricName(containerType, MemMappedfile)] = curStats.Memory.MappedFile
	metric.fields[MetricName(containerType, MemWorkingset)] = curStats.Memory.WorkingSet

	multiplier := float64(time.Second)
	assignRateValueToField(&m.rateCalculator, metric.fields, MetricName(containerType, MemPgfault), info.Name,
		float64(curStats.Memory.ContainerData.Pgfault), curStats.Timestamp, multiplier)
	assignRateValueToField(&m.rateCalculator, metric.fields, MetricName(containerType, MemPgmajfault), info.Name,
		float64(curStats.Memory.ContainerData.Pgmajfault), curStats.Timestamp, multiplier)
	assignRateValueToField(&m.rateCalculator, metric.fields, MetricName(containerType, MemHierarchicalPgfault), info.Name,
		float64(curStats.Memory.HierarchicalData.Pgfault), curStats.Timestamp, multiplier)
	assignRateValueToField(&m.rateCalculator, metric.fields, MetricName(containerType, MemHierarchicalPgmajfault), info.Name,
		float64(curStats.Memory.HierarchicalData.Pgmajfault), curStats.Timestamp, multiplier)

	memoryCapacity := mInfo.GetMemoryCapacity()
	if metric.fields[MetricName(containerType, MemWorkingset)] != nil && memoryCapacity != 0 {
		metric.fields[MetricName(containerType, MemUtilization)] = float64(metric.fields[MetricName(containerType, MemWorkingset)].(uint64)) / float64(memoryCapacity) * 100
	}

	if containerType == TypeNode {
		metric.fields[MetricName(containerType, MemLimit)] = int64(memoryCapacity)
	}

	metrics = append(metrics, metric)
	return metrics
}

func NewMemMetricExtractor(logger *zap.Logger) *MemMetricExtractor {
	return &MemMetricExtractor{
		logger:         logger,
		rateCalculator: newFloat64RateCalculator(),
	}
}
