## Performance Report

**Version:** [v0.27.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.27.0)

**Commit ID:** [e0e063cda879a0a81a98de9c84e7de52b3b36463](https://github.com/aws-observability/aws-otel-collector/commit/e0e063cda879a0a81a98de9c84e7de52b3b36463)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 70.80 | 0.20 | 71.23 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 69.49 | 0.20 | 69.56 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.38 | 0.10 | 66.49 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.02 | 0.20 | 68.58 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 67.74 | 0.20 | 68.46 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.95 | 0.20 | 67.17 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.09 | 83.26 | 0.30 | 83.88 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 66.91 | 0.10 | 67.50 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 67.95 | 0.20 | 67.96 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.51 | 0.20 | 68.63 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.95 | 0.20 | 67.00 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 68.72 | 0.10 | 69.67 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 68.25 | 0.20 | 69.08 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.93 | 0.20 | 69.70 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.01 | 110.71 | 1.70 | 116.19 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 66.36 | 0.20 | 66.88 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 69.01 | 0.20 | 69.14 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.58 | 0.10 | 68.24 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 68.06 | 0.10 | 68.45 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.67 | 0.10 | 66.67 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 67.27 | 0.20 | 67.46 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.07 | 0.20 | 68.22 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 5.84 | 242.79 | 9.40 | 270.29 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 68.05 | 0.10 | 68.49 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.96 | 81.65 | 4.10 | 82.57 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.87 | 87.92 | 15.70 | 90.01 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.46 | 85.72 | 5.20 | 87.51 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 4.93 | 83.89 | 5.30 | 84.85 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.87 | 141.06 | 5.20 | 187.81 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 4.01 | 84.58 | 4.80 | 85.07 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.30 | 81.88 | 4.50 | 82.59 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 2.95 | 94.65 | 3.00 | 95.37 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 4.66 | 85.54 | 16.80 | 87.58 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.01 | 85.38 | 19.40 | 88.03 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 24.42 | 157.72 | 40.40 | 187.99 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 29.48 | 89.52 | 30.00 | 91.47 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 27.79 | 84.01 | 28.60 | 86.16 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 26.38 | 734.53 | 37.20 | 1250.87 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 27.30 | 81.89 | 28.60 | 83.65 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 27.25 | 84.23 | 29.20 | 86.08 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 25.09 | 96.87 | 25.70 | 97.69 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 32.95 | 272.03 | 49.30 | 355.04 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 27.13 | 96.98 | 28.50 | 103.16 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 25.20 | 191.98 | 46.50 | 220.43 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 106.69 | 95.73 | 110.50 | 106.73 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 103.18 | 86.68 | 118.00 | 89.63 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 95.74 | 3167.41 | 159.99 | 6088.91 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 89.56 | 84.67 | 95.70 | 86.43 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 112.90 | 17073.47 | 360.20 | 28426.53 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 86.06 | 98.92 | 92.90 | 100.15 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 32.65 | 368.31 | 47.50 | 458.36 |
