package extractors

import (
	"log"
	"testing"

	. "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	"github.com/stretchr/testify/assert"
)

func TestDiskIOStats(t *testing.T) {

	result := loadContainerInfo("./testdata/PreInfoContainer.json")
	result2 := loadContainerInfo("./testdata/CurInfoContainer.json")
	//for eks node-level metrics
	containerType := TypeNode
	extractor := NewDiskIOMetricExtractor(nil)

	var cMetrics []*CAdvisorMetric
	if extractor.HasValue(result[0]) {
		cMetrics = extractor.GetValue(result[0], nil, containerType)
	}

	if extractor.HasValue(result2[0]) {
		cMetrics = extractor.GetValue(result2[0], nil, containerType)
	}

	for _, cadvisorMetric := range cMetrics {
		log.Printf("cadvisor Metrics received:\n %v \n", *cadvisorMetric)
	}
	expectedFieldsService := map[string]interface{}{
		"node_diskio_io_service_bytes_write": float64(10000),
		"node_diskio_io_service_bytes_total": float64(10010),
		"node_diskio_io_service_bytes_async": float64(10000),
		"node_diskio_io_service_bytes_sync":  float64(10000),
		"node_diskio_io_service_bytes_read":  float64(10),
	}
	expectedFieldsServiced := map[string]interface{}{
		"node_diskio_io_serviced_async": float64(10),
		"node_diskio_io_serviced_sync":  float64(10),
		"node_diskio_io_serviced_read":  float64(10),
		"node_diskio_io_serviced_write": float64(10),
		"node_diskio_io_serviced_total": float64(20),
	}
	expectedTags := map[string]string{
		"device": "/dev/xvda",
		"Type":   "NodeDiskIO",
	}
	AssertContainsTaggedField(t, cMetrics[0], expectedFieldsService, expectedTags)
	AssertContainsTaggedField(t, cMetrics[1], expectedFieldsServiced, expectedTags)

	//for ecs node-level metrics
	containerType = TypeInstance
	extractor = NewDiskIOMetricExtractor(nil)

	if extractor.HasValue(result[0]) {
		cMetrics = extractor.GetValue(result[0], nil, containerType)
	}

	if extractor.HasValue(result2[0]) {
		cMetrics = extractor.GetValue(result2[0], nil, containerType)
	}

	for _, cadvisorMetric := range cMetrics {
		log.Printf("cadvisor Metrics received:\n %v \n", *cadvisorMetric)
	}
	expectedFieldsService = map[string]interface{}{
		"instance_diskio_io_service_bytes_write": float64(10000),
		"instance_diskio_io_service_bytes_total": float64(10010),
		"instance_diskio_io_service_bytes_async": float64(10000),
		"instance_diskio_io_service_bytes_sync":  float64(10000),
		"instance_diskio_io_service_bytes_read":  float64(10),
	}
	expectedFieldsServiced = map[string]interface{}{
		"instance_diskio_io_serviced_async": float64(10),
		"instance_diskio_io_serviced_sync":  float64(10),
		"instance_diskio_io_serviced_read":  float64(10),
		"instance_diskio_io_serviced_write": float64(10),
		"instance_diskio_io_serviced_total": float64(20),
	}
	expectedTags = map[string]string{
		"device": "/dev/xvda",
		"Type":   "InstanceDiskIO",
	}
	AssertContainsTaggedField(t, cMetrics[0], expectedFieldsService, expectedTags)
	AssertContainsTaggedField(t, cMetrics[1], expectedFieldsServiced, expectedTags)

	//for non supported type
	containerType = TypeContainerDiskIO
	extractor = NewDiskIOMetricExtractor(nil)

	if extractor.HasValue(result[0]) {
		cMetrics = extractor.GetValue(result[0], nil, containerType)
	}

	if extractor.HasValue(result2[0]) {
		cMetrics = extractor.GetValue(result2[0], nil, containerType)
	}

	assert.Equal(t, len(cMetrics), 0)
}
