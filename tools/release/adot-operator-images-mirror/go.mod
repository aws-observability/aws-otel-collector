module github.com/aws-observability/aws-otel-collector/tools/release/adot-operator-images-mirror

go 1.17

require (
	github.com/aws/aws-sdk-go-v2/config v1.10.3
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.8.1
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/docker/docker v20.10.11+incompatible
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/Microsoft/go-winio v0.5.0 // indirect
	github.com/Microsoft/hcsshim v0.8.23 // indirect
	github.com/aws/aws-sdk-go-v2 v1.11.1 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.6.3 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.8.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.0.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.5.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.6.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.11.0 // indirect
	github.com/aws/smithy-go v1.9.0 // indirect
	github.com/bits-and-blooms/bitset v1.2.0 // indirect
	github.com/containerd/cgroups v1.0.1 // indirect
	github.com/containerd/containerd v1.5.10 // indirect
	github.com/containerd/continuity v0.1.0 // indirect
	github.com/containerd/fifo v1.0.0 // indirect
	github.com/containerd/ttrpc v1.1.0 // indirect
	github.com/containerd/typeurl v1.0.2 // indirect
	github.com/docker/distribution v2.8.0-beta.1+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-events v0.0.0-20190806004212-e31b211e4f1c // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/gogo/googleapis v1.4.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/klauspost/compress v1.11.13 // indirect
	github.com/moby/locker v1.0.1 // indirect
	github.com/moby/sys/mountinfo v0.4.1 // indirect
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/opencontainers/runc v1.0.2 // indirect
	github.com/opencontainers/runtime-spec v1.0.3-0.20210326190908-1c3f411f0417 // indirect
	github.com/opencontainers/selinux v1.8.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	go.opencensus.io v0.22.3 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a // indirect
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22 // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/time v0.0.0-20211116232009-f0f3c7e86c11 // indirect
	google.golang.org/genproto v0.0.0-20201110150050-8816d57aaa9a // indirect
	google.golang.org/grpc v1.33.2 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

// see https://github.com/aws-observability/aws-otel-collector/issues/977
exclude github.com/docker/distribution v2.8.0+incompatible
