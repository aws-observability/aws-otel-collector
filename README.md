![CI](https://github.com/aws-observability/aws-otel-collector/actions/workflows/CI.yml/badge.svg)
![CD](https://github.com/aws-observability/aws-otel-collector/actions/workflows/CD.yml/badge.svg)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/aws-observability/aws-otel-collector)

### Overview

AWS Distro for OpenTelemetry Collector (ADOT Collector) is an AWS supported version of the upstream OpenTelemetry Collector and is distributed by Amazon. It supports the selected components from the OpenTelemetry community. It is fully compatible with AWS computing platforms including EC2, ECS, and EKS. It enables users to send telemetry data to AWS CloudWatch Metrics, Traces, and Logs backends as well as the other supported backends.

See the [AWS Distro for OpenTelemetry documentation](https://aws-otel.github.io/docs/getting-started/collector) for more information.

### Getting Help

Use the community resources below for getting help with the ADOT Collector.
* Use [GitHub issues](https://github.com/aws-observability/aws-otel-collector/issues) to report bugs and request features.
* Join our GitHub [Community](https://github.com/aws-observability/aws-otel-community) for AWS Distro for OpenTelemetry to ask your questions, file issues, or request enhancements.
* Open a support ticket with [AWS Support](http://docs.aws.amazon.com/awssupport/latest/user/getting-started.html).
* If you think you may have found a bug, open a [bug report](https://github.com/aws-observability/aws-otel-collector/issues/new?template=bug_report.md).
* For contributing guidelines, refer to [CONTRIBUTING.md](CONTRIBUTING.md).

#### ADOT Collector Built-in Components

This table represents the supported components of the ADOT Collector. The highlighted components below are developed by AWS in-house. The rest of the components in the table are the essential default components that the ADOT Collector will support.

| Receiver                        | Processor                     | Exporter                           | Extensions             |
|---------------------------------|-------------------------------|------------------------------------|------------------------|
| prometheusreceiver              | attributesprocessor           | `awsxrayexporter`                  | healthcheckextension   |
| otlpreceiver                    | resourceprocessor             | `awsemfexporter`                   | pprofextension         |
| `awsecscontainermetricsreceiver`| batchprocessor                | `awsprometheusremotewriteexporter`* | zpagesextension        |
| `awsxrayreceiver`               | memorylimiterprocessor        | loggingexporter                    | `ecsobserver`          |
| statsdreceiver                  | probabilisticsamplerprocessor | otlpexporter                       | `awsproxy`             |
| zipkinreceiver                  | metricstransformprocessor     | fileexporter                       | ballastextention       |
| jaegerreceiver                  | spanprocessor                 | otlphttpexporter                   | `sigv4authextension`   |
| `awscontainerinsightreceiver`   | filterprocessor               | prometheusexporter                 |                        |
|                                 | resourcedetectionprocessor    | datadogexporter                    |                        |
|                                 | `metricsgenerationprocessor`  | dynatraceexporter                  |                        |
|                                 | cumulativetodeltaprocessor    | sapmexporter                       |                        |
|                                 | deltatorateprocessor          | signalfxexporter                   |                        |
|                                 |                               | logzioexporter                     |                        |
|                                 |                               | prometheusremotewriteexporter      |                        |

\* Note that the `awsprometheusremotewriteexporter` will be removed at some point after v0.19.0. Users who want to send metrics to Amazon Managed Service for Prometheus will need to instead use the [Prometheus Remote Write Exporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/exporter/prometheusremotewriteexporter/README.md) along with the [Sigv4 Authenticator Extension](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/extension/sigv4authextension/README.md) to achieve the same result.

#### ADOT Collector AWS Components

* [OpenTelemetry Collector](https://github.com/open-telemetry/opentelemetry-collector/)
* [Trace X-Ray Exporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/awsxrayexporter)
* [Metrics EMF Exporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/awsemfexporter)
* [ECS Container Metrics Receiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/awsecscontainermetricsreceiver)
* [StatsD Receiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/statsdreceiver)
* [ECS Observer Extension](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/observer/ecsobserver)

### Getting Started

#### Prerequisites

To build the ADOT Collector locally, you will need to have Golang installed. You can download and install Golang [here](https://golang.org/doc/install).

#### ADOT Collector Configuration

We built in a [default configuration](https://github.com/aws-observability/aws-otel-collector/blob/main/config.yaml) to our docker image and other format of release.
So, you can run the ADOT Collector out of the box with the default settings.
Also, the ADOT Collector configuration uses the same configuration syntax/design from [OpenTelemetry Collector](https://github.com/open-telemetry/opentelemetry-collector)
so you can customize or port your OpenTelemetry Collector configuration files when running ADOT Collector. Please refer to the `Try out ADOT Collector` section on configuring ADOT Collector.

#### Try out the ADOT Collector

The ADOT Collector supports all AWS computing platforms and Docker/Kubernetes. Here are some examples on how to run the ADOT Collector to send telemetry data:

* [Run it with Docker](docs/developers/docker-demo.md)
* [Run it with ECS](docs/developers/ecs-demo.md)
* [Run it with EKS](docs/developers/eks-demo.md)
* [Run it on EC2 (Amazon Linux 2)](docs/developers/linux-rpm-demo.md)
* [Run it on EC2 (Windows)](docs/developers/windows-other-demo.md)
* [Run it on EC2 (Debian)](docs/developers/debian-deb-demo.md)

#### Build Your Own Artifacts

Use the following instructions to build your own ADOT Collector artifacts:

* [Build Docker Image](docs/developers/build-docker.md)
* [Build RPM/Deb/MSI](docs/developers/build-aoc.md)

### Development

See [docs/developers](docs/developers/README.md)

### Benchmark

The latest performance report is [here](https://aws-observability.github.io/aws-otel-collector/benchmark/report), while the trends by testcase can be found [here](https://aws-observability.github.io/aws-otel-collector/benchmark/trend).
Both are updated on each successful CI run. The charts use the [github-action-benchmark](https://github.com/benchmark-action/github-action-benchmark) action and uses a modified layout to group the testcases.
The performance test can be conducted by following the [instructions](https://github.com/aws-observability/aws-otel-test-framework/blob/terraform/docs/get-performance-model.md) here.

### Support

For each merged pull request, a corresponding image with the naming convention of ```[ADOT_COLLECTOR_VERSION]-[GITHUB_SHA]``` is pushed to [public.ecr.aws/aws-otel-test/adot-collector-integration-test](https://gallery.ecr.aws/aws-otel-test/adot-collector-integration-test). 
This image is used for the integration tests. You can pull any of the images from there, however, we will not support any issues and pull requests for these test images.

### Supported Versions

Each ADOT Collector release is supported until there are two newer minor releases. For example, ADOT collector v0.16.1 will be supported until v0.18.0 is released.


### License

ADOT Collector is licensed under an Apache 2.0 license.
