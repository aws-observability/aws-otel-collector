module aws-observability.io/collector/lambdacomponents

go 1.15

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.15.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.15.0
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/collector v0.15.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray => github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray v0.14.1-0.20201111210848-994cabe5d596
