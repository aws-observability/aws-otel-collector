## Performance Report

**Version:** [v0.15.0-1650955773](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.15.0-1650955773)

**Commit ID:** [f57ae1bc811622cc0ec86ae0dd096fcf9e7e1b4b](https://github.com/aws-observability/aws-otel-collector/commit/f57ae1bc811622cc0ec86ae0dd096fcf9e7e1b4b)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.62 | 0.20 | 60.20 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.79 | 0.20 | 61.02 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.09 | 0.20 | 61.49 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.32 | 0.20 | 58.38 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.04 | 58.77 | 0.20 | 59.04 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.95 | 0.20 | 60.09 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.13 | 72.59 | 0.30 | 74.74 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 0.13 | 72.42 | 0.40 | 73.84 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 0.56 | 67.37 | 0.70 | 67.98 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 59.91 | 0.20 | 60.38 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.48 | 0.40 | 63.84 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.67 | 0.20 | 58.68 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.98 | 0.20 | 61.34 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.89 | 0.20 | 60.35 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.08 | 59.66 | 0.20 | 60.54 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.00 | 0.20 | 60.45 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.31 | 115.16 | 3.00 | 119.90 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 1.30 | 116.67 | 2.50 | 120.71 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 5.17 | 68.96 | 5.50 | 69.58 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 58.89 | 0.20 | 59.43 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.78 | 0.20 | 60.05 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.82 | 0.20 | 60.13 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.01 | 0.20 | 60.15 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 60.19 | 0.20 | 60.29 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.05 | 58.81 | 0.20 | 58.94 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 60.27 | 0.20 | 60.61 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.87 | 256.43 | 14.60 | 307.60 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 8.08 | 258.53 | 13.70 | 303.01 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 25.83 | 68.83 | 26.79 | 69.25 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 59.30 | 0.20 | 59.95 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 4.71 | 141.62 | 6.20 | 207.48 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.36 | 75.18 | 2.70 | 77.28 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 6.29 | 72.07 | 6.70 | 74.49 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.70 | 89.18 | 4.00 | 90.58 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.77 | 129.18 | 5.20 | 183.48 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.18 | 69.51 | 3.50 | 69.89 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.88 | 71.02 | 4.20 | 71.88 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 4.06 | 71.80 | 4.40 | 72.60 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.61 | 84.03 | 4.00 | 84.75 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.57 | 77.92 | 7.00 | 83.34 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 34.16 | 579.24 | 48.20 | 1046.22 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.82 | 152.34 | 22.63 | 188.96 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 33.18 | 74.23 | 35.09 | 75.02 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 28.39 | 101.91 | 29.30 | 107.71 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 28.50 | 661.41 | 39.89 | 1182.65 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 27.65 | 71.60 | 29.30 | 72.43 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 32.76 | 76.31 | 42.21 | 78.46 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 35.86 | 77.83 | 44.61 | 85.32 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 25.33 | 85.90 | 26.21 | 86.52 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 28.93 | 458.96 | 37.51 | 523.61 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 54.02 | 1044.25 | 66.94 | 1675.58 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.97 | 169.14 | 22.63 | 190.69 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 131.67 | 82.44 | 136.20 | 83.96 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 105.16 | 123.67 | 108.50 | 129.83 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 124.05 | 3032.34 | 185.02 | 5615.56 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 121.70 | 76.69 | 122.60 | 77.98 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 140.71 | 13436.02 | 490.67 | 26426.86 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 167.07 | 12132.93 | 437.56 | 23175.12 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 104.55 | 91.97 | 106.09 | 94.74 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 28.55 | 495.05 | 40.02 | 558.65 |
