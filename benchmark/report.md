## Performance Report

**Version:** [v0.15.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.15.0)

**Commit ID:** [5f4a11c4bce4678adda475f24cd1f64bada1aa72](https://github.com/aws-observability/aws-otel-collector/commit/5f4a11c4bce4678adda475f24cd1f64bada1aa72)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.22 | 0.30 | 61.85 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 60.21 | 0.10 | 60.83 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.12 | 0.20 | 59.41 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.29 | 0.20 | 59.54 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 58.68 | 0.20 | 58.95 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.09 | 0.20 | 62.39 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.14 | 72.77 | 0.50 | 74.38 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 59.28 | 0.20 | 59.36 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 60.92 | 0.40 | 62.14 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.59 | 0.20 | 61.24 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 58.75 | 0.10 | 59.63 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.76 | 0.10 | 60.43 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 58.98 | 0.20 | 59.35 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.74 | 0.20 | 62.10 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.51 | 108.64 | 3.40 | 115.18 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 58.12 | 0.20 | 58.18 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.51 | 0.40 | 63.71 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.35 | 0.20 | 59.39 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 59.75 | 0.20 | 60.01 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.66 | 0.20 | 59.01 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 60.03 | 0.20 | 60.21 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.51 | 0.20 | 59.62 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.11 | 257.53 | 12.11 | 281.43 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 59.85 | 0.20 | 60.14 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.64 | 140.91 | 5.90 | 204.98 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.61 | 76.79 | 3.20 | 78.64 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 5.15 | 72.91 | 6.00 | 75.22 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.29 | 90.03 | 3.90 | 93.04 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.44 | 134.34 | 4.80 | 189.27 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.48 | 70.87 | 3.70 | 71.50 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.91 | 71.16 | 4.30 | 72.10 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.21 | 84.31 | 3.90 | 85.29 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.62 | 78.49 | 6.90 | 83.03 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 26.46 | 510.37 | 39.61 | 850.23 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.29 | 145.54 | 20.60 | 175.14 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 33.43 | 75.11 | 34.10 | 75.84 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 26.83 | 103.10 | 27.61 | 107.62 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 27.63 | 743.68 | 40.10 | 1237.01 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 26.69 | 72.35 | 27.89 | 72.98 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 32.89 | 75.82 | 42.99 | 77.13 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 26.56 | 86.27 | 27.11 | 86.59 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 29.89 | 471.40 | 36.83 | 557.19 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 37.92 | 719.00 | 51.19 | 1230.36 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 17.89 | 166.86 | 22.87 | 192.20 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 132.81 | 82.84 | 134.31 | 84.53 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 111.56 | 125.00 | 116.19 | 132.85 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 113.72 | 3557.10 | 180.82 | 6436.36 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 117.26 | 77.32 | 118.60 | 78.04 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 141.14 | 14957.96 | 496.88 | 25733.21 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 119.36 | 92.38 | 120.71 | 95.33 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 30.19 | 513.66 | 39.10 | 552.38 |
