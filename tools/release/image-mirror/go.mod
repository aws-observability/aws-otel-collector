module github.com/aws-observability/aws-otel-collector/tools/release/adot-operator-images-mirror

go 1.21

require (
	github.com/aws/aws-sdk-go-v2/config v1.27.18
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.25.7
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/docker/docker v27.1.1+incompatible
	github.com/google/go-containerregistry v0.19.1
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/aws/aws-sdk-go-v2 v1.30.5 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.18 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.5 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.17 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.17 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.11.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.11.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.20.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.24.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.28.12 // indirect
	github.com/aws/smithy-go v1.20.4 // indirect
	github.com/containerd/log v0.1.0 // indirect
	github.com/containerd/stargz-snapshotter/estargz v0.15.1 // indirect
	github.com/distribution/reference v0.6.0 // indirect
	github.com/docker/cli v26.1.4+incompatible // indirect
	github.com/docker/distribution v2.8.3+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.8.2 // indirect
	github.com/docker/go-connections v0.5.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/moby/docker-image-spec v1.3.1 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/vbatts/tar-split v0.11.5 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.52.0 // indirect
	go.opentelemetry.io/otel v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.22.0 // indirect
	go.opentelemetry.io/otel/metric v1.27.0 // indirect
	go.opentelemetry.io/otel/sdk v1.22.0 // indirect
	go.opentelemetry.io/otel/trace v1.27.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/time v0.0.0-20211116232009-f0f3c7e86c11 // indirect
)

// see https://github.com/aws-observability/aws-otel-collector/issues/977
exclude github.com/docker/distribution v2.8.0+incompatible

replace github.com/docker/docker v25.0.3+incompatible => github.com/docker/docker v25.0.6+incompatible
