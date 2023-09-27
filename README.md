[![C/I](https://github.com/aws-observability/aws-otel-collector/actions/workflows/CI.yml/badge.svg?branch=main)](https://github.com/aws-observability/aws-otel-collector/actions/workflows/CI.yml)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/aws-observability/aws-otel-collector)



### Overview

AWS Distro for OpenTelemetry Collector (ADOT Collector) is an AWS supported version of the upstream OpenTelemetry Collector and is distributed by Amazon. It supports the selected components from the OpenTelemetry community. It is fully compatible with AWS computing platforms including EC2, ECS, and EKS. It enables users to send telemetry data to AWS CloudWatch Metrics, Traces, and Logs backends as well as the other supported backends.

See the [AWS Distro for OpenTelemetry documentation](https://aws-otel.github.io/docs/getting-started/collector) for more information. Additionally, the ADOT Collector is now generally available for metrics.

### Getting Help

Use the community resources below for getting help with the ADOT Collector.
* Open a support ticket with [AWS Support](http://docs.aws.amazon.com/awssupport/latest/user/getting-started.html).
* Use [GitHub issues](https://github.com/aws-observability/aws-otel-collector/issues) to report bugs and request features.
* Join our GitHub [Community](https://github.com/aws-observability/aws-otel-community) for AWS Distro for OpenTelemetry to ask your questions, file issues, or request enhancements.
* If you think you may have found a bug, open a [bug report](https://github.com/aws-observability/aws-otel-collector/issues/new?template=bug_report.md).
* For contributing guidelines, refer to [CONTRIBUTING.md](CONTRIBUTING.md).

### Notice: ADOT Collector v0.33.0 Breaking Change
Users of the `statsd` receiver, please refer to GitHub Issue - [Warning: StatsD Receiver → EMF Exporter Metric Pipeline Breaking Change](https://github.com/aws-observability/aws-otel-collector/issues/2249)
for information on an upcoming breaking change.

### Notices: ADOT Collector v0.35.0 Breaking Changes
* Users of the 'awscontainerinsightreceiver', please refer to the GitHub Issue - [Warning: Container Image Default User Change → Important consideration for AWSContainerInsightReceiver](https://github.com/aws-observability/aws-otel-collector/issues/2317) for more information on an upcoming breaking change.
* Users of the `prometheus` or `prometheusremotewrite` exporters please refer to the GitHub Issue [Warning: ADOT Collector v0.35.0 breaking change](https://github.com/aws-observability/aws-otel-collector/issues/2367)
for information on an upcoming breaking change.

#### ADOT Collector Built-in Components

This table represents the supported components of the ADOT Collector. The highlighted components below are developed by AWS in-house. The rest of the components in the table are the essential default components that the ADOT Collector will support.

| Receiver                                                                                                                                                | Processor                                                                                                                                                                             | Exporter                                                                                                                                                                             | Extensions                                                                                                                                      |
|---------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------|
| [prometheusreceiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver#prometheus-receiver)       | [attributesprocessor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/attributesprocessor#attributes-processor)                                 | [`awsxrayexporter`](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/awsxrayexporter)                                                            | [healthcheckextension](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/healthcheckextension#health-check) |
| [otlpreceiver](https://github.com/open-telemetry/opentelemetry-collector/tree/main/receiver/otlpreceiver#otlp-receiver)                                 | [resourceprocessor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/resourceprocessor#resource-processor)                                       | [`awsemfexporter`](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/awsemfexporter)                                                              | [pprofextension](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/pprofextension#performance-profiler)     |
| [`awsecscontainermetricsreceiver`](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/awsecscontainermetricsreceiver) | [batchprocessor](https://github.com/open-telemetry/opentelemetry-collector/tree/main/processor/batchprocessor#batch-processor)                                                        | [prometheusremotewriteexporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/prometheusremotewriteexporter#prometheus-remote-write-exporter) | [zpagesextension](https://github.com/open-telemetry/opentelemetry-collector/tree/main/extension/zpagesextension#zpages)                         |
| [`awsxrayreceiver`](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/awsxrayreceiver)                               | [memorylimiterprocessor](https://github.com/open-telemetry/opentelemetry-collector/tree/main/processor/memorylimiterprocessor#memory-limiter-processor)                               | [loggingexporter](https://github.com/open-telemetry/opentelemetry-collector/tree/main/exporter/loggingexporter#logging-exporter)                                                     | [`ecsobserver`](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/observer/ecsobserver)                     |
| [statsdreceiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/statsdreceiver#statsd-receiver)                   | [probabilisticsamplerprocessor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/probabilisticsamplerprocessor#probabilistic-sampling-processor) | [otlpexporter](https://github.com/open-telemetry/opentelemetry-collector/tree/main/exporter/otlpexporter)                                                                            | [`awsproxy`](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/awsproxy)                                    |
| [zipkinreceiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/zipkinreceiver#zipkin-receiver)                   | [metricstransformprocessor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/metricstransformprocessor#metrics-transform-processor)              | [fileexporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/fileexporter#file-exporter)                                                      | [ballastextention](https://github.com/open-telemetry/opentelemetry-collector/tree/main/extension/ballastextension#memory-ballast)               |
| [jaegerreceiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/jaegerreceiver#jaeger-receiver)                   | [spanprocessor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/spanprocessor#span-processor)                                                   | [otlphttpexporter](https://github.com/open-telemetry/opentelemetry-collector/tree/main/exporter/otlphttpexporter#otlphttp-exporter)                                                  | [`sigv4authextension`](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/sigv4authextension)                |
| [`awscontainerinsightreceiver`](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/awscontainerinsightreceiver)       | [filterprocessor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/filterprocessor#filter-processor)                                             | [prometheusexporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/prometheusexporter#prometheus-exporter)                                    |                                                                                                                                                 |
| [kafka](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/kafkareceiver)                                             | [resourcedetectionprocessor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/resourcedetectionprocessor#resource-detection-processor)           | [datadogexporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/datadogexporter#datadog-exporter)                                             |                                                                                                                                                 |
|                                                                                                                                                         | [metricsgenerationprocessor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/metricsgenerationprocessor#metrics-generation-processor)           | [dynatraceexporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/dynatraceexporter#dynatrace-exporter)                                       |                                                                                                                                                 |
|                                                                                                                                                         | [cumulativetodeltaprocessor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/cumulativetodeltaprocessor#cumulative-to-delta-processor)          | [sapmexporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/sapmexporter#sapm-exporter)                                                      |                                                                                                                                                 |
|                                                                                                                                                         | [deltatorateprocessor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/deltatorateprocessor#delta-to-rate-processor)                            | [signalfxexporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/signalfxexporter#signalfx-metrics-exporter)                                  |                                                                                                                                                 |
|                                                                                                                                                         | [groupbytraceprocessor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/groupbytraceprocessor)                                                  | [logzioexporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/logzioexporter#logzio-exporter)                                                |                                                                                                                                                 |
|                                                                                                                                                         | [tailsamplingprocessor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/tailsamplingprocessor)                                                  | [kafka](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/kafkaexporter)                                                                          |                                                                                                                                                 |
|                                                                                                                                                         | [k8sattributesprocessor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/k8sattributesprocessor)                                                | [loadbalancingexporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/loadbalancingexporter)                                                  |                                                                                                                                                 |


Besides the components that interact with telemetry signals directly from the previous table, there is also support to the following confmap providers:

* file
* env
* YAML
* s3
* http
* https

More documentation for confmap providers can be found [here](https://aws-otel.github.io/docs/components/confmap-providers).

### Getting Started

#### Prerequisites

To build the ADOT Collector locally, you will need to have Golang installed. You can download and install Golang [here](https://golang.org/doc/install).

#### ADOT Collector Configuration

The ADOT Collector is built with a [default configuration](https://github.com/aws-observability/aws-otel-collector/blob/main/config.yaml).
The ADOT Collector configuration uses the same configuration syntax/design from [OpenTelemetry Collector](https://github.com/open-telemetry/opentelemetry-collector). For more information regarding OpenTelemetry Collector configuration please refer to the [upstream documentation](https://opentelemetry.io/docs/collector/configuration/). 
so you can customize or port your OpenTelemetry Collector configuration files when running ADOT Collector. Please refer to the `Try out the ADOT Collector` section on configuring ADOT Collector.

#### Try out the ADOT Collector

The ADOT Collector supports all AWS computing platforms and Docker/Kubernetes. Here are some examples on how to run the ADOT Collector to send telemetry data:

* [Run it with Docker](docs/developers/docker-demo.md)
* [Run it with ECS](https://aws-otel.github.io/docs/setup/ecs)
* [Run it with EKS](https://aws-otel.github.io/docs/getting-started/adot-eks-add-on)
* [Run it on EC2 (Amazon Linux 2)](docs/developers/linux-rpm-demo.md)
* [Run it on EC2 (Windows)](docs/developers/windows-other-demo.md)
* [Run it on EC2 (Debian)](docs/developers/debian-deb-demo.md)

#### Build Your Own Artifacts

Use the following instructions to build your own ADOT Collector artifacts:

* [Build Docker Image](docs/developers/build-docker.md)
* [Build RPM/Deb/MSI](docs/developers/build-aoc.md)

### Development

See [docs/developers](docs/developers/README.md).

### Benchmark

The latest performance report is [here](https://aws-observability.github.io/aws-otel-collector/benchmark/report), while the trends by testcase can be found [here](https://aws-observability.github.io/aws-otel-collector/benchmark/trend).
Both are updated on each successful CI run. The charts use the [github-action-benchmark](https://github.com/benchmark-action/github-action-benchmark) action and uses a modified layout to group the testcases.
The performance test can be conducted by following the [instructions](https://github.com/aws-observability/aws-otel-test-framework/blob/terraform/docs/get-performance-model.md) here.

### Support

 Please note that as per policy, we're providing support via GitHub on a best effort basis. However, if you have AWS Enterprise Support you can create a ticket and we will provide direct support within the respective SLAs.

For each merged pull request, a corresponding image with the naming convention of ```[ADOT_COLLECTOR_VERSION]-[GITHUB_SHA]``` is pushed to [public.ecr.aws/aws-otel-test/adot-collector-integration-test](https://gallery.ecr.aws/aws-otel-test/adot-collector-integration-test). 
This image is used for the integration tests. You can pull any of the images from there, however, we will not support any issues and pull requests for these test images.

### Supported Versions

Each ADOT Collector release is supported until there are two newer minor releases. For example, ADOT collector v0.16.1 will be supported until v0.18.0 is released.

### Security issue notifications
If you discover a potential security issue in this project we ask that you notify AWS/Amazon Security via our [vulnerability reporting page](http://aws.amazon.com/security/vulnerability-reporting/). Please do **not** create a public github issue.

### License

ADOT Collector is licensed under an Apache 2.0 license.
