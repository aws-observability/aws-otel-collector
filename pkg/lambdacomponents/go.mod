module github.com/aws-observability/aws-otel-collector/pkg/lambdacomponents

go 1.14

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.19.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.19.0
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.19.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray => github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray v0.19.0
