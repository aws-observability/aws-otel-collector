## Performance Report

**Version:** [v0.42.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.42.0)

**Commit ID:** [b79e473186d9a1cb5ef114b4f87717c3de33f1c2](https://github.com/aws-observability/aws-otel-collector/commit/b79e473186d9a1cb5ef114b4f87717c3de33f1c2)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.42 | 126.25 | 0.60 | 127.15 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 1.97 | 131.93 | 8.80 | 145.63 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.19 | 112.42 | 0.40 | 113.66 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.19 | 106.12 | 0.40 | 107.81 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.20 | 114.24 | 0.40 | 115.77 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.36 | 121.46 | 0.60 | 122.53 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.23 | 123.69 | 0.50 | 126.37 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.07 | 119.27 | 0.30 | 120.41 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 95.57 | 0.20 | 98.13 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 2.00 | 139.06 | 2.20 | 141.13 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.46 | 136.73 | 0.70 | 138.67 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 9.38 | 148.00 | 10.20 | 151.93 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.49 | 130.46 | 0.70 | 132.92 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.54 | 140.20 | 0.80 | 143.96 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 1.46 | 134.40 | 1.70 | 136.40 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.75 | 140.22 | 1.00 | 143.37 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.68 | 158.88 | 1.30 | 169.57 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 94.85 | 0.20 | 97.22 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 9.47 | 154.24 | 10.20 | 161.47 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 9.09 | 147.58 | 10.90 | 151.42 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 1.77 | 155.43 | 6.80 | 158.38 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 1.54 | 137.80 | 1.80 | 142.31 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 1.82 | 143.39 | 2.60 | 148.12 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 6.53 | 144.01 | 7.20 | 146.57 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 3.45 | 143.99 | 3.80 | 149.01 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.30 | 270.73 | 7.70 | 303.36 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 94.56 | 0.20 | 96.71 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.53 | 122.83 | 3.80 | 123.51 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.05 | 93.80 | 0.20 | 96.15 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 98.11 | 0.30 | 100.51 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.09 | 101.28 | 0.40 | 104.74 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.21 | 100.13 | 0.50 | 102.67 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 94.77 | 0.20 | 97.18 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 94.97 | 0.20 | 97.22 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 96.70 | 0.20 | 99.04 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 94.70 | 0.20 | 96.92 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.05 | 94.83 | 0.30 | 96.14 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 93.34 | 0.20 | 95.64 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.05 | 94.32 | 0.20 | 96.56 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 16.85 | 125.81 | 17.20 | 127.48 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 95.79 | 0.20 | 98.14 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 99.04 | 0.20 | 101.59 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.16 | 101.32 | 0.30 | 106.12 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 100.63 | 0.20 | 105.83 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 95.13 | 0.20 | 96.89 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.02 | 96.00 | 0.20 | 98.51 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 95.27 | 0.20 | 97.69 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 94.43 | 0.20 | 96.82 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 94.47 | 0.30 | 96.86 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 96.61 | 0.20 | 98.71 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 94.49 | 0.30 | 96.65 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 23.80 | 139.14 | 24.60 | 143.14 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.06 | 95.63 | 0.20 | 97.88 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 98.70 | 0.20 | 101.23 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 101.11 | 0.30 | 105.33 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.07 | 101.18 | 0.30 | 106.17 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 95.84 | 0.20 | 97.78 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 95.36 | 0.20 | 97.88 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 97.00 | 0.20 | 99.25 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 95.79 | 0.20 | 98.55 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.05 | 92.88 | 0.20 | 95.16 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 95.41 | 0.20 | 96.69 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 96.83 | 0.20 | 98.60 |
