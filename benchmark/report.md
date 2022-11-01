## Performance Report

**Version:** [v0.23.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.23.0)

**Commit ID:** [14a43733e8396b3094d847f8a38f2ba5e8ecaa99](https://github.com/aws-observability/aws-otel-collector/commit/14a43733e8396b3094d847f8a38f2ba5e8ecaa99)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.52 | 0.20 | 65.90 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.19 | 0.20 | 64.60 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.95 | 0.20 | 66.42 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.90 | 0.10 | 65.22 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 65.70 | 0.20 | 66.32 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.96 | 0.20 | 65.22 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.10 | 78.83 | 0.30 | 79.98 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 63.52 | 0.20 | 64.00 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.35 | 0.20 | 69.39 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.22 | 0.20 | 64.32 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.61 | 0.20 | 64.29 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 63.93 | 0.20 | 64.51 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 65.44 | 0.20 | 65.70 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.26 | 0.20 | 66.43 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.95 | 106.77 | 1.60 | 116.15 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 64.45 | 0.20 | 65.06 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.10 | 0.20 | 69.73 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.26 | 0.20 | 64.45 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.78 | 0.20 | 65.22 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.90 | 0.20 | 64.11 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 64.33 | 0.20 | 64.79 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.04 | 0.20 | 66.28 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 5.59 | 225.01 | 9.00 | 259.33 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 64.25 | 0.20 | 64.55 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.85 | 78.24 | 4.20 | 79.97 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.10 | 72.65 | 3.10 | 82.53 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.05 | 79.83 | 4.50 | 82.79 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 4.26 | 77.96 | 4.50 | 80.36 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.19 | 147.15 | 4.50 | 195.51 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 4.16 | 79.18 | 4.60 | 80.61 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.96 | 78.57 | 4.30 | 79.81 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.54 | 89.99 | 3.80 | 92.02 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.49 | 83.80 | 17.60 | 88.05 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.10 | 81.89 | 19.70 | 85.59 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.00 | 75.92 | 0.00 | 86.95 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 34.17 | 85.39 | 35.60 | 88.12 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 28.98 | 80.07 | 29.60 | 81.96 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 28.31 | 686.42 | 39.70 | 1198.88 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 27.05 | 78.86 | 27.40 | 80.08 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 29.33 | 81.51 | 31.00 | 84.07 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 26.10 | 91.73 | 26.60 | 93.54 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 36.32 | 306.25 | 52.50 | 443.37 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 26.51 | 95.17 | 27.80 | 101.14 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 25.80 | 127.69 | 30.80 | 172.62 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 127.77 | 96.35 | 136.00 | 102.45 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 121.09 | 82.71 | 122.30 | 86.44 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 124.60 | 3336.64 | 183.51 | 5721.90 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 114.61 | 81.07 | 118.20 | 82.78 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 107.04 | 17938.75 | 270.12 | 31720.27 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 114.08 | 97.08 | 116.10 | 99.34 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 38.06 | 434.26 | 62.30 | 589.98 |
