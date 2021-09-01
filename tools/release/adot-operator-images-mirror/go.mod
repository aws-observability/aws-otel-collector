module github.com/aws-observability/aws-otel-collector/tools/release/adot-operator-images-mirror

go 1.16

require (
	github.com/aws/aws-sdk-go-v2/config v1.7.0
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.5.0
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/fsouza/go-dockerclient v1.7.4
	gopkg.in/yaml.v2 v2.4.0
)
