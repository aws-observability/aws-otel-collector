module github.com/aws-observability/aws-otel-collector/internal/aws

go 1.14

require (
	github.com/aws/aws-sdk-go v1.37.23
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.23.0
	go.opentelemetry.io/otel v0.19.0
	go.uber.org/zap v1.16.0
	k8s.io/api v0.20.4
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v0.20.4
)
