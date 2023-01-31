## Performance Report

**Version:** [v0.25.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.25.0)

**Commit ID:** [3fed655247fd3c3758770f1f826fc97684bf3f61](https://github.com/aws-observability/aws-otel-collector/commit/3fed655247fd3c3758770f1f826fc97684bf3f61)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 69.02 | 0.20 | 69.15 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.60 | 0.20 | 68.45 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 68.69 | 0.20 | 69.57 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.05 | 0.20 | 68.71 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 68.72 | 0.20 | 69.48 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 69.12 | 0.20 | 69.23 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.09 | 82.34 | 0.30 | 83.52 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 67.68 | 0.20 | 68.46 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 67.67 | 0.20 | 68.25 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.19 | 0.20 | 67.64 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 67.61 | 0.20 | 68.19 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.52 | 0.20 | 67.10 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 67.66 | 0.20 | 68.53 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 70.41 | 0.20 | 70.64 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.07 | 113.68 | 1.80 | 117.22 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 67.52 | 0.20 | 67.81 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 68.94 | 0.20 | 69.53 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.42 | 0.20 | 67.90 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.50 | 0.20 | 68.51 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.96 | 0.20 | 67.45 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 66.60 | 0.20 | 66.67 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 69.12 | 0.20 | 69.65 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 5.93 | 233.31 | 9.90 | 266.61 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 67.77 | 0.10 | 68.48 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.73 | 80.96 | 3.90 | 82.31 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.98 | 87.26 | 15.70 | 89.49 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 3.32 | 84.89 | 3.60 | 88.70 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 4.06 | 82.44 | 4.70 | 84.33 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.32 | 143.56 | 4.60 | 190.24 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.94 | 82.95 | 4.60 | 85.55 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.97 | 82.31 | 4.10 | 84.39 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.38 | 93.81 | 3.80 | 95.85 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.17 | 86.08 | 17.40 | 89.28 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.48 | 84.69 | 19.90 | 87.42 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 26.48 | 157.55 | 44.90 | 196.28 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 28.27 | 93.42 | 30.10 | 96.75 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 26.92 | 81.94 | 27.30 | 84.14 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 25.38 | 738.35 | 39.00 | 1216.47 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 24.92 | 82.92 | 25.50 | 84.65 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 27.34 | 84.37 | 29.20 | 87.17 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 25.35 | 95.74 | 26.10 | 96.80 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 35.71 | 342.43 | 50.50 | 520.50 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 27.29 | 99.86 | 29.10 | 107.34 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 24.16 | 182.09 | 41.60 | 218.49 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 120.69 | 103.03 | 125.81 | 109.27 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 119.41 | 86.29 | 124.01 | 89.91 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 104.91 | 3417.42 | 160.30 | 5871.00 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 107.29 | 84.44 | 110.00 | 85.82 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 108.62 | 18717.60 | 297.71 | 32053.62 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 101.39 | 98.96 | 102.60 | 100.84 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 36.85 | 445.42 | 59.00 | 620.81 |
