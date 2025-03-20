## Performance Report

**Version:** [v0.43.1](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.43.1)

**Commit ID:** [141783c41718b8984c2ec510583ade1a829f2b7d](https://github.com/aws-observability/aws-otel-collector/commit/141783c41718b8984c2ec510583ade1a829f2b7d)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.45 | 112.56 | 1.30 | 114.07 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.18 | 107.13 | 0.30 | 109.37 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 9.78 | 127.95 | 10.90 | 132.83 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.18 | 100.67 | 0.40 | 102.52 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.19 | 111.42 | 0.40 | 114.11 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.36 | 106.14 | 0.50 | 106.95 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.25 | 111.97 | 0.50 | 115.34 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.07 | 109.08 | 0.20 | 111.53 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 89.51 | 0.10 | 89.85 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 2.10 | 120.36 | 2.30 | 122.70 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 2.54 | 123.88 | 2.90 | 126.35 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.48 | 121.52 | 0.70 | 123.31 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.45 | 111.72 | 0.70 | 114.21 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.53 | 125.66 | 0.80 | 128.79 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 1.58 | 115.84 | 1.70 | 118.67 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.82 | 123.03 | 1.00 | 126.56 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.72 | 140.94 | 1.40 | 152.06 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 90.98 | 0.10 | 91.22 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 9.79 | 135.33 | 10.30 | 142.11 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 9.30 | 131.03 | 10.40 | 135.29 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 1.70 | 122.84 | 2.00 | 128.37 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 1.57 | 119.49 | 1.80 | 122.31 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 1.88 | 125.41 | 2.10 | 130.81 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 6.99 | 124.67 | 7.30 | 127.75 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 3.70 | 124.22 | 4.20 | 129.78 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.47 | 263.64 | 7.80 | 293.29 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 91.08 | 0.10 | 91.82 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.48 | 106.91 | 3.70 | 108.10 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 91.33 | 0.20 | 93.27 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 94.55 | 0.20 | 96.08 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.17 | 98.48 | 0.30 | 99.93 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 95.39 | 0.20 | 97.43 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 92.80 | 0.20 | 93.47 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 91.11 | 0.10 | 91.48 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 91.78 | 0.20 | 92.13 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 91.92 | 0.20 | 92.79 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 92.13 | 0.20 | 93.11 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 90.56 | 0.10 | 91.65 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.05 | 89.24 | 0.20 | 90.02 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 18.59 | 110.77 | 20.50 | 112.86 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 89.01 | 0.10 | 89.90 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 93.57 | 0.20 | 94.21 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 95.77 | 0.20 | 97.26 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.18 | 96.55 | 0.30 | 98.58 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 91.28 | 0.20 | 91.92 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 90.98 | 0.20 | 91.89 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 89.59 | 0.10 | 90.09 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 90.62 | 0.20 | 90.78 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 90.43 | 0.20 | 91.19 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 90.94 | 0.20 | 91.50 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 91.99 | 0.10 | 92.53 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.06 | 125.61 | 26.30 | 133.77 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 92.18 | 0.20 | 92.70 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 94.54 | 0.20 | 95.55 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.07 | 96.49 | 0.30 | 97.81 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.07 | 96.09 | 0.20 | 97.60 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 90.06 | 0.20 | 90.21 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 91.23 | 0.20 | 92.00 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 92.22 | 0.20 | 92.64 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 88.90 | 0.10 | 89.55 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 89.88 | 0.10 | 89.95 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 92.37 | 0.20 | 93.82 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 90.88 | 0.20 | 91.11 |
