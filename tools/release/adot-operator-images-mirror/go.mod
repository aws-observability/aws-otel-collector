module github.com/aws-observability/aws-otel-collector/tools/release/adot-operator-images-mirror

go 1.19

require (
	github.com/aws/aws-sdk-go-v2/config v1.18.20
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.15.7
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/docker/docker v23.0.3+incompatible
	github.com/google/go-containerregistry v0.13.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/Microsoft/go-winio v0.6.0 // indirect
	github.com/aws/aws-sdk-go-v2 v1.17.8 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.13.19 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.32 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.26 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.33 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.26 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.12.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.14.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.18.8 // indirect
	github.com/aws/smithy-go v1.13.5 // indirect
	github.com/containerd/stargz-snapshotter/estargz v0.12.1 // indirect
	github.com/docker/cli v20.10.20+incompatible // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.7.0 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/klauspost/compress v1.15.11 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0-rc2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/vbatts/tar-split v0.11.2 // indirect
	golang.org/x/mod v0.10.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/time v0.0.0-20211116232009-f0f3c7e86c11 // indirect
	golang.org/x/tools v0.8.0 // indirect
)

// see https://github.com/aws-observability/aws-otel-collector/issues/977
exclude github.com/docker/distribution v2.8.0+incompatible
