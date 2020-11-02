module aws-observability.io/collector

go 1.14

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.13.1-0.20201102154212-c742efac0db4
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.13.1-0.20201102154212-c742efac0db4
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor v0.12.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver v0.13.1-0.20201022225622-7a23d7606a89
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/collector v0.13.1-0.20201101004512-f4e4382d0e0e
	go.uber.org/zap v1.16.0
	golang.org/x/sys v0.0.0-20200821140526-fda516888d29
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray => github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray v0.13.1-0.20201102154212-c742efac0db4
