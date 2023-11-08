## Performance Report

**Version:** [v0.35.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.35.0)

**Commit ID:** [e06f194abb0590756f8097dca0a10416a94cd374](https://github.com/aws-observability/aws-otel-collector/commit/e06f194abb0590756f8097dca0a10416a94cd374)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 85.95 | 0.20 | 86.75 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.55 | 0.20 | 81.55 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.17 | 86.23 | 0.30 | 87.17 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.17 | 85.97 | 0.40 | 91.15 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.82 | 0.20 | 81.53 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 79.44 | 0.20 | 80.60 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 79.60 | 0.20 | 80.47 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 81.18 | 0.10 | 81.94 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.07 | 105.20 | 0.30 | 107.28 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 79.11 | 0.10 | 80.02 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 84.12 | 0.20 | 84.82 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 78.48 | 0.20 | 79.39 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 84.92 | 0.20 | 90.37 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 85.32 | 0.20 | 86.09 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.46 | 0.10 | 82.51 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 80.76 | 0.20 | 81.53 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 80.66 | 0.20 | 81.32 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 84.00 | 0.20 | 85.07 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.67 | 128.67 | 1.20 | 137.52 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 79.57 | 0.10 | 80.26 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 83.38 | 0.20 | 84.62 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 79.98 | 0.20 | 80.15 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 84.55 | 0.20 | 85.34 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 85.17 | 0.20 | 86.13 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.25 | 0.20 | 82.07 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 79.44 | 0.20 | 79.78 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 79.62 | 0.10 | 80.79 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 79.44 | 0.10 | 80.48 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.45 | 252.14 | 7.50 | 290.89 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 78.95 | 0.10 | 79.88 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.55 | 106.76 | 3.70 | 107.42 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 1.34 | 110.86 | 1.50 | 112.38 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.12 | 114.21 | 4.60 | 115.13 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 24.38 | 141.86 | 26.70 | 154.01 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 4.59 | 108.99 | 4.90 | 109.70 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 4.20 | 114.67 | 4.70 | 116.36 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 5.24 | 130.15 | 6.00 | 147.19 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.29 | 135.01 | 4.00 | 144.68 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.66 | 113.47 | 4.10 | 114.63 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 5.08 | 106.39 | 5.30 | 106.52 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 4.48 | 116.86 | 4.90 | 117.08 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 2.50 | 113.51 | 2.70 | 117.74 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 18.11 | 108.95 | 18.70 | 111.66 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 10.93 | 118.37 | 11.30 | 123.86 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 28.97 | 118.10 | 29.70 | 122.43 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 42.53 | 109.01 | 43.10 | 109.39 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 50.24 | 110.04 | 51.30 | 111.68 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 28.04 | 111.95 | 29.30 | 113.65 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 46.62 | 165.27 | 47.80 | 169.25 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 25.86 | 482.67 | 28.40 | 533.59 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 24.62 | 113.69 | 25.30 | 115.36 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 26.60 | 109.57 | 27.10 | 110.35 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 24.05 | 120.81 | 24.60 | 121.28 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 22.15 | 146.15 | 23.30 | 171.97 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 24.77 | 120.27 | 26.10 | 126.07 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 55.50 | 126.73 | 57.20 | 138.09 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 106.44 | 124.00 | 111.80 | 131.12 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 133.48 | 119.78 | 141.50 | 123.74 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 138.04 | 156.60 | 143.90 | 163.81 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 97.50 | 110.79 | 101.60 | 112.84 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 189.51 | 212.57 | 195.78 | 216.19 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 101.51 | 1989.56 | 116.39 | 2245.28 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 95.73 | 110.83 | 103.80 | 112.67 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 113.70 | 15932.91 | 367.82 | 31929.79 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 87.10 | 123.78 | 93.80 | 125.61 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 100.69 | 252.81 | 105.69 | 350.42 |
