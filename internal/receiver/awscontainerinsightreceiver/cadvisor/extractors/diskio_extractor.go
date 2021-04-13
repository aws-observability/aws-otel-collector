package extractors

import (
	"fmt"
	"strings"
	"time"

	aws "github.com/aws-observability/aws-otel-collector/internal/aws"
	. "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	cinfo "github.com/google/cadvisor/info/v1"
	"go.uber.org/zap"
)

type DiskIOMetricExtractor struct {
	logger         *zap.Logger
	rateCalculator aws.MetricCalculator
}

func (d *DiskIOMetricExtractor) HasValue(info *cinfo.ContainerInfo) bool {
	return info.Spec.HasDiskIo
}

func (d *DiskIOMetricExtractor) GetValue(info *cinfo.ContainerInfo, _ MachineInfoProvider, containerType string) []*CAdvisorMetric {
	var metrics []*CAdvisorMetric
	if containerType != TypeNode && containerType != TypeInstance {
		return metrics
	}

	curStats := GetStats(info)
	metrics = append(metrics, d.extractIoMetrics(curStats.DiskIo.IoServiceBytes, DiskIOServiceBytesPrefix, containerType, info.Name, curStats.Timestamp)...)
	metrics = append(metrics, d.extractIoMetrics(curStats.DiskIo.IoServiced, DiskIOServicedPrefix, containerType, info.Name, curStats.Timestamp)...)
	return metrics
}

func (d *DiskIOMetricExtractor) extractIoMetrics(curStatsSet []cinfo.PerDiskStats, namePrefix string, containerType string, infoName string, curTime time.Time) []*CAdvisorMetric {
	var metrics []*CAdvisorMetric
	expectedKey := []string{DiskIOAsync, DiskIOSync, DiskIORead, DiskIOWrite, DiskIOTotal}
	for _, cur := range curStatsSet {
		curDevName := devName(cur)
		metric := newCadvisorMetric(getDiskIOMetricType(containerType, d.logger))
		metric.tags[DiskDev] = curDevName
		for _, key := range expectedKey {
			if curVal, curOk := cur.Stats[key]; curOk {
				mname := MetricName(containerType, ioMetricName(namePrefix, key))
				assignRateValueToField(&d.rateCalculator, metric.fields, mname, infoName, float64(curVal), curTime, float64(time.Second))
			}
		}
		if len(metric.fields) > 0 {
			metrics = append(metrics, metric)
		}
	}
	return metrics
}

func ioMetricName(prefix, key string) string {
	return prefix + strings.ToLower(key)
}

func devName(dStats cinfo.PerDiskStats) string {
	devName := dStats.Device
	if devName == "" {
		devName = fmt.Sprintf("%d:%d", dStats.Major, dStats.Minor)
	}
	return devName
}

func NewDiskIOMetricExtractor(logger *zap.Logger) *DiskIOMetricExtractor {
	return &DiskIOMetricExtractor{
		logger:         logger,
		rateCalculator: newFloat64RateCalculator(),
	}
}

func getDiskIOMetricType(containerType string, logger *zap.Logger) string {
	metricType := ""
	switch containerType {
	case TypeNode:
		metricType = TypeNodeDiskIO
	case TypeInstance:
		metricType = TypeInstanceDiskIO
	case TypeContainer:
		metricType = TypeContainerDiskIO
	default:
		logger.Warn("diskio_extractor: diskIO metric extractor is parsing unexpected containerType", zap.String("containerType", containerType))
	}
	return metricType
}
