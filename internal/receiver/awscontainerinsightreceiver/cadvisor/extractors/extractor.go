package extractors

import (
	"log"
	"time"

	aws "github.com/aws-observability/aws-otel-collector/internal/aws"
	common "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	cinfo "github.com/google/cadvisor/info/v1"
)

const (
	containerNameLable = "io.kubernetes.container.name"
	// TODO: https://github.com/containerd/cri/issues/922#issuecomment-423729537 the container name can be empty on containerd
	infraContainerName = "POD"
	Metrics            = "Metrics"
	Dimensions         = "Dimensions"
)

func GetStats(info *cinfo.ContainerInfo) *cinfo.ContainerStats {
	if len(info.Stats) == 0 {
		return nil
	}
	// When there is more than one stats point, always use the last one
	return info.Stats[len(info.Stats)-1]
}

type MachineInfoProvider interface {
	GetNumCores() int
	GetMemoryCapacity() uint64
}

type MetricExtractor interface {
	HasValue(*cinfo.ContainerInfo) bool
	GetValue(*cinfo.ContainerInfo, MachineInfoProvider, string) []*CAdvisorMetric
}

type CAdvisorMetric struct {
	//key/value pairs that are typed and contain the metric (numerical) data
	fields map[string]interface{}
	//key/value string pairs that are used to identify the metrics
	tags map[string]string
}

func newCadvisorMetric(mType string) *CAdvisorMetric {
	metric := &CAdvisorMetric{
		fields: make(map[string]interface{}),
		tags:   make(map[string]string),
	}
	metric.tags[common.MetricType] = mType
	return metric
}

func (c *CAdvisorMetric) GetTags() map[string]string {
	return c.tags
}

func (c *CAdvisorMetric) GetFields() map[string]interface{} {
	return c.fields
}

func (c *CAdvisorMetric) GetMetricType() string {
	return c.tags[common.MetricType]
}

func (c *CAdvisorMetric) AddTags(tags map[string]string) {
	for k, v := range tags {
		c.tags[k] = v
	}
}

func (c *CAdvisorMetric) Merge(src *CAdvisorMetric) {
	// If there is any conflict, keep the fields with earlier timestamp
	for k, v := range src.fields {
		if _, ok := c.fields[k]; ok {
			log.Printf("D! metric being merged has conflict in fields, src: %v, dest: %v \n", *src, *c)
			if c.tags[common.Timestamp] < src.tags[common.Timestamp] {
				continue
			}
		}
		c.fields[k] = v
	}
}

func newFloat64RateCalculator() aws.MetricCalculator {
	return aws.NewMetricCalculator(func(prev *aws.MetricValue, val interface{}, timestamp time.Time) (interface{}, bool) {
		if prev != nil {
			deltaNs := timestamp.Sub(prev.Timestamp).Nanoseconds()
			deltaValue := val.(float64) - prev.RawValue.(float64)
			if deltaNs > common.MinTimeDiff && deltaValue >= 0 {
				return deltaValue / float64(deltaNs), true
			}
		}
		return float64(0), false
	})
}

func assignRateValueToField(rateCalculator *aws.MetricCalculator, fields map[string]interface{}, metricName string,
	cinfoName string, curVal interface{}, curTime time.Time, multiplier float64) {
	key := cinfoName + metricName
	if val, ok := rateCalculator.Calculate(key, nil, curVal, curTime); ok {
		fields[metricName] = val.(float64) * multiplier
	}
}
