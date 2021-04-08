package structuredlogsadapter

import (
	"encoding/json"

	common "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	"github.com/aws-observability/aws-otel-collector/internal/processor/awscontainerinsightprocessor/stores"
	"go.opentelemetry.io/collector/consumer/pdata"
)

func TagMetricSource(attributes pdata.AttributeMap) {
	metricType := stores.GetStringValFromAttributes(common.MetricType, attributes)
	if metricType == "" {
		return
	}

	var sources []string
	switch metricType {
	case common.TypeNode:
		sources = append(sources, []string{"cadvisor", "/proc", "pod", "calculated"}...)
	case common.TypeNodeFS:
		sources = append(sources, []string{"cadvisor", "calculated"}...)
	case common.TypeNodeNet:
		sources = append(sources, []string{"cadvisor", "calculated"}...)
	case common.TypeNodeDiskIO:
		sources = append(sources, []string{"cadvisor"}...)
	case common.TypePod:
		sources = append(sources, []string{"cadvisor", "pod", "calculated"}...)
	case common.TypePodNet:
		sources = append(sources, []string{"cadvisor", "calculated"}...)
	case common.TypeContainer:
		sources = append(sources, []string{"cadvisor", "pod", "calculated"}...)
	case common.TypeContainerFS:
		sources = append(sources, []string{"cadvisor", "calculated"}...)
	case common.TypeContainerDiskIO:
		sources = append(sources, []string{"cadvisor"}...)
	case common.TypeCluster, common.TypeClusterService, common.TypeClusterNamespace:
		sources = append(sources, []string{"apiserver"}...)
	}

	if len(sources) > 0 {
		sourcesInfo, err := json.Marshal(sources)
		if err != nil {
			return
		}
		stores.AddAttribute(common.SourcesKey, string(sourcesInfo), attributes)
	}
}

func AddKubernetesInfo(attributes pdata.AttributeMap, kubernetesBlob map[string]interface{}) {
	needMoveToKubernetes := map[string]string{common.ContainerNamekey: "container_name", common.K8sPodNameKey: "pod_name",
		common.PodIdKey: "pod_id"}
	needCopyToKubernetes := map[string]string{common.K8sNamespace: "namespace_name", common.TypeService: "service_name", common.NodeNameKey: "host"}

	for k, v := range needMoveToKubernetes {
		if attVal := stores.GetStringValFromAttributes(k, attributes); attVal != "" {
			kubernetesBlob[v] = attVal
			stores.RemoveAttribute(k, attributes)
		}
	}
	for k, v := range needCopyToKubernetes {
		if attVal := stores.GetStringValFromAttributes(k, attributes); attVal != "" {
			kubernetesBlob[v] = attVal
		}
	}

	if len(kubernetesBlob) > 0 {
		kubernetesInfo, err := json.Marshal(kubernetesBlob)
		if err != nil {
			return
		}
		stores.AddAttribute(common.Kubernetes, string(kubernetesInfo), attributes)
	}
}
