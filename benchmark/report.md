## Performance Report

**Version:** [v0.35.1](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.35.1)

**Commit ID:** [3a1b6e62be30bcfe379e24174f9428b1608cc078](https://github.com/aws-observability/aws-otel-collector/commit/3a1b6e62be30bcfe379e24174f9428b1608cc078)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 83.03 | 0.20 | 83.73 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.18 | 0.20 | 81.30 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.16 | 88.01 | 0.30 | 89.26 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.07 | 85.05 | 0.20 | 89.82 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.74 | 0.20 | 81.99 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 79.23 | 0.20 | 80.02 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 79.61 | 0.20 | 80.44 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.75 | 0.20 | 81.92 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.07 | 103.76 | 0.30 | 106.47 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 80.30 | 0.10 | 80.93 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 83.57 | 0.20 | 84.82 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 82.74 | 0.20 | 83.40 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 85.74 | 0.20 | 91.40 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 86.19 | 0.20 | 87.09 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 80.55 | 0.20 | 81.97 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 81.15 | 0.20 | 82.31 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 82.64 | 0.20 | 83.26 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.07 | 0.20 | 82.03 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.65 | 128.23 | 1.10 | 137.67 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 82.16 | 0.10 | 83.11 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 82.67 | 0.20 | 82.97 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 79.72 | 0.20 | 80.45 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 86.56 | 0.20 | 92.07 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.18 | 86.12 | 0.30 | 91.96 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.00 | 0.20 | 81.91 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.88 | 0.20 | 81.96 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 80.19 | 0.10 | 81.23 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.26 | 0.30 | 82.31 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.48 | 252.39 | 7.90 | 285.41 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 79.87 | 0.10 | 80.77 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.60 | 107.42 | 3.80 | 108.11 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 1.32 | 110.89 | 1.50 | 112.93 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.08 | 113.07 | 4.60 | 114.68 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 5.46 | 106.49 | 7.20 | 107.00 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 23.01 | 135.66 | 26.30 | 150.55 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.87 | 113.72 | 4.90 | 115.04 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 5.31 | 126.52 | 5.90 | 144.37 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.59 | 133.38 | 4.30 | 143.76 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.39 | 114.61 | 3.80 | 115.82 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.33 | 105.66 | 3.70 | 105.97 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.09 | 116.23 | 3.50 | 116.73 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 2.43 | 111.65 | 2.90 | 116.30 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 17.55 | 110.02 | 18.20 | 111.39 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 10.94 | 113.98 | 12.00 | 117.10 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 30.01 | 121.73 | 30.40 | 124.80 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 59.97 | 137.87 | 62.30 | 149.03 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 41.48 | 114.02 | 45.40 | 115.52 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 26.56 | 112.51 | 28.00 | 114.55 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 46.23 | 164.61 | 50.50 | 167.61 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 25.75 | 446.64 | 29.20 | 533.10 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 24.21 | 112.23 | 24.70 | 115.78 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 26.21 | 110.08 | 27.90 | 110.92 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 24.84 | 117.92 | 26.20 | 118.21 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 22.17 | 140.92 | 23.20 | 177.25 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.23 | 118.48 | 26.40 | 123.09 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 55.95 | 126.20 | 59.50 | 139.24 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 106.71 | 122.24 | 111.90 | 126.02 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 128.70 | 121.61 | 148.00 | 123.44 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 130.14 | 128.41 | 155.60 | 159.56 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 100.31 | 111.45 | 102.80 | 113.14 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 188.06 | 212.22 | 200.29 | 217.92 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 101.31 | 1831.65 | 114.70 | 2242.74 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 91.37 | 111.72 | 97.39 | 113.41 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 110.32 | 15116.63 | 318.52 | 26069.44 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 83.32 | 124.01 | 87.60 | 125.90 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 100.97 | 242.00 | 104.30 | 390.97 |
