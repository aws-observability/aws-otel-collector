module github.com/aws-observability/aws-otel-collector/pkg/lambdacomponents

go 1.17

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.41.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsprometheusremotewriteexporter v0.41.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.41.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter v0.41.0
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.41.0
	go.uber.org/multierr v1.7.0
)

require (
	github.com/aws/aws-sdk-go v1.42.21 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.1.2 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/containerd/containerd v1.5.9 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/knadh/koanf v1.3.3 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/mostynb/go-grpc-compression v1.1.15 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter v0.41.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil v0.41.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics v0.41.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray v0.41.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry v0.41.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/prometheus/prometheus v1.8.2-0.20210621150501-ff58416a0b02 // indirect
	github.com/rs/cors v1.8.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/tidwall/gjson v1.10.2 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/tinylru v1.1.0 // indirect
	github.com/tidwall/wal v1.1.7 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.opentelemetry.io/collector/model v0.41.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.27.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.27.0 // indirect
	go.opentelemetry.io/otel v1.2.0 // indirect
	go.opentelemetry.io/otel/internal/metric v0.25.0 // indirect
	go.opentelemetry.io/otel/metric v0.25.0 // indirect
	go.opentelemetry.io/otel/trace v1.2.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/zap v1.19.1 // indirect
	golang.org/x/net v0.0.0-20210917221730-978cfadd31cf // indirect
	golang.org/x/sys v0.0.0-20211124211545-fe61309f8881 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211206160659-862468c7d6e0 // indirect
	google.golang.org/grpc v1.42.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray => github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray v0.41.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics => github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics v0.41.0

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil => github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil v0.41.0
