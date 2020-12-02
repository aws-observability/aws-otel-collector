module  github.com/aws-observability/aws-otel-collector/pkg/lambdacomponents

go 1.14

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.15.1-0.20201121135752-0c3b74245b41
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.14.1-0.20201111210848-994cabe5d596
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/collector v0.15.1-0.20201120151746-8ceddba7ea03
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray => github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray v0.14.1-0.20201111210848-994cabe5d596
