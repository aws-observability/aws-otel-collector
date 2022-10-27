## Performance Report

**Version:** [v0.22.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.22.0)

**Commit ID:** [723bcd5f509e0a4dcd3173bb7daaf12642d8c320](https://github.com/aws-observability/aws-otel-collector/commit/723bcd5f509e0a4dcd3173bb7daaf12642d8c320)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.21 | 0.20 | 65.35 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.08 | 0.20 | 63.18 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.55 | 0.20 | 64.18 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.00 | 0.20 | 64.32 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 62.85 | 0.20 | 63.46 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.27 | 0.20 | 66.71 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.09 | 79.55 | 0.30 | 80.68 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 65.78 | 0.10 | 66.31 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.74 | 0.10 | 65.19 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.91 | 0.20 | 65.08 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.37 | 0.20 | 64.66 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 64.49 | 0.20 | 64.71 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 65.63 | 0.20 | 66.24 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 66.05 | 0.20 | 66.79 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.94 | 105.99 | 1.60 | 110.55 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 64.54 | 0.10 | 64.98 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.10 | 0.20 | 66.40 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 63.54 | 0.20 | 64.27 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.17 | 0.20 | 64.45 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 62.79 | 0.20 | 63.05 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 63.81 | 0.20 | 64.67 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.18 | 0.20 | 65.25 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 5.42 | 226.37 | 9.40 | 253.03 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 63.66 | 0.10 | 64.16 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.85 | 77.84 | 4.30 | 79.16 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.00 | 75.54 | 3.40 | 83.14 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 3.79 | 81.38 | 4.00 | 84.75 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 5.23 | 79.59 | 5.60 | 81.68 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.22 | 136.64 | 4.90 | 191.63 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.99 | 79.13 | 4.30 | 80.68 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.81 | 78.58 | 4.10 | 80.07 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.27 | 90.44 | 3.60 | 92.33 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.30 | 83.19 | 17.50 | 88.04 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.32 | 81.79 | 19.80 | 85.38 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.00 | 89.20 | 0.00 | 103.99 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 29.83 | 84.80 | 30.80 | 86.89 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 31.10 | 79.57 | 32.40 | 81.71 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 27.15 | 714.09 | 37.60 | 1198.97 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 28.13 | 79.25 | 29.50 | 80.80 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 29.08 | 81.39 | 31.20 | 85.46 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 27.32 | 91.95 | 29.00 | 94.00 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 37.31 | 340.76 | 52.10 | 500.85 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 28.36 | 91.88 | 29.80 | 100.53 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 26.31 | 157.47 | 37.90 | 191.34 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 136.24 | 93.00 | 137.61 | 97.33 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 127.05 | 82.46 | 132.99 | 86.98 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 123.11 | 3156.94 | 182.50 | 6123.09 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 118.63 | 81.07 | 120.39 | 82.69 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 113.95 | 17726.05 | 344.70 | 30936.20 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 112.08 | 96.73 | 112.81 | 98.75 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 35.56 | 451.97 | 49.70 | 619.03 |
