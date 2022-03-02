## Performance Report

**Version:** [v0.17.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.17.0)

**Commit ID:** [ba4622c9611e15b059e915bb01d85caa011615be](https://github.com/aws-observability/aws-otel-collector/commit/ba4622c9611e15b059e915bb01d85caa011615be)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 65.35 | 0.20 | 66.40 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.58 | 0.20 | 62.85 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 62.86 | 0.30 | 63.29 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.18 | 0.20 | 63.49 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 63.02 | 0.20 | 63.36 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.42 | 0.20 | 62.87 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.14 | 74.63 | 0.40 | 75.42 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 61.82 | 0.10 | 61.91 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 65.68 | 0.40 | 66.55 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.79 | 0.20 | 61.88 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 61.23 | 0.20 | 61.45 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 62.42 | 0.40 | 62.75 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.06 | 62.02 | 0.30 | 62.27 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.90 | 0.30 | 64.69 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.30 | 116.09 | 2.80 | 120.11 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 62.41 | 0.10 | 63.10 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.53 | 0.40 | 65.59 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.59 | 0.20 | 62.43 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.53 | 0.30 | 64.49 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.54 | 0.20 | 62.87 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.06 | 62.28 | 0.20 | 62.32 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 63.64 | 0.20 | 63.75 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.29 | 267.64 | 14.40 | 300.16 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 62.11 | 0.20 | 62.19 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.78 | 119.92 | 5.80 | 155.77 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.10 | 78.70 | 3.50 | 81.19 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.95 | 78.77 | 5.70 | 79.34 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.90 | 94.01 | 4.30 | 94.45 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.34 | 139.38 | 4.60 | 185.29 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 5.21 | 75.13 | 6.10 | 76.38 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.84 | 74.90 | 5.80 | 75.98 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 4.35 | 87.62 | 4.80 | 87.77 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 7.11 | 80.23 | 8.70 | 83.97 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 26.28 | 370.08 | 36.81 | 557.86 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 20.09 | 157.36 | 25.47 | 187.54 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 48.11 | 79.45 | 54.60 | 80.80 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 28.30 | 107.71 | 30.61 | 111.43 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 29.22 | 744.15 | 43.20 | 1223.07 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 28.78 | 78.05 | 29.60 | 78.57 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 45.51 | 80.70 | 57.40 | 81.81 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 28.25 | 89.24 | 35.41 | 89.76 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 32.97 | 471.05 | 39.30 | 545.17 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 36.42 | 499.71 | 45.21 | 753.94 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 20.09 | 175.22 | 26.40 | 212.63 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 119.43 | 85.87 | 127.10 | 86.88 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 120.29 | 127.97 | 124.22 | 134.38 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 134.88 | 3354.57 | 199.78 | 6142.25 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 127.75 | 79.68 | 129.36 | 80.30 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 145.94 | 13265.15 | 408.96 | 24483.70 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 124.53 | 94.59 | 138.80 | 97.65 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 32.09 | 518.97 | 40.56 | 572.24 |
