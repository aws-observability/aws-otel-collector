## Performance Report

**Version:** [v0.35.1](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.35.1)

**Commit ID:** [a0723fdba7d59ffd8b249bccc0b875f6af94cf12](https://github.com/aws-observability/aws-otel-collector/commit/a0723fdba7d59ffd8b249bccc0b875f6af94cf12)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 85.22 | 0.20 | 86.20 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.25 | 0.20 | 80.57 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.17 | 88.06 | 0.30 | 92.05 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 86.98 | 0.20 | 90.59 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 84.04 | 0.20 | 85.36 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 81.33 | 0.20 | 82.58 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 80.30 | 0.20 | 81.61 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.30 | 0.10 | 81.80 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.07 | 106.73 | 0.40 | 109.12 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 80.41 | 0.10 | 81.35 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 84.20 | 0.20 | 85.89 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 82.62 | 0.10 | 83.80 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 88.75 | 0.20 | 94.10 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.17 | 87.59 | 0.30 | 93.37 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 82.97 | 0.20 | 84.51 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 80.53 | 0.20 | 81.05 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 80.91 | 0.10 | 82.03 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.88 | 0.20 | 83.34 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.71 | 129.67 | 1.30 | 140.66 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 80.90 | 0.20 | 82.23 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 83.38 | 0.20 | 84.84 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 82.52 | 0.20 | 83.34 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 85.13 | 0.20 | 89.67 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 86.35 | 0.20 | 87.04 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 82.01 | 0.20 | 82.92 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 82.43 | 0.20 | 84.02 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 81.23 | 0.20 | 82.35 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.44 | 0.20 | 82.60 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.54 | 259.28 | 7.60 | 298.09 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 80.50 | 0.10 | 81.33 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.91 | 107.11 | 4.10 | 107.93 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 1.27 | 111.98 | 1.40 | 113.67 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.61 | 114.81 | 4.80 | 116.04 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 6.19 | 111.46 | 6.40 | 111.64 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 6.01 | 109.62 | 6.60 | 110.14 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.61 | 114.50 | 4.00 | 115.86 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 5.09 | 127.77 | 5.40 | 146.52 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.82 | 136.46 | 4.70 | 143.49 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 4.14 | 115.01 | 5.10 | 117.71 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.07 | 106.29 | 4.40 | 106.45 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 4.36 | 119.45 | 4.60 | 120.14 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 2.47 | 113.84 | 2.90 | 118.29 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 18.69 | 110.46 | 19.50 | 112.10 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 10.74 | 114.05 | 11.30 | 118.01 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 29.52 | 118.32 | 30.30 | 122.08 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 62.56 | 139.92 | 64.79 | 153.02 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 60.85 | 141.31 | 64.70 | 152.65 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 29.21 | 113.24 | 31.80 | 115.33 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 47.34 | 164.29 | 49.20 | 168.23 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 28.18 | 475.69 | 30.30 | 535.27 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 29.24 | 112.57 | 30.80 | 114.54 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 28.23 | 109.85 | 29.20 | 110.98 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 27.92 | 120.43 | 29.90 | 120.45 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 21.86 | 136.42 | 24.30 | 166.94 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.44 | 121.17 | 27.20 | 127.70 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 55.12 | 126.13 | 57.70 | 137.37 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 108.70 | 121.73 | 120.39 | 126.25 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 126.48 | 123.74 | 133.10 | 126.01 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 135.35 | 124.47 | 144.91 | 125.98 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 109.91 | 115.26 | 117.00 | 119.28 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 187.62 | 215.23 | 193.00 | 219.60 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 108.34 | 1968.70 | 125.60 | 2247.13 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 99.66 | 111.37 | 105.60 | 112.62 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 116.77 | 14773.06 | 420.12 | 24885.33 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 102.10 | 124.44 | 114.59 | 126.48 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 101.91 | 240.41 | 105.79 | 345.75 |
