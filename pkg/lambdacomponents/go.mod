module github.com/aws-observability/aws-otel-collector/pkg/lambdacomponents

go 1.14

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.29.1-0.20210630203112-81d57601b1bc
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.29.1-0.20210630203112-81d57601b1bc
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.29.1-0.20210630003519-14d917479ef3
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray => github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray v0.29.1-0.20210630203112-81d57601b1bc

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics => github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics v0.29.1-0.20210630203112-81d57601b1bc

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil => github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil v0.29.1-0.20210630203112-81d57601b1bc
