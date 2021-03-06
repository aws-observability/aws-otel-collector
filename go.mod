module github.com/aws-observability/aws-otel-collector

go 1.14

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.21.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsprometheusremotewriteexporter v0.21.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.21.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter v0.21.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dynatraceexporter v0.21.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/newrelicexporter v0.21.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter v0.21.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter v0.21.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor v0.21.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver v0.21.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver v0.21.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver v0.21.1-0.20210302030508-0050d3df8ce5
	github.com/opencontainers/runc v1.0.0-rc92
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.21.1-0.20210225192722-e6319ac4c6fc
	go.uber.org/zap v1.16.0
	golang.org/x/sys v0.0.0-20201214210602-f9fddec55a1e
	gopkg.in/jcmturner/gokrb5.v7 v7.5.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray => github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray v0.21.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk => github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk v0.21.0

replace github.com/aws-observability/aws-otel-collector/pkg/lambdacomponents => ./pkg/lambdacomponents

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.21.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr => github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr v0.21.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata => github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata v0.21.0
