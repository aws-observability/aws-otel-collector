module github.com/aws-observability/aws-otel-collector

go 1.14

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.16.1-0.20201209014323-5008dc15f83a
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsprometheusremotewriteexporter v0.16.1-0.20201207164444-473669511a75
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.16.1-0.20201209014323-5008dc15f83a
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor v0.16.1-0.20201207164444-473669511a75
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver v0.16.1-0.20201207164444-473669511a75
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver v0.16.1-0.20201207164444-473669511a75
	github.com/opencontainers/runc v1.0.0-rc92
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/collector v0.16.1-0.20201207152538-326931de8c32
	go.uber.org/zap v1.16.0
	golang.org/x/sys v0.0.0-20201015000850-e3ed0017c211
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray => github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray v0.16.0

replace github.com/aws-observability/aws-otel-collector/pkg/lambdacomponents => ./pkg/lambdacomponents
