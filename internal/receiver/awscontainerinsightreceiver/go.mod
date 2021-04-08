module github.com/aws-observability/aws-otel-collector/internal/receiver/awscontainerinsightreceiver

go 1.14

require (
	github.com/aws-observability/aws-otel-collector/internal/aws v0.0.0-00010101000000-000000000000
	github.com/aws/aws-sdk-go v1.37.26
	github.com/google/cadvisor v0.39.0
	github.com/onsi/ginkgo v1.14.1 // indirect
	github.com/onsi/gomega v1.10.2 // indirect
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.24.0
	go.uber.org/zap v1.16.0
	gopkg.in/ini.v1 v1.57.0 // indirect
	k8s.io/api v0.20.4
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v0.20.4
	k8s.io/klog v1.0.0
)

replace github.com/aws-observability/aws-otel-collector/internal/aws => ../../aws
