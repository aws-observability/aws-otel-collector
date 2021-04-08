package k8sapiserver

import (
	"fmt"
	"os"
	"testing"

	common "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	"github.com/aws-observability/aws-otel-collector/internal/aws/k8scommon/k8sclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"
)

var mockClient = new(MockClient)

var mockK8sClient = &k8sclient.K8sClient{
	Pod:  mockClient,
	Node: mockClient,
	Ep:   mockClient,
}

func mockGet() *k8sclient.K8sClient {
	return mockK8sClient
}

type MockClient struct {
	k8sclient.PodClient
	k8sclient.NodeClient
	k8sclient.EpClient

	mock.Mock
}

// k8sclient.PodClient
func (client *MockClient) NamespaceToRunningPodNum() map[string]int {
	args := client.Called()
	return args.Get(0).(map[string]int)
}

// k8sclient.NodeClient
func (client *MockClient) ClusterFailedNodeCount() int {
	args := client.Called()
	return args.Get(0).(int)
}

func (client *MockClient) ClusterNodeCount() int {
	args := client.Called()
	return args.Get(0).(int)
}

// k8sclient.EpClient
func (client *MockClient) ServiceToPodNum() map[k8sclient.Service]int {
	args := client.Called()
	return args.Get(0).(map[k8sclient.Service]int)
}

func (client *MockClient) Init() {
}

func (client *MockClient) Shutdown() {
}

func getStringAttrVal(m pdata.Metrics, key string) string {
	rm := m.ResourceMetrics().At(0)
	attributes := rm.Resource().Attributes()
	if attributeValue, ok := attributes.Get(key); ok {
		return attributeValue.StringVal()
	} else {
		return ""
	}
}

func assertMetricValueEqual(t *testing.T, m pdata.Metrics, metricName string, expected int64) {
	rm := m.ResourceMetrics().At(0)
	ilms := rm.InstrumentationLibraryMetrics()

	for j := 0; j < ilms.Len(); j++ {
		metricSlice := ilms.At(j).Metrics()
		for i := 0; i < metricSlice.Len(); i++ {
			metric := metricSlice.At(i)
			if metric.Name() == metricName {
				if metric.DataType() == pdata.MetricDataTypeIntGauge {
					assert.Equal(t, expected, metric.IntGauge().DataPoints().At(0).Value())
					return
				} else {
					msg := fmt.Sprintf("Metric with name: %v has wrong type.", metricName)
					assert.Fail(t, msg)
				}
			}
		}
	}

	msg := fmt.Sprintf("No metric with name: %v", metricName)
	assert.Fail(t, msg)
}

func TestK8sAPIServer_GetMetrics(t *testing.T) {
	hostName, err := os.Hostname()
	assert.NoError(t, err)
	k8sApiServer := &K8sAPIServer{
		nodeName: hostName,
		leading:  true,
		logger:   zap.NewNop(),
	}

	k8sclient.Get = mockGet

	mockClient.On("NamespaceToRunningPodNum").Return(map[string]int{"default": 2})
	mockClient.On("ClusterFailedNodeCount").Return(1)
	mockClient.On("ClusterNodeCount").Return(1)
	mockClient.On("ServiceToPodNum").Return(map[k8sclient.Service]int{k8sclient.NewService("service1", "kube-system"): 1, k8sclient.NewService("service2", "kube-system"): 1})

	metrics := k8sApiServer.GetMetrics()
	assert.NoError(t, err)

	/*
		tags: map[Timestamp:1557291396709 Type:Cluster], fields: map[cluster_failed_node_count:1 cluster_node_count:1],
		tags: map[Service:service2 Timestamp:1557291396709 Type:ClusterService], fields: map[service_number_of_running_pods:1],
		tags: map[Service:service1 Timestamp:1557291396709 Type:ClusterService], fields: map[service_number_of_running_pods:1],
		tags: map[Namespace:default Timestamp:1557291396709 Type:ClusterNamespace], fields: map[namespace_number_of_running_pods:2],
	*/
	for _, metric := range metrics {

		// log.Printf("measurement: %v, tags: %v, fields: %v, time: %v\n", metric.Measurement, metric.Tags, metric.Fields, metric.Time)
		if metricType := getStringAttrVal(metric, common.MetricType); metricType == common.TypeCluster {
			assertMetricValueEqual(t, metric, "cluster_failed_node_count", int64(1))
			assertMetricValueEqual(t, metric, "cluster_node_count", int64(1))
		} else if metricType == common.TypeClusterService {
			assertMetricValueEqual(t, metric, "service_number_of_running_pods", int64(1))
			if serviceTag := getStringAttrVal(metric, common.TypeService); serviceTag != "service1" && serviceTag != "service2" {
				assert.Fail(t, "Expect to see a tag named as Service")
			}
			if namespaceTag := getStringAttrVal(metric, common.K8sNamespace); namespaceTag != "kube-system" {
				assert.Fail(t, "Expect to see a tag named as Namespace")
			}
		} else if metricType == common.TypeClusterNamespace {
			assertMetricValueEqual(t, metric, "namespace_number_of_running_pods", int64(2))
			assert.Equal(t, "default", getStringAttrVal(metric, common.K8sNamespace))
		} else {
			assert.Fail(t, "Unexpected metric type: "+metricType)
		}
	}

}
