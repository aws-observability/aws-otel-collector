package extractors

import (
	aws "github.com/aws-observability/aws-otel-collector/internal/aws"
	. "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	cInfo "github.com/google/cadvisor/info/v1"
	"go.uber.org/zap"
)

const (
	decimalToMillicores = 1000
)

type CpuMetricExtractor struct {
	logger         *zap.Logger
	rateCalculator aws.MetricCalculator
}

func (c *CpuMetricExtractor) HasValue(info *cInfo.ContainerInfo) bool {
	return info.Spec.HasCpu
}

func (c *CpuMetricExtractor) GetValue(info *cInfo.ContainerInfo, mInfo MachineInfoProvider, containerType string) []*CAdvisorMetric {
	var metrics []*CAdvisorMetric
	if info.Spec.Labels[containerNameLable] == infraContainerName {
		return metrics
	}

	// When there is more than one stats point, always use the last one
	curStats := GetStats(info)
	metric := newCadvisorMetric(containerType)
	multiplier := float64(decimalToMillicores)
	assignRateValueToField(&c.rateCalculator, metric.fields, MetricName(containerType, CpuTotal), info.Name, float64(curStats.Cpu.Usage.Total), curStats.Timestamp, multiplier)
	assignRateValueToField(&c.rateCalculator, metric.fields, MetricName(containerType, CpuUser), info.Name, float64(curStats.Cpu.Usage.User), curStats.Timestamp, multiplier)
	assignRateValueToField(&c.rateCalculator, metric.fields, MetricName(containerType, CpuSystem), info.Name, float64(curStats.Cpu.Usage.System), curStats.Timestamp, multiplier)

	numCores := mInfo.GetNumCores()
	if metric.fields[MetricName(containerType, CpuTotal)] != nil && numCores != 0 {
		metric.fields[MetricName(containerType, CpuUtilization)] = metric.fields[MetricName(containerType, CpuTotal)].(float64) / float64(numCores*decimalToMillicores) * 100
	}

	if containerType == TypeNode {
		metric.fields[MetricName(containerType, CpuLimit)] = int64(numCores * decimalToMillicores)
	}

	metrics = append(metrics, metric)
	return metrics
}

func NewCpuMetricExtractor(logger *zap.Logger) *CpuMetricExtractor {
	return &CpuMetricExtractor{
		logger:         logger,
		rateCalculator: newFloat64RateCalculator(),
	}
}
