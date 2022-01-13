## Performance Report

**Version:** [v0.15.0-1684262232](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.15.0-1684262232)

**Commit ID:** [543db00eb9f7f469d512df6f9ab643a4697df1c4](https://github.com/aws-observability/aws-otel-collector/commit/543db00eb9f7f469d512df6f9ab643a4697df1c4)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.01 | 0.30 | 63.16 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.16 | 0.20 | 61.15 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 60.25 | 0.20 | 60.92 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 58.88 | 0.20 | 59.29 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 59.25 | 0.20 | 59.74 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.55 | 0.20 | 62.70 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.14 | 72.03 | 0.40 | 73.25 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 58.96 | 0.10 | 60.05 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.86 | 0.30 | 63.02 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.00 | 0.20 | 59.04 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 60.88 | 0.20 | 61.08 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.45 | 0.20 | 60.60 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 60.38 | 0.20 | 60.56 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.32 | 0.20 | 60.44 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.48 | 105.92 | 3.30 | 115.32 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 57.97 | 0.10 | 58.07 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.45 | 0.30 | 61.93 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.44 | 0.20 | 58.70 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.32 | 0.20 | 60.62 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.92 | 0.10 | 59.33 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 59.29 | 0.20 | 59.35 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.44 | 0.20 | 62.22 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.57 | 259.22 | 14.10 | 283.62 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 58.50 | 0.10 | 58.50 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 5.22 | 155.48 | 6.50 | 211.50 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.34 | 76.31 | 2.70 | 77.89 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.84 | 72.20 | 5.50 | 75.03 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.01 | 91.62 | 3.20 | 93.38 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 2.97 | 134.46 | 4.30 | 186.85 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.32 | 70.36 | 3.70 | 70.80 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.78 | 71.65 | 5.10 | 72.48 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 4.08 | 85.45 | 4.50 | 85.83 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.68 | 79.54 | 6.40 | 84.07 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.71 | 546.60 | 38.19 | 874.18 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.58 | 151.25 | 21.57 | 179.76 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 31.76 | 74.76 | 32.80 | 76.15 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 27.66 | 104.16 | 28.41 | 108.45 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 28.45 | 714.09 | 41.79 | 1243.02 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 27.74 | 71.05 | 28.92 | 71.70 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 32.97 | 76.52 | 42.78 | 78.20 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 26.96 | 86.26 | 28.10 | 86.74 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 29.97 | 495.85 | 36.84 | 614.53 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 37.15 | 785.78 | 49.05 | 1278.42 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 18.26 | 169.45 | 25.12 | 196.56 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 119.58 | 82.81 | 123.68 | 84.10 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 109.82 | 124.79 | 111.19 | 131.26 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 123.33 | 3143.67 | 191.33 | 5719.70 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 118.25 | 76.46 | 120.71 | 77.44 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 149.18 | 16972.31 | 493.24 | 27905.10 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 111.98 | 91.29 | 115.59 | 94.34 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 30.55 | 507.55 | 41.50 | 564.31 |
