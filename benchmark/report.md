## Performance Report

**Version:** [v0.26.2](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.26.2)

**Commit ID:** [885999751bec9edda538d0d2797bac9bfb227e5b](https://github.com/aws-observability/aws-otel-collector/commit/885999751bec9edda538d0d2797bac9bfb227e5b)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 69.27 | 0.20 | 69.49 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 68.18 | 0.20 | 69.07 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.92 | 0.10 | 68.56 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.18 | 0.20 | 67.25 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 68.22 | 0.20 | 68.80 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.57 | 0.20 | 68.05 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.10 | 81.43 | 0.30 | 82.49 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 66.54 | 0.10 | 66.67 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 68.40 | 0.10 | 69.17 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.45 | 0.20 | 67.68 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 68.00 | 0.10 | 68.68 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.39 | 0.10 | 67.45 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 67.56 | 0.10 | 67.83 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.54 | 0.20 | 67.64 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.15 | 112.61 | 1.80 | 117.33 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 67.17 | 0.10 | 67.35 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 70.51 | 0.20 | 71.11 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.47 | 0.20 | 68.90 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.97 | 0.20 | 68.35 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.43 | 0.20 | 69.65 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 68.27 | 0.20 | 68.80 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.79 | 0.20 | 69.16 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 6.01 | 237.76 | 9.90 | 266.47 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 67.25 | 0.10 | 67.42 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.93 | 79.40 | 4.20 | 81.17 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.04 | 87.06 | 12.70 | 89.31 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 3.73 | 84.11 | 4.20 | 87.71 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 4.22 | 80.76 | 4.70 | 82.80 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.39 | 139.60 | 4.50 | 190.89 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.91 | 80.89 | 4.20 | 82.71 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.55 | 80.04 | 4.80 | 81.41 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.13 | 93.10 | 3.60 | 94.86 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 4.86 | 85.00 | 17.50 | 88.10 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.39 | 83.71 | 20.00 | 87.03 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 25.30 | 154.51 | 40.80 | 190.91 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 28.28 | 89.10 | 29.80 | 92.14 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 28.37 | 81.02 | 29.60 | 83.89 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 25.54 | 703.99 | 37.00 | 1248.78 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 26.02 | 82.76 | 26.50 | 84.47 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 28.75 | 83.95 | 29.60 | 86.26 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 24.78 | 94.53 | 25.20 | 95.99 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 35.01 | 281.99 | 50.90 | 422.59 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 26.49 | 95.91 | 27.80 | 101.41 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 25.88 | 181.01 | 43.30 | 211.41 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 103.62 | 96.94 | 112.40 | 102.85 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 102.26 | 86.20 | 108.91 | 89.79 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 100.33 | 3211.99 | 159.60 | 5760.70 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 100.86 | 84.60 | 105.00 | 86.79 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 118.49 | 15634.23 | 342.18 | 30798.65 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 88.37 | 99.56 | 92.61 | 101.11 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 33.28 | 397.25 | 48.90 | 525.59 |
