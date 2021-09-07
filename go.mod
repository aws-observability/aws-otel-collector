module github.com/aws-observability/aws-otel-collector

go 1.14

require (
	github.com/crossdock/crossdock-go v0.0.0-20160816171116-049aabb0122b
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsprometheusremotewriteexporter v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dynatraceexporter v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logzioexporter v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/newrelicexporter v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanprocessor v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver v0.34.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver v0.34.0
	github.com/opencontainers/runc v1.0.1
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.34.0
	go.uber.org/zap v1.19.0
	golang.org/x/sys v0.0.0-20210806184541-e5e7981a1069
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray => github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray v0.34.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk => github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk v0.34.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics => github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics v0.34.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil => github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil v0.34.0

replace github.com/aws-observability/aws-otel-collector/pkg/lambdacomponents => ./pkg/lambdacomponents

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.34.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr => github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr v0.34.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata => github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata v0.34.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/containerinsight => github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/containerinsight v0.34.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/k8s => github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/k8s v0.34.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig => github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig v0.34.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/kubelet => github.com/open-telemetry/opentelemetry-collector-contrib/internal/kubelet v0.34.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/proxy => github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/proxy v0.34.0

// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/4433
exclude github.com/StackExchange/wmi v1.2.0

// https://github.com/google/flatbuffers/issues/6466
exclude github.com/google/flatbuffers v1.12.0
