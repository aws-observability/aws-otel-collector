module  github.com/aws-observability/aws-otel-collector/pkg/lambdacomponents

go 1.14

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.16.1-0.20201209014323-5008dc15f83a
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.16.1-0.20201209014323-5008dc15f83a
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/collector v0.16.1-0.20201207152538-326931de8c32
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray => github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray v0.16.0
