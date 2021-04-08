module github.com/aws-observability/aws-otel-collector

go 1.14

require (
	github.com/crossdock/crossdock-go v0.0.0-20160816171116-049aabb0122b
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.24.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsprometheusremotewriteexporter v0.24.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.24.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter v0.24.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dynatraceexporter v0.24.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logzioexporter v0.24.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/newrelicexporter v0.24.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter v0.24.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter v0.24.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver v0.24.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor v0.24.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor v0.24.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver v0.24.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver v0.24.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver v0.24.0
	github.com/opencontainers/runc v1.0.0-rc92
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.24.0
	go.uber.org/zap v1.16.0
	golang.org/x/sys v0.0.0-20210315160823-c6e025ad8005
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray => github.com/open-telemetry/opentelemetry-collector-contrib/internal/awsxray v0.24.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk => github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk v0.24.0

replace github.com/aws-observability/aws-otel-collector/pkg/lambdacomponents => ./pkg/lambdacomponents

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.24.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr => github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr v0.24.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata => github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata v0.24.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws => github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws v0.24.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver => ./internal/extension/ecsobserver
