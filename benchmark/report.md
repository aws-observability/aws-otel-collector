## Performance Report

**Version:** [v0.17.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.17.0)

**Commit ID:** [b841c9eb6b29b6239af38e71c027f4b7c2c314c9](https://github.com/aws-observability/aws-otel-collector/commit/b841c9eb6b29b6239af38e71c027f4b7c2c314c9)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.80 | 0.20 | 66.16 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.30 | 0.20 | 62.52 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.67 | 0.20 | 63.87 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.53 | 0.30 | 64.84 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 61.59 | 0.30 | 61.85 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.72 | 0.20 | 65.33 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.13 | 75.55 | 0.40 | 76.38 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 62.54 | 0.20 | 63.36 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 65.06 | 0.30 | 65.60 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.86 | 0.30 | 63.28 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.74 | 0.30 | 64.75 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.53 | 0.20 | 62.73 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.06 | 61.90 | 0.60 | 62.33 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.67 | 0.20 | 64.86 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.35 | 111.65 | 3.60 | 118.51 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 62.16 | 0.10 | 62.44 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 65.27 | 0.20 | 65.92 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.19 | 0.20 | 63.25 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.11 | 0.20 | 63.24 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.15 | 0.20 | 63.48 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 62.44 | 0.30 | 62.53 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.61 | 0.20 | 63.33 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.95 | 266.17 | 15.40 | 294.94 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 61.20 | 0.30 | 61.46 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.68 | 117.72 | 5.69 | 152.78 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.16 | 80.14 | 3.80 | 81.94 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 6.17 | 78.33 | 7.50 | 78.91 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.76 | 93.67 | 4.00 | 94.45 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 4.03 | 138.02 | 5.90 | 191.35 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 4.82 | 75.84 | 5.80 | 77.30 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.71 | 75.59 | 5.10 | 76.23 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.75 | 88.75 | 5.10 | 89.46 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 7.34 | 81.16 | 8.70 | 84.74 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 27.98 | 319.96 | 38.80 | 515.01 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 20.89 | 158.48 | 27.78 | 199.92 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 33.93 | 77.85 | 34.48 | 79.86 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 34.18 | 107.84 | 46.00 | 112.65 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 34.87 | 712.78 | 51.20 | 1224.66 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 29.62 | 77.68 | 32.51 | 78.29 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 37.44 | 79.62 | 47.13 | 80.79 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 28.94 | 91.28 | 30.80 | 91.93 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 33.48 | 463.09 | 40.56 | 572.79 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 36.58 | 453.68 | 47.63 | 752.33 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 20.48 | 178.90 | 26.66 | 207.04 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 138.62 | 86.65 | 147.17 | 87.73 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 108.97 | 127.82 | 110.79 | 134.37 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 112.72 | 3341.77 | 182.27 | 6011.06 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 110.65 | 80.70 | 116.41 | 81.19 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 142.58 | 15457.59 | 475.01 | 26574.42 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 112.67 | 97.56 | 118.94 | 99.54 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 30.67 | 502.29 | 37.97 | 555.17 |
