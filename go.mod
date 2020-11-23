module aws-observability.io/collector

go 1.14

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.14.1-0.20201117192543-4a81c809e720
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.14.1-0.20201111210848-994cabe5d596
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter v0.14.1-0.20201111210848-994cabe5d596
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dynatraceexporter v0.15.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter v0.14.1-0.20201111210848-994cabe5d596
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor v0.14.1-0.20201111210848-994cabe5d596
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver v0.14.1-0.20201118171217-2398a00e4656
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver v0.14.1-0.20201111210848-994cabe5d596
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/collector v0.15.0
	go.uber.org/zap v1.16.0
	golang.org/x/sys v0.0.0-20201015000850-e3ed0017c211
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray => github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray v0.14.1-0.20201111210848-994cabe5d596

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk => github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk v0.14.1-0.20201111210848-994cabe5d596
