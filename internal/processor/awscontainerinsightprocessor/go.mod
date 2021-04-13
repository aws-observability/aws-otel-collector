module github.com/aws-observability/aws-otel-collector/internal/processor/awscontainerinsightprocessor

go 1.14

require (
	github.com/aws-observability/aws-otel-collector/internal/aws v0.0.0-00010101000000-000000000000
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.24.0
	go.uber.org/zap v1.16.0
	gopkg.in/ini.v1 v1.57.0 // indirect
	k8s.io/api v0.20.4
)

replace github.com/aws-observability/aws-otel-collector/internal/aws => ../../aws
