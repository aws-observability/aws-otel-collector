module github.com/aws-observability/aws-otel-collector

go 1.20

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/confmap/provider/s3provider v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awscloudwatchlogsexporter v0.93.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logzioexporter v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/awsproxy v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/sigv4authextension v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/cumulativetodeltaprocessor v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatorateprocessor v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbytraceprocessor v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricsgenerationprocessor v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanprocessor v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filelogreceiver v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver v0.92.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver v0.92.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.8.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/collector/component v0.93.0
	go.opentelemetry.io/collector/confmap v0.93.0
	go.opentelemetry.io/collector/exporter v0.93.0
	go.opentelemetry.io/collector/exporter/loggingexporter v0.92.0
	go.opentelemetry.io/collector/exporter/otlpexporter v0.92.0
	go.opentelemetry.io/collector/exporter/otlphttpexporter v0.92.0
	go.opentelemetry.io/collector/extension v0.93.0
	go.opentelemetry.io/collector/extension/ballastextension v0.92.0
	go.opentelemetry.io/collector/extension/zpagesextension v0.92.0
	go.opentelemetry.io/collector/featuregate v1.0.1
	go.opentelemetry.io/collector/otelcol v0.92.0
	go.opentelemetry.io/collector/processor v0.92.0
	go.opentelemetry.io/collector/processor/batchprocessor v0.92.0
	go.opentelemetry.io/collector/processor/memorylimiterprocessor v0.92.0
	go.opentelemetry.io/collector/receiver v0.93.0
	go.opentelemetry.io/collector/receiver/otlpreceiver v0.92.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.26.0
	golang.org/x/sys v0.16.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require (
	cloud.google.com/go/compute v1.23.3 // indirect
	cloud.google.com/go/compute/metadata v0.2.4-0.20230617002413-005d2dfb6b68 // indirect
	contrib.go.opencensus.io/exporter/prometheus v0.4.2 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.8.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.4.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.3.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4 v4.2.1 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2 v2.2.1 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.1.1 // indirect
	github.com/DataDog/agent-payload/v5 v5.0.102 // indirect
	github.com/DataDog/datadog-agent/pkg/obfuscate v0.50.2 // indirect
	github.com/DataDog/datadog-agent/pkg/proto v0.50.2 // indirect
	github.com/DataDog/datadog-agent/pkg/remoteconfig/state v0.50.2 // indirect
	github.com/DataDog/datadog-agent/pkg/trace v0.50.2 // indirect
	github.com/DataDog/datadog-agent/pkg/util/cgroups v0.50.2 // indirect
	github.com/DataDog/datadog-agent/pkg/util/log v0.50.2 // indirect
	github.com/DataDog/datadog-agent/pkg/util/pointer v0.50.2 // indirect
	github.com/DataDog/datadog-agent/pkg/util/scrubber v0.50.2 // indirect
	github.com/DataDog/datadog-api-client-go/v2 v2.20.0 // indirect
	github.com/DataDog/datadog-go/v5 v5.1.1 // indirect
	github.com/DataDog/go-sqllexer v0.0.8 // indirect
	github.com/DataDog/go-tuf v1.0.2-0.5.2 // indirect
	github.com/DataDog/gohai v0.0.0-20230524154621-4316413895ee // indirect
	github.com/DataDog/opentelemetry-mapping-go/pkg/inframetadata v0.11.0 // indirect
	github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/attributes v0.11.0 // indirect
	github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/logs v0.11.0 // indirect
	github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/metrics v0.11.0 // indirect
	github.com/DataDog/opentelemetry-mapping-go/pkg/quantile v0.11.0 // indirect
	github.com/DataDog/sketches-go v1.4.4 // indirect
	github.com/DataDog/zstd v1.5.2 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp v1.21.0 // indirect
	github.com/IBM/sarama v1.42.1 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/Showmax/go-fqdn v1.0.0 // indirect
	github.com/alecthomas/participle/v2 v2.1.1 // indirect
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/apache/thrift v0.19.0 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/aws/aws-sdk-go v1.50.2 // indirect
	github.com/aws/aws-sdk-go-v2 v1.24.1 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.5.4 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.26.3 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.16.14 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.14.11 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.7.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.2.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.10.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.2.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.10.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.16.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.48.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.18.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.21.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.26.7 // indirect
	github.com/aws/smithy-go v1.19.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/bmatcuk/doublestar/v4 v4.6.1 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/checkpoint-restore/go-criu/v5 v5.3.0 // indirect
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575 // indirect
	github.com/cilium/ebpf v0.9.1 // indirect
	github.com/cncf/xds/go v0.0.0-20230607035331-e9ce68804cb4 // indirect
	github.com/containerd/cgroups/v3 v3.0.2 // indirect
	github.com/containerd/console v1.0.3 // indirect
	github.com/containerd/ttrpc v1.2.2 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/cyphar/filepath-securejoin v0.2.4 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dennwc/varint v1.0.0 // indirect
	github.com/digitalocean/godo v1.104.1 // indirect
	github.com/docker/distribution v2.8.2+incompatible // indirect
	github.com/docker/docker v24.0.7+incompatible // indirect
	github.com/docker/go-connections v0.5.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/eapache/go-resiliency v1.4.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/envoyproxy/go-control-plane v0.11.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.2 // indirect
	github.com/euank/go-kmsg-parser v2.0.0+incompatible // indirect
	github.com/expr-lang/expr v1.15.8 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/frankban/quicktest v1.14.3 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-openapi/jsonpointer v0.20.0 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	github.com/go-zookeeper/zk v1.0.3 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/godbus/dbus/v5 v5.0.6 // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.0.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/cadvisor v0.48.1 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.0 // indirect
	github.com/gophercloud/gophercloud v1.7.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grafana/regexp v0.0.0-20221122212121-6b5c0a4cb7fd // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.0 // indirect
	github.com/hashicorp/consul/api v1.26.1 // indirect
	github.com/hashicorp/cronexpr v1.1.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.6.2 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.4 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/hashicorp/nomad/api v0.0.0-20230721134942-515895c7690c // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/hetznercloud/hcloud-go/v2 v2.4.0 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/influxdata/go-syslog/v3 v3.0.1-0.20230911200830-875f5bc594a4 // indirect
	github.com/ionos-cloud/sdk-go/v6 v6.1.9 // indirect
	github.com/jaegertracing/jaeger v1.53.0 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/karrick/godirwalk v1.17.0 // indirect
	github.com/klauspost/compress v1.17.4 // indirect
	github.com/knadh/koanf v1.5.0 // indirect
	github.com/knadh/koanf/v2 v2.0.1 // indirect
	github.com/kolo/xmlrpc v0.0.0-20220921171641-a4b6fa1dd06b // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/leodido/ragel-machinery v0.0.0-20181214104525-299bdde78165 // indirect
	github.com/lightstep/go-expohisto v1.0.0 // indirect
	github.com/linode/linodego v1.23.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20220913051719-115f729f3c8c // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/miekg/dns v1.1.56 // indirect
	github.com/mistifyio/go-zfs v2.1.2-0.20190413222219-f784269be439+incompatible // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/hashstructure/v2 v2.0.2 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20231216201459-8508981c8b6c // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/sys/mountinfo v0.6.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/go-grpc-compression v1.2.2 // indirect
	github.com/mrunalp/fileutils v0.5.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil v0.93.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/containerinsight v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/cwlogs v0.93.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/ecsutil v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/k8s v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/proxy v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.93.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/kafka v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/kubelet v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchpersignal v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.93.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/azure v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheus v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/signalfx v0.92.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin v0.92.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0-rc5 // indirect
	github.com/opencontainers/runc v1.1.9 // indirect
	github.com/opencontainers/runtime-spec v1.1.0-rc.3 // indirect
	github.com/opencontainers/selinux v1.10.0 // indirect
	github.com/openshift/api v3.9.0+incompatible // indirect
	github.com/openshift/client-go v0.0.0-20210521082421-73d9475a9142 // indirect
	github.com/openzipkin/zipkin-go v0.4.2 // indirect
	github.com/outcaste-io/ristretto v0.2.1 // indirect
	github.com/ovh/go-ovh v1.4.3 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.18 // indirect
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20220216144756-c35f1ee13d7c // indirect
	github.com/prometheus/client_golang v1.18.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.46.0 // indirect
	github.com/prometheus/common/sigv4 v0.1.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/prometheus/prometheus v0.48.1 // indirect
	github.com/prometheus/statsd_exporter v0.22.7 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/relvacode/iso8601 v1.3.0 // indirect
	github.com/rs/cors v1.10.1 // indirect
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.21 // indirect
	github.com/seccomp/libseccomp-golang v0.9.2-0.20220502022130-f33da4d89646 // indirect
	github.com/secure-systems-lab/go-securesystemslib v0.7.0 // indirect
	github.com/shirou/gopsutil/v3 v3.23.12 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/signalfx/com_signalfx_metrics_protobuf v0.0.3 // indirect
	github.com/signalfx/sapm-proto v0.14.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/syndtr/gocapability v0.0.0-20200815063812-42c35b437635 // indirect
	github.com/tidwall/gjson v1.12.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/tinylru v1.1.0 // indirect
	github.com/tidwall/wal v1.1.7 // indirect
	github.com/tinylib/msgp v1.1.9 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/vishvananda/netlink v1.1.1-0.20210330154013-f5de75959ad5 // indirect
	github.com/vishvananda/netns v0.0.0-20210104183010-2eb08e3e575f // indirect
	github.com/vultr/govultr/v2 v2.17.2 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	go.etcd.io/bbolt v1.3.8 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector v0.93.0 // indirect
	go.opentelemetry.io/collector/config/configauth v0.92.0 // indirect
	go.opentelemetry.io/collector/config/configcompression v0.92.0 // indirect
	go.opentelemetry.io/collector/config/configgrpc v0.92.0 // indirect
	go.opentelemetry.io/collector/config/confighttp v0.92.0 // indirect
	go.opentelemetry.io/collector/config/confignet v0.92.0 // indirect
	go.opentelemetry.io/collector/config/configopaque v0.92.0 // indirect
	go.opentelemetry.io/collector/config/configretry v0.93.0 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.93.0 // indirect
	go.opentelemetry.io/collector/config/configtls v0.92.0 // indirect
	go.opentelemetry.io/collector/config/internal v0.92.0 // indirect
	go.opentelemetry.io/collector/connector v0.92.0 // indirect
	go.opentelemetry.io/collector/consumer v0.93.0 // indirect
	go.opentelemetry.io/collector/extension/auth v0.92.0 // indirect
	go.opentelemetry.io/collector/pdata v1.0.1 // indirect
	go.opentelemetry.io/collector/semconv v0.93.0 // indirect
	go.opentelemetry.io/collector/service v0.92.0 // indirect
	go.opentelemetry.io/contrib/config v0.2.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.47.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.46.1 // indirect
	go.opentelemetry.io/contrib/propagators/b3 v1.21.1 // indirect
	go.opentelemetry.io/contrib/zpages v0.46.1 // indirect
	go.opentelemetry.io/otel v1.22.0 // indirect
	go.opentelemetry.io/otel/bridge/opencensus v0.44.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.44.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v0.44.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.22.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.22.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.22.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.45.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v0.44.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.22.0 // indirect
	go.opentelemetry.io/otel/metric v1.22.0 // indirect
	go.opentelemetry.io/otel/sdk v1.22.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.22.0 // indirect
	go.opentelemetry.io/otel/trace v1.22.0 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/exp v0.0.0-20240103183307-be819d1f06fc // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/oauth2 v0.16.0 // indirect
	golang.org/x/term v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.4.0 // indirect
	golang.org/x/tools v0.16.1 // indirect
	gonum.org/v1/gonum v0.14.0 // indirect
	google.golang.org/api v0.150.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20240102182953-50ed04b92917 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240102182953-50ed04b92917 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240108191215-35c7eff3a6b1 // indirect
	google.golang.org/grpc v1.60.1 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gopkg.in/zorkian/go-datadog-api.v2 v2.30.0 // indirect
	k8s.io/api v0.28.4 // indirect
	k8s.io/apimachinery v0.28.4 // indirect
	k8s.io/client-go v0.28.4 // indirect
	k8s.io/klog v1.0.0 // indirect
	k8s.io/klog/v2 v2.110.1 // indirect
	k8s.io/kube-openapi v0.0.0-20230717233707-2695361300d9 // indirect
	k8s.io/utils v0.0.0-20240102154912-e7106e64919e // indirect
	sigs.k8s.io/controller-runtime v0.16.3 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.3.0 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/4433
exclude github.com/StackExchange/wmi v1.2.0

// https://github.com/google/flatbuffers/issues/6466
exclude github.com/google/flatbuffers v1.12.0

// see https://github.com/aws-observability/aws-otel-collector/issues/977
exclude github.com/docker/distribution v2.8.0+incompatible

replace github.com/go-chi/chi/v4 => github.com/go-chi/chi v4.0.0+incompatible

// see https://github.com/ionos-cloud/sdk-go/issues/27
exclude github.com/ionos-cloud/sdk-go/v6 v6.0.5851

// https://github.com/go-openapi/spec/issues/156
replace github.com/go-openapi/spec v0.20.5 => github.com/go-openapi/spec v0.20.6

// This replace can be removed once it is patched to v0.3.2 upstream
// see https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/exporter/datadogexporter/go.mod#L116
// Also see the source version from the datadog agent
// https://github.com/DataDog/datadog-agent/blob/main/go.mod#L26
replace github.com/theupdateframework/go-tuf v0.3.1 => github.com/DataDog/go-tuf v0.3.0--fix-localmeta

// excluded as DataDog/agent-payload/v5 v5.0.59 has been removed from source directory but is still present in proxy package
exclude github.com/DataDog/agent-payload/v5 v5.0.59 // indirect

// It appears that the v0.2.0 tag was modified.  Replacing with v0.2.1
replace github.com/outcaste-io/ristretto v0.2.0 => github.com/outcaste-io/ristretto v0.2.1

// openshift removed all tags from their repo, use the pseudoversion from the release-3.9 branch HEAD
replace github.com/openshift/api v3.9.0+incompatible => github.com/openshift/api v0.0.0-20180801171038-322a19404e37

replace github.com/cyphar/filepath-securejoin v0.2.3 => github.com/cyphar/filepath-securejoin v0.2.4

replace github.com/sirupsen/logrus v1.9.0 => github.com/sirupsen/logrus v1.9.3
