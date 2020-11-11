module aws-observability.io/collector

go 1.14

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.14.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.14.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor v0.14.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver v0.14.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver v0.14.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/collector v0.14.0
	go.uber.org/zap v1.16.0
	golang.org/x/sys v0.0.0-20200821140526-fda516888d29
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray => github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray v0.14.0
