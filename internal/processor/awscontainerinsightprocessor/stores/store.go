package stores

import "go.opentelemetry.io/collector/consumer/pdata"

type K8sStore interface {
	Decorate(metric pdata.ResourceMetrics, kubernetesBlob map[string]interface{}) bool
	RefreshTick()
}
