module github.com/aws-observability/aws-otel-collector

go 1.14

require (
	github.com/aws-observability/aws-otel-collector/internal/processor/awscontainerinsightprocessor v0.0.0-00010101000000-000000000000
	github.com/aws-observability/aws-otel-collector/internal/receiver/awscontainerinsightreceiver v0.0.0-00010101000000-000000000000
	github.com/crossdock/crossdock-go v0.0.0-20160816171116-049aabb0122b
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.24.1-0.20210415165634-9a981e34f21d
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsprometheusremotewriteexporter v0.24.1-0.20210415165634-9a981e34f21d
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.24.1-0.20210415165634-9a981e34f21d
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter v0.24.1-0.20210415165634-9a981e34f21d
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dynatraceexporter v0.24.1-0.20210415165634-9a981e34f21d
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logzioexporter v0.24.1-0.20210415165634-9a981e34f21d
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/newrelicexporter v0.24.1-0.20210415165634-9a981e34f21d
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter v0.24.1-0.20210415165634-9a981e34f21d
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter v0.24.1-0.20210415165634-9a981e34f21d
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor v0.24.1-0.20210415165634-9a981e34f21d
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor v0.24.1-0.20210415165634-9a981e34f21d
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver v0.24.1-0.20210415165634-9a981e34f21d
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver v0.24.1-0.20210415165634-9a981e34f21d
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver v0.24.1-0.20210415165634-9a981e34f21d
	github.com/opencontainers/runc v1.0.0-rc93
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.24.1-0.20210414213607-0f75efa9f327
	go.uber.org/zap v1.16.0
	golang.org/x/sys v0.0.0-20210320140829-1e4c9ba3b0c4
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray => github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray v0.0.0-20210415165634-9a981e34f21d

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk => github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk v0.24.1-0.20210415165634-9a981e34f21d

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics => github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics v0.0.0-20210415165634-9a981e34f21d

replace github.com/aws-observability/aws-otel-collector/pkg/lambdacomponents => ./pkg/lambdacomponents

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.24.1-0.20210415165634-9a981e34f21d

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr => github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr v0.24.1-0.20210415165634-9a981e34f21d

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata => github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata v0.24.1-0.20210415165634-9a981e34f21d

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws => github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws v0.24.1-0.20210415165634-9a981e34f21d

replace github.com/aws-observability/aws-otel-collector/internal/aws => ./internal/aws

replace github.com/aws-observability/aws-otel-collector/internal/receiver/awscontainerinsightreceiver => ./internal/receiver/awscontainerinsightreceiver

replace github.com/aws-observability/aws-otel-collector/internal/processor/awscontainerinsightprocessor => ./internal/processor/awscontainerinsightprocessor

// replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter => ../opentelemetry-collector-contrib/exporter/awsemfexporter
// replace go.opentelemetry.io/collector => go.opentelemetry.io/collector v0.23.0
