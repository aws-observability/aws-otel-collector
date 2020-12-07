module github.com/aws-observability/aws-otel-collector

go 1.14

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.16.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsprometheusremotewriteexporter v0.16.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.16.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter v0.16.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dynatraceexporter v0.16.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/newrelicexporter v0.16.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter v0.16.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter v0.16.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor v0.16.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver v0.16.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver v0.16.0
	github.com/opencontainers/runc v1.0.0-rc92
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/collector v0.16.0
	go.uber.org/zap v1.16.0
	golang.org/x/sys v0.0.0-20201015000850-e3ed0017c211
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray => github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray v0.16.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk => github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk v0.16.0

replace github.com/aws-observability/aws-otel-collector/pkg/lambdacomponents => ./pkg/lambdacomponents

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver => github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver v0.16.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig => github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig v0.16.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.16.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr => github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr v0.16.0
