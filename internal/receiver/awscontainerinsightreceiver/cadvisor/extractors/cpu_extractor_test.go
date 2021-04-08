package extractors

import (
	"testing"

	common "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
)

func TestCPUStats(t *testing.T) {
	mockMachineInfo := MockMachineInfo{}

	result := loadContainerInfo("./testdata/PreInfoContainer.json")
	result2 := loadContainerInfo("./testdata/CurInfoContainer.json")

	//test container type
	containerType := common.TypeContainer
	extractor := NewCpuMetricExtractor(nil)

	var cMetrics []*CAdvisorMetric
	if extractor.HasValue(result[0]) {
		cMetrics = extractor.GetValue(result[0], mockMachineInfo, containerType)
	}

	if extractor.HasValue(result2[0]) {
		cMetrics = extractor.GetValue(result2[0], mockMachineInfo, containerType)
	}

	AssertContainsTaggedFloat(t, cMetrics[0], "container_cpu_usage_total", 10, 0)
	AssertContainsTaggedFloat(t, cMetrics[0], "container_cpu_usage_user", 10, 0)
	AssertContainsTaggedFloat(t, cMetrics[0], "container_cpu_usage_system", 10, 0)
	AssertContainsTaggedFloat(t, cMetrics[0], "container_cpu_utilization", 0.5, 0)

	//test node type
	containerType = common.TypeNode
	extractor = NewCpuMetricExtractor(nil)

	if extractor.HasValue(result[0]) {
		cMetrics = extractor.GetValue(result[0], mockMachineInfo, containerType)
	}

	if extractor.HasValue(result2[0]) {
		cMetrics = extractor.GetValue(result2[0], mockMachineInfo, containerType)
	}

	AssertContainsTaggedFloat(t, cMetrics[0], "node_cpu_usage_total", 10, 0)
	AssertContainsTaggedFloat(t, cMetrics[0], "node_cpu_usage_user", 10, 0)
	AssertContainsTaggedFloat(t, cMetrics[0], "node_cpu_usage_system", 10, 0)
	AssertContainsTaggedFloat(t, cMetrics[0], "node_cpu_utilization", 0.5, 0)
	AssertContainsTaggedInt(t, cMetrics[0], "node_cpu_limit", 2000)
}
