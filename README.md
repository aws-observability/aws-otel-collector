[![codecov](https://codecov.io/gh/mxiamxia/aws-opentelemetry-collector/branch/master/graph/badge.svg)](https://codecov.io/gh/mxiamxia/aws-opentelemetry-collector)
![CI](https://github.com/mxiamxia/aws-opentelemetry-collector/workflows/CI/badge.svg)
![CD](https://github.com/mxiamxia/aws-opentelemetry-collector/workflows/CD/badge.svg)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/mxiamxia/aws-opentelemetry-collector)


### Overview

AWS Observability Collector(AOC) is a certified OpenTelemetry Collector and distributed by Amazon. It supports the selected components from the OpenTelemetry community and is fully compatible with AWS computing platforms including EC2, ECS and EKS. It will send telemetry data to AWS CloudWatch Metrics, Traces and Logs backends as well as the other backends we have claimed to support with verification.

### Getting Help

Use the community resources below for getting help with AWS Observability Collector. We use [GitHub issues](https://github.com/mxiamxia/aws-opentelemetry-collector/issues) for tracking bugs and feature requests.

* Ask a question in [AWS CloudWatch Forum](https://forums.aws.amazon.com/forum.jspa?forumID=138).
* Open a support ticket with [AWS Support](http://docs.aws.amazon.com/awssupport/latest/user/getting-started.html).
* If you think you may have found a bug, open an [issue](https://github.com/mxiamxia/aws-opentelemetry-collector/issues/new).
* For contributing guidelines refer [CONTRIBUTING.md](https://github.com/mxiamxia/aws-opentelemetry-collector/blob/master/CONTRIBUTING.md).

#### AOC Built-in Components (in 2020)

This table represents the supported components of AWS observability collector in 2020. The highlighted components below are developed by AWS in-house. The rest of the components in the table are the essential default components that AOC will support.
 
| Receiver             | Processor                     | Exporter           | Extensions             |
|----------------------|-------------------------------|--------------------|------------------------|
| prometheusreceiver   | attributesprocessor           | awsxrayexporter    | healthcheckextension   |
| otlpreceiver         | resourceprocessor             | awsemfexporter     | pprofextension         |
| statsdreceiver       | queuedprocessor               | prometheusexporter | zpagesextension        |
| xrayreceiver.        | batchprocessor                | loggingexporter    |                        |
| ecsmetadatareceiver  | memorylimite                  | otlpexporter       |                        |
|                      | tailsamplingprocessor         |                    |                        |
|                      | probabilisticsamplerprocessor |                    |                        |
|                      | spanprocessor                 |                    |                        |
|                      | filterprocessor               |                    |                        |

#### AOC AWS Components
* [OpenTelemetry Collector-v0.7.0](https://github.com/open-telemetry/opentelemetry-collector/)
* [Trace X-Ray Exporter-v0.7.0](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/master/exporter/awsxrayexporter)
* [Metrics EMF Exporter](https://github.com/shaochengwang/opentelemetry-collector-contrib/blob/emf_exporter_pr/exporter/awsemfexporter/README.md)
* More coming

### Getting Started
#### Prerequisites
To build AOC locally, you will need to have Golang installed. You can download and install Golang [here](https://golang.org/doc/install).

#### Build Your Own Artifacts
Use the following instruction to build your own AOC artifacts:
* [Build Docker Image](docs/developers/build-docker.md)
* [Build RPM/Deb/MSI](docs/developers/build-aoc.md)

#### Try out AOC Beta
AOC supports all AWS computing platforms and docker/kubernetes. Here are some examples on how to run AOC to send telemetry data:
* [Run it with Docker](docs/developers/docker-demo.md)
* [Run it with ECS/EKS](docs/developers/emf-demo.md)
* [Run it on AWS Linux EC2](docs/developers/linux-rpm-demo.md)
* [Run it on AWS Windows EC2](docs/developers/windows-other-demo.md)
* [Run it on AWS Debian EC2](docs/developers/debian-deb-demo.md)

### Release Process
* [Release new version](docs/developers/release-new-version.md)

### Benchmark

aws-observability-collector is based on open-telemetry-collector. Here is the benchmark of AWSXray trace exporter and AWSEMF metrics exporter.

This table shows the performance of AWSEMF exporter against 1kData/sec, 5kData/sec, and 10kData/sec metrics:

| Test                | Result | Duration | CPU Avg% | CPU Max% | RAM Avg MiB | RAM Max MiB | Sent Items | Received Items |
|---------------------|--------|----------|----------|----------|-------------|-------------|------------|----------------|
| Metric1kDPS/AWSEmf  | PASS   | 17s      |       34 |     36.3 |          59 |          82 |     105000 |         105000 |
| Metric5kDPS/AWSEmf  | PASS   | 43s      |     58.5 |    101.5 |         505 |         678 |     508200 |         508200 |
| Metric10kDPS/AWSEmf | PASS   | 72s      |       63 |    145.7 |         971 |        1178 |     849100 |         849100 |

This table shows the performance of AWSXray  exporter against 1kData/sec,5kData/sec and 10kData/sec spans(traces).

| Test                | Result | Duration | CPU Avg% | CPU Max% | RAM Avg MiB | RAM Max MiB | Sent Items | Received Items |
|---------------------|--------|----------|----------|----------|-------------|-------------|------------|----------------|
| Trace1kSPS/AwsXray  | PASS   | 15s      |      8.5 |     11.6 |          32 |          36 |      15000 |          15000 |
| Trace5kSPS/AwsXray  | PASS   | 15s      |    26.12 |     27.8 |          33 |          38 |      74400 |          74400 |
| Trace10kSPS/AwsXray | PASS   | 15s      |     43.8 |     45.3 |          37 |          43 |     132500 |         132500 |


### License
aws-observability-collector is under Apache 2.0 license.
