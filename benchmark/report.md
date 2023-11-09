## Performance Report

**Version:** [v0.35.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.35.0)

**Commit ID:** [146aafd548a8b0859adb602644a2f55340313f89](https://github.com/aws-observability/aws-otel-collector/commit/146aafd548a8b0859adb602644a2f55340313f89)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 82.96 | 0.20 | 84.20 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 79.87 | 0.20 | 80.37 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 87.08 | 0.20 | 87.66 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.17 | 87.72 | 0.40 | 92.42 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.26 | 0.20 | 82.46 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 79.63 | 0.20 | 79.80 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 84.92 | 0.10 | 86.07 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 79.63 | 0.10 | 80.43 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.08 | 104.67 | 0.30 | 106.82 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 79.31 | 0.10 | 80.13 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 82.10 | 0.30 | 82.75 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.18 | 0.20 | 80.68 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.17 | 88.38 | 0.30 | 92.32 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 86.42 | 0.20 | 89.82 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.74 | 0.20 | 82.17 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 79.93 | 0.20 | 80.41 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 80.98 | 0.20 | 81.48 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 80.85 | 0.20 | 82.08 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.66 | 129.47 | 1.10 | 137.81 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 78.73 | 0.20 | 79.54 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 81.57 | 0.20 | 82.76 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.23 | 0.20 | 82.22 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 84.38 | 0.20 | 85.46 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 86.46 | 0.20 | 86.63 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.96 | 0.20 | 82.13 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 78.76 | 0.20 | 79.46 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 80.16 | 0.10 | 81.15 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.99 | 0.30 | 81.91 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.41 | 252.56 | 7.80 | 289.89 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 78.69 | 0.10 | 79.53 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.43 | 105.93 | 3.60 | 106.45 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 1.28 | 110.68 | 1.50 | 112.78 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.96 | 112.93 | 5.40 | 114.11 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 5.17 | 108.45 | 5.90 | 108.62 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 5.31 | 108.25 | 6.00 | 108.50 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.41 | 112.19 | 3.70 | 113.24 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 6.10 | 125.81 | 7.50 | 142.70 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.70 | 137.64 | 4.40 | 145.47 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.12 | 114.18 | 3.80 | 118.37 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.21 | 106.73 | 3.50 | 106.98 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.77 | 117.15 | 4.40 | 118.12 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 2.47 | 113.09 | 2.70 | 118.16 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 17.52 | 109.22 | 17.90 | 110.75 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 11.02 | 113.53 | 11.50 | 118.85 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 29.00 | 122.94 | 29.70 | 125.61 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 60.33 | 141.45 | 63.00 | 151.33 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 61.22 | 139.56 | 64.40 | 150.89 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 27.71 | 111.57 | 28.40 | 114.27 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 46.46 | 162.16 | 47.60 | 169.12 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 26.97 | 487.04 | 29.40 | 534.96 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 24.87 | 112.33 | 26.30 | 113.92 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 27.69 | 108.55 | 28.30 | 109.66 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 25.23 | 118.97 | 25.90 | 118.98 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 21.87 | 141.96 | 23.00 | 172.42 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.69 | 117.36 | 26.90 | 122.63 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 55.42 | 125.15 | 56.50 | 131.74 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 111.92 | 121.52 | 115.09 | 128.26 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 133.13 | 118.50 | 141.90 | 123.78 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 121.20 | 123.12 | 151.41 | 151.58 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 103.07 | 111.92 | 105.89 | 113.70 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 188.30 | 211.33 | 196.90 | 217.15 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 98.40 | 1983.78 | 113.81 | 2247.59 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 95.19 | 110.65 | 104.90 | 112.59 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 117.73 | 14629.13 | 348.57 | 28737.33 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 87.13 | 124.73 | 95.20 | 126.82 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 103.39 | 241.99 | 107.00 | 313.46 |
