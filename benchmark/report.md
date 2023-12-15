## Performance Report

**Version:** [v0.36.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.36.0)

**Commit ID:** [834d7948e108c84a1a9efbbfaf34f6baddb466bc](https://github.com/aws-observability/aws-otel-collector/commit/834d7948e108c84a1a9efbbfaf34f6baddb466bc)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 84.94 | 0.20 | 86.73 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.69 | 0.20 | 82.85 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 87.24 | 0.20 | 88.46 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.17 | 88.49 | 0.30 | 92.59 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.36 | 0.20 | 82.46 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.67 | 0.10 | 82.94 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 82.36 | 0.20 | 83.76 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 82.06 | 0.20 | 83.66 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.07 | 106.00 | 0.30 | 108.29 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 81.72 | 0.10 | 83.19 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 84.31 | 0.20 | 85.80 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.50 | 0.20 | 81.88 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.17 | 87.71 | 0.30 | 91.11 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 86.29 | 0.20 | 91.43 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 83.18 | 0.20 | 84.57 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.74 | 0.20 | 81.90 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 82.02 | 0.20 | 83.04 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.86 | 0.10 | 83.35 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.62 | 129.26 | 1.30 | 139.63 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 79.64 | 0.10 | 80.89 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 83.94 | 0.20 | 84.79 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.46 | 0.20 | 81.68 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 86.03 | 0.20 | 90.43 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 87.78 | 0.20 | 93.27 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 82.80 | 0.20 | 84.32 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 82.88 | 0.20 | 84.01 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 82.17 | 0.20 | 83.09 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 82.30 | 0.10 | 83.12 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.42 | 257.51 | 7.50 | 290.42 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 80.34 | 0.10 | 81.71 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.37 | 108.37 | 3.50 | 108.85 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 1.30 | 113.38 | 1.50 | 114.83 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.06 | 115.64 | 5.50 | 118.22 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 5.73 | 111.28 | 6.30 | 111.40 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 5.05 | 109.98 | 5.60 | 110.43 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 4.47 | 115.11 | 5.20 | 116.57 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 5.40 | 124.70 | 6.70 | 144.78 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.25 | 137.91 | 3.90 | 146.34 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.70 | 115.96 | 3.90 | 117.21 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.67 | 107.85 | 3.80 | 108.58 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.30 | 119.49 | 3.60 | 120.10 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 2.51 | 113.23 | 2.80 | 116.83 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 17.25 | 110.61 | 17.60 | 112.14 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 10.77 | 115.22 | 11.30 | 119.20 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 30.90 | 121.63 | 33.00 | 125.26 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 41.64 | 110.18 | 42.20 | 110.44 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 61.95 | 142.12 | 65.10 | 153.46 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 28.07 | 114.23 | 28.50 | 116.42 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 48.11 | 164.40 | 50.00 | 168.45 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 27.15 | 486.08 | 30.60 | 536.68 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 25.86 | 112.56 | 26.90 | 114.84 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 28.09 | 110.72 | 30.30 | 112.25 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 26.80 | 119.92 | 27.90 | 120.22 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 21.89 | 139.55 | 22.70 | 172.52 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.14 | 120.56 | 26.00 | 126.31 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 55.22 | 128.07 | 56.90 | 136.45 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 111.24 | 125.18 | 118.70 | 136.84 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 159.20 | 152.91 | 167.00 | 162.09 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 138.10 | 123.53 | 149.20 | 126.50 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 110.84 | 115.16 | 121.80 | 119.67 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 187.79 | 216.69 | 193.50 | 223.83 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 108.05 | 1915.13 | 121.90 | 2249.60 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 96.82 | 112.71 | 102.80 | 114.13 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 130.45 | 15160.41 | 453.89 | 27516.61 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 92.88 | 125.63 | 101.90 | 128.49 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 102.26 | 250.43 | 106.40 | 381.25 |
