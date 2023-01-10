## Performance Report

**Version:** [v0.24.1](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.24.1)

**Commit ID:** [c97ef241bc12d4fb23fbc87df6ec23b45da53298](https://github.com/aws-observability/aws-otel-collector/commit/c97ef241bc12d4fb23fbc87df6ec23b45da53298)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 71.81 | 0.20 | 73.22 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.42 | 0.30 | 68.15 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.54 | 0.20 | 69.24 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.97 | 0.20 | 69.53 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 67.31 | 0.20 | 68.06 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 69.03 | 0.20 | 70.21 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.09 | 81.94 | 0.20 | 83.50 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 67.35 | 0.10 | 67.96 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 73.30 | 0.20 | 74.60 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.05 | 0.20 | 67.40 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 68.01 | 0.20 | 68.72 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 69.04 | 0.30 | 70.10 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 68.13 | 0.20 | 68.68 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 69.76 | 0.20 | 70.29 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.05 | 112.44 | 1.80 | 115.98 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 67.70 | 0.10 | 68.12 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 73.13 | 0.20 | 73.99 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 67.40 | 0.30 | 67.94 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.24 | 0.20 | 69.16 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 68.25 | 0.20 | 68.77 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 68.28 | 0.20 | 69.13 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.99 | 0.20 | 69.53 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 6.00 | 236.17 | 9.80 | 262.68 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 68.29 | 0.20 | 69.15 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.82 | 79.92 | 4.10 | 81.38 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.04 | 87.21 | 15.90 | 90.47 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.32 | 85.18 | 4.60 | 88.32 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 4.47 | 83.50 | 4.80 | 85.05 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 4.05 | 143.45 | 5.50 | 193.71 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 4.55 | 82.08 | 5.20 | 84.05 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.70 | 82.78 | 3.90 | 84.38 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.88 | 96.41 | 4.10 | 97.73 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.41 | 87.33 | 18.00 | 92.75 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.68 | 85.46 | 20.40 | 88.62 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 27.10 | 157.28 | 43.50 | 189.09 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 28.71 | 92.36 | 29.30 | 93.99 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 27.84 | 82.50 | 28.40 | 85.09 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 26.16 | 723.89 | 37.40 | 1221.80 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 26.29 | 82.50 | 26.80 | 84.92 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 29.06 | 84.94 | 30.70 | 87.22 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 24.47 | 97.24 | 25.20 | 98.55 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 35.55 | 344.16 | 49.00 | 539.40 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 27.90 | 98.35 | 29.50 | 105.84 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 24.46 | 182.00 | 38.80 | 210.29 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 118.60 | 97.46 | 120.20 | 101.22 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 115.47 | 86.26 | 116.20 | 89.58 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 109.13 | 3306.51 | 170.20 | 5814.41 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 103.96 | 85.00 | 105.20 | 87.54 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 101.78 | 19166.49 | 380.18 | 32210.55 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 105.88 | 99.19 | 111.30 | 101.35 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 35.04 | 427.76 | 48.80 | 539.32 |
