package structuredlogsadapter

import (
	"strconv"
	"testing"
	"time"

	. "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	. "github.com/aws-observability/aws-otel-collector/internal/processor/awscontainerinsightprocessor/stores"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestUtils_addKubernetesInfo(t *testing.T) {
	fields := map[string]interface{}{MetricName(TypePod, CpuTotal): float64(1)}
	tags := map[string]string{
		ContainerNamekey: "testContainer",
		K8sPodNameKey:    "testPod",
		PodIdKey:         "123",
		K8sNamespace:     "testNamespace",
		TypeService:      "testService",
		NodeNameKey:      "testNode",
		Timestamp:        strconv.FormatInt(time.Now().UnixNano(), 10),
	}
	md := ConvertToOTLPMetrics(fields, tags, zap.NewNop())
	rm := md.ResourceMetrics().At(0)
	attributes := rm.Resource().Attributes()

	kubernetesBlob := map[string]interface{}{}
	AddKubernetesInfo(attributes, kubernetesBlob)
	assert.Equal(t, "", GetStringValFromAttributes(ContainerNamekey, attributes))
	assert.Equal(t, "", GetStringValFromAttributes(K8sPodNameKey, attributes))
	assert.Equal(t, "", GetStringValFromAttributes(PodIdKey, attributes))
	assert.Equal(t, "testNamespace", GetStringValFromAttributes(K8sNamespace, attributes))
	assert.Equal(t, "testService", GetStringValFromAttributes(TypeService, attributes))
	assert.Equal(t, "testNode", GetStringValFromAttributes(NodeNameKey, attributes))

	expectedKubeBlob := map[string]interface{}{"container_name": "testContainer", "host": "testNode", "namespace_name": "testNamespace", "pod_id": "123", "pod_name": "testPod", "service_name": "testService"}
	assert.Equal(t, expectedKubeBlob, kubernetesBlob)
}
