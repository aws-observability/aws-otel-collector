## Performance Report

**Version:** [v0.15.0-634e4d9](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.15.0-634e4d9)

**Commit ID:** [634e4d9453e7a6525dbb4061c6390fa802e16660](https://github.com/aws-observability/aws-otel-collector/commit/634e4d9453e7a6525dbb4061c6390fa802e16660)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.73 | 0.30 | 62.94 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 58.74 | 0.20 | 58.91 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 60.08 | 0.20 | 60.12 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.24 | 0.20 | 59.62 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 60.44 | 0.20 | 60.73 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.58 | 0.20 | 60.86 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.13 | 73.22 | 0.40 | 74.42 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 59.27 | 0.20 | 59.47 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.72 | 0.30 | 63.33 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.77 | 0.20 | 61.16 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 60.09 | 0.10 | 60.82 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.57 | 0.20 | 60.06 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 58.85 | 0.20 | 59.21 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.78 | 0.20 | 62.05 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.30 | 105.75 | 3.50 | 115.10 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 59.90 | 0.20 | 60.19 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.26 | 0.20 | 59.88 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.04 | 0.20 | 59.52 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.14 | 0.20 | 60.19 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.30 | 0.20 | 58.51 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 59.03 | 0.20 | 59.07 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.97 | 0.20 | 60.08 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.71 | 260.15 | 15.30 | 285.29 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 59.88 | 0.20 | 60.27 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.85 | 143.30 | 6.00 | 205.44 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.44 | 75.85 | 2.70 | 77.76 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 5.49 | 72.67 | 6.30 | 74.57 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 2.89 | 90.15 | 3.20 | 92.27 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.46 | 129.40 | 4.80 | 180.11 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.63 | 69.85 | 4.10 | 70.80 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.03 | 71.40 | 4.40 | 72.32 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.04 | 83.98 | 3.30 | 84.11 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.58 | 78.56 | 6.60 | 82.51 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.76 | 484.94 | 35.20 | 797.91 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.97 | 148.62 | 21.89 | 171.79 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 31.76 | 75.02 | 32.21 | 76.08 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 26.79 | 102.89 | 27.22 | 107.11 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 29.25 | 673.14 | 42.10 | 1185.23 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 26.30 | 72.18 | 27.00 | 72.85 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 33.47 | 75.20 | 43.27 | 76.48 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 26.53 | 86.64 | 27.00 | 87.47 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 29.84 | 479.91 | 36.20 | 578.11 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 37.94 | 752.23 | 50.67 | 1232.67 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.79 | 170.67 | 22.88 | 196.41 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 111.31 | 82.88 | 113.98 | 84.26 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 121.78 | 124.17 | 123.85 | 132.05 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 132.25 | 3474.73 | 201.29 | 6173.17 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 122.16 | 77.38 | 123.39 | 78.07 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 138.73 | 16279.07 | 504.03 | 27881.95 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 116.12 | 91.76 | 119.21 | 94.13 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 28.49 | 507.33 | 36.33 | 556.47 |
