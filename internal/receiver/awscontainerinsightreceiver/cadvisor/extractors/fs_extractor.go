package extractors

import (
	"regexp"

	. "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	cinfo "github.com/google/cadvisor/info/v1"
	"go.uber.org/zap"
)

const (
	allowList = "^tmpfs$|^/dev/|^overlay$"
)

type FileSystemMetricExtractor struct {
	allowListRegexP *regexp.Regexp
	logger          *zap.Logger
}

func (f *FileSystemMetricExtractor) HasValue(info *cinfo.ContainerInfo) bool {
	return info.Spec.HasFilesystem
}

func (f *FileSystemMetricExtractor) GetValue(info *cinfo.ContainerInfo, _ MachineInfoProvider, containerType string) []*CAdvisorMetric {
	var metrics []*CAdvisorMetric
	if containerType == TypePod || info.Spec.Labels[containerNameLable] == infraContainerName {
		return metrics
	}

	containerType = getFSMetricType(containerType, f.logger)
	stats := GetStats(info)

	for _, v := range stats.Filesystem {
		metric := newCadvisorMetric(containerType)
		if v.Device == "" {
			continue
		}
		if f.allowListRegexP != nil && !f.allowListRegexP.MatchString(v.Device) {
			continue
		}

		metric.tags[DiskDev] = v.Device
		metric.tags[FSType] = v.Type

		metric.fields[MetricName(containerType, FSUsage)] = v.Usage
		metric.fields[MetricName(containerType, FSCapacity)] = v.Limit
		metric.fields[MetricName(containerType, FSAvailable)] = v.Available

		if v.Limit != 0 {
			metric.fields[MetricName(containerType, FSUtilization)] = float64(v.Usage) / float64(v.Limit) * 100
		}

		if v.HasInodes {
			metric.fields[MetricName(containerType, FSInodes)] = v.Inodes
			metric.fields[MetricName(containerType, FSInodesfree)] = v.InodesFree
		}

		metrics = append(metrics, metric)
	}
	return metrics
}

func NewFileSystemMetricExtractor(logger *zap.Logger) *FileSystemMetricExtractor {
	fse := &FileSystemMetricExtractor{
		logger: logger,
	}
	if p, err := regexp.Compile(allowList); err == nil {
		fse.allowListRegexP = p
	} else {
		logger.Error("NewFileSystemMetricExtractor set regex failed", zap.Error(err))
	}

	return fse
}

func getFSMetricType(containerType string, logger *zap.Logger) string {
	metricType := ""
	switch containerType {
	case TypeNode:
		metricType = TypeNodeFS
	case TypeInstance:
		metricType = TypeInstanceFS
	case TypeContainer:
		metricType = TypeContainerFS
	default:
		logger.Warn("fs_extractor: fs metric extractor is parsing unexpected containerType", zap.String("containerType", containerType))
	}
	return metricType
}
