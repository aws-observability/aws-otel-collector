module github.com/aws-observability/aws-otel-collector/pkg/lambdacomponents

go 1.17

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.51.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsprometheusremotewriteexporter v0.51.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.51.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter v0.53.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter v0.51.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/sigv4authextension v0.51.0
	github.com/stretchr/testify v1.7.2
	go.opentelemetry.io/collector v0.53.0
	go.uber.org/multierr v1.8.0
)

require (
	github.com/aws/aws-sdk-go v1.44.28 // indirect
	github.com/aws/aws-sdk-go-v2 v1.9.2 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.8.3 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.4.3 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.6.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.2.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.3.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.4.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.7.2 // indirect
	github.com/aws/smithy-go v1.8.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/distribution v2.8.0-beta.1+incompatible // indirect
	github.com/docker/go-connections v0.4.1-0.20210727194412-58542c764a11 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.6 // indirect
	github.com/knadh/koanf v1.4.1 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/go-grpc-compression v1.1.16 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil v0.51.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/cwlogs v0.51.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics v0.51.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray v0.51.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry v0.53.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite v0.51.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.12.2 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.34.0 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/prometheus/prometheus v0.36.0 // indirect
	github.com/rs/cors v1.8.2 // indirect
	github.com/tidwall/gjson v1.10.2 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/tinylru v1.1.0 // indirect
	github.com/tidwall/wal v1.1.7 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.opentelemetry.io/collector/pdata v0.53.0 // indirect
	go.opentelemetry.io/collector/semconv v0.53.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.32.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.32.0 // indirect
	go.opentelemetry.io/otel v1.7.0 // indirect
	go.opentelemetry.io/otel/metric v0.30.0 // indirect
	go.opentelemetry.io/otel/trace v1.7.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/net v0.0.0-20220526153639-5463443f8c37 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220527130721-00d5c0f3be58 // indirect
	google.golang.org/grpc v1.47.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// see https://github.com/aws-observability/aws-otel-collector/issues/977
exclude github.com/docker/distribution v2.8.0+incompatible
