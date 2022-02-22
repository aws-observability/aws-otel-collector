## Performance Report

**Version:** [v0.16.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.16.0)

**Commit ID:** [feb258384ebfb418991529190148036595f802b8](https://github.com/aws-observability/aws-otel-collector/commit/feb258384ebfb418991529190148036595f802b8)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.13 | 0.30 | 65.00 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.84 | 0.20 | 63.01 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 63.16 | 0.30 | 63.66 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.11 | 0.20 | 64.18 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 62.15 | 0.20 | 62.77 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.46 | 0.20 | 63.72 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.14 | 75.61 | 0.50 | 76.38 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 62.96 | 0.20 | 63.32 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 63.24 | 0.30 | 63.90 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.95 | 0.20 | 61.19 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.63 | 0.20 | 63.92 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.50 | 0.20 | 62.91 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 61.14 | 0.20 | 61.17 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.25 | 0.30 | 64.23 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.34 | 113.01 | 3.70 | 120.46 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 61.90 | 0.20 | 62.24 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 66.00 | 0.40 | 67.48 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.62 | 0.20 | 63.17 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.57 | 0.20 | 62.77 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.15 | 0.20 | 64.74 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.06 | 61.68 | 0.30 | 62.47 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.05 | 0.30 | 64.36 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.20 | 266.12 | 13.50 | 299.28 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 63.28 | 0.20 | 63.32 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 5.65 | 117.06 | 6.60 | 152.08 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.04 | 78.99 | 3.50 | 81.31 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 5.20 | 78.82 | 5.90 | 79.48 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.29 | 94.74 | 3.50 | 95.21 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.21 | 129.74 | 4.60 | 177.05 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 4.65 | 75.65 | 5.70 | 77.06 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.29 | 73.69 | 4.80 | 74.75 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.51 | 87.76 | 3.70 | 87.97 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 6.95 | 81.85 | 7.90 | 87.90 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 26.49 | 315.72 | 34.91 | 483.79 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 20.16 | 155.91 | 26.33 | 186.61 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 36.29 | 78.07 | 38.50 | 80.08 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 27.56 | 106.52 | 28.40 | 110.59 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 27.46 | 762.24 | 40.12 | 1209.32 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 35.24 | 76.55 | 39.22 | 77.04 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 34.71 | 78.33 | 46.19 | 79.73 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 43.62 | 89.34 | 56.80 | 89.77 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 34.54 | 466.62 | 40.00 | 547.53 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 37.62 | 456.47 | 52.40 | 749.63 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 19.42 | 177.09 | 26.94 | 206.37 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 133.59 | 85.58 | 145.62 | 86.46 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 133.91 | 127.78 | 143.40 | 135.43 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 144.89 | 3632.34 | 214.00 | 5874.77 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 122.14 | 80.39 | 151.04 | 81.15 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 154.28 | 15356.57 | 453.92 | 26360.11 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 111.97 | 95.39 | 113.81 | 97.71 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 31.32 | 502.54 | 40.37 | 567.03 |
