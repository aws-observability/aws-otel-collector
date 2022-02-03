## Performance Report

**Version:** [v0.16.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.16.0)

**Commit ID:** [56cf8ba0f6d766d2ea1407c48231f8957972e18b](https://github.com/aws-observability/aws-otel-collector/commit/56cf8ba0f6d766d2ea1407c48231f8957972e18b)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 65.44 | 0.40 | 66.70 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.94 | 0.20 | 63.55 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.64 | 0.10 | 63.74 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.66 | 0.10 | 62.97 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 62.44 | 0.20 | 62.53 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.00 | 0.20 | 63.26 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.13 | 75.53 | 0.40 | 76.77 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 61.05 | 0.10 | 62.79 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.50 | 0.30 | 66.00 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.10 | 0.20 | 63.24 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.61 | 0.20 | 62.77 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.67 | 0.20 | 63.20 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 62.46 | 0.20 | 62.99 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.01 | 0.20 | 64.23 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.40 | 113.52 | 2.90 | 119.41 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 61.14 | 0.10 | 61.70 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.31 | 0.20 | 64.51 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.97 | 0.20 | 61.58 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.06 | 0.20 | 63.80 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.45 | 0.20 | 62.55 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 61.53 | 0.20 | 62.19 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.52 | 0.30 | 64.13 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 8.10 | 267.73 | 14.40 | 298.20 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 62.62 | 0.10 | 63.28 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.98 | 147.53 | 6.20 | 209.34 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.29 | 80.49 | 2.60 | 82.21 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.29 | 77.56 | 5.00 | 78.39 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.26 | 95.76 | 3.70 | 97.36 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.12 | 140.64 | 4.40 | 191.05 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.26 | 74.32 | 3.60 | 74.92 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.82 | 75.72 | 4.20 | 76.96 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.53 | 88.68 | 3.90 | 89.03 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.88 | 81.80 | 7.00 | 86.47 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 26.07 | 496.48 | 34.80 | 817.40 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 17.69 | 154.85 | 23.19 | 188.47 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 33.26 | 78.73 | 34.20 | 79.59 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 29.31 | 105.43 | 30.81 | 110.11 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 27.22 | 749.27 | 39.99 | 1308.23 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 26.24 | 75.80 | 26.80 | 76.41 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 33.33 | 79.23 | 43.67 | 80.74 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 25.72 | 89.43 | 26.60 | 89.87 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 29.97 | 476.30 | 36.50 | 531.38 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 37.55 | 758.51 | 49.30 | 1255.73 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 17.02 | 170.97 | 21.99 | 200.68 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 124.16 | 85.00 | 133.15 | 86.02 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 107.94 | 126.79 | 110.80 | 133.88 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 115.23 | 3537.62 | 186.42 | 6340.06 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 106.59 | 81.83 | 111.40 | 82.96 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 152.30 | 15587.01 | 542.10 | 29726.26 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 120.24 | 95.36 | 121.52 | 97.26 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 29.65 | 507.94 | 37.39 | 574.03 |
