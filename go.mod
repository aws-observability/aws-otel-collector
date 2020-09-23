module aws-observability.io/collector

go 1.14

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.0.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.7.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/collector v0.7.0
	go.uber.org/zap v1.15.0
	golang.org/x/sys v0.0.0-20200805065543-0cf7623e9dbd
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter => ./pkg/devexporter/awsemfexporter