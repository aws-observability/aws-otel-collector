## Performance Report

**Version:** [v0.15.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.15.0)

**Commit ID:** [f6dddb4ad9bb61d34cb8c94d50f542dec8cd4b8c](https://github.com/aws-observability/aws-otel-collector/commit/f6dddb4ad9bb61d34cb8c94d50f542dec8cd4b8c)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.24 | 0.30 | 62.05 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.08 | 0.20 | 59.16 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 59.80 | 0.20 | 60.03 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 59.00 | 0.20 | 59.12 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 59.32 | 0.20 | 59.47 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.52 | 0.20 | 61.62 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.13 | 73.20 | 0.40 | 74.71 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 58.53 | 0.10 | 58.90 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.21 | 0.20 | 59.25 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.19 | 0.20 | 58.72 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.80 | 0.20 | 58.93 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.40 | 0.20 | 60.37 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 59.06 | 0.20 | 59.08 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.17 | 0.20 | 61.39 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.44 | 105.93 | 3.10 | 115.32 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 60.58 | 0.20 | 60.78 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.98 | 0.30 | 64.11 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.57 | 0.20 | 60.60 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 60.81 | 0.20 | 61.32 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.36 | 0.20 | 59.49 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 59.56 | 0.10 | 59.70 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.01 | 0.20 | 61.56 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.45 | 256.06 | 14.20 | 286.08 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 60.14 | 0.20 | 60.28 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.76 | 152.05 | 5.70 | 208.05 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.44 | 75.84 | 2.80 | 77.74 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 6.15 | 73.91 | 7.00 | 75.41 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.22 | 90.79 | 3.60 | 91.80 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.12 | 127.31 | 4.50 | 182.76 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 4.19 | 69.96 | 4.60 | 70.33 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.76 | 72.25 | 5.20 | 73.04 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.71 | 84.60 | 4.10 | 85.11 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.68 | 78.26 | 6.80 | 82.10 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.89 | 478.55 | 39.62 | 820.29 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.55 | 148.47 | 22.13 | 170.42 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 31.49 | 74.58 | 33.20 | 75.67 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 26.43 | 101.91 | 27.09 | 106.25 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 29.86 | 658.72 | 41.80 | 1223.37 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 28.13 | 71.14 | 29.41 | 72.09 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 33.73 | 75.10 | 43.50 | 76.71 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 26.60 | 86.64 | 27.30 | 87.06 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 31.60 | 486.96 | 38.89 | 556.49 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 37.07 | 789.99 | 49.48 | 1233.90 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 17.71 | 166.39 | 23.69 | 196.94 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 132.53 | 82.58 | 136.41 | 84.18 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 104.00 | 123.80 | 107.01 | 131.44 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 120.29 | 3591.28 | 187.50 | 6271.46 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 115.00 | 76.91 | 116.87 | 77.63 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 144.17 | 13896.77 | 453.41 | 24078.16 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 115.57 | 92.39 | 117.01 | 95.77 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 29.78 | 499.46 | 37.30 | 552.13 |
