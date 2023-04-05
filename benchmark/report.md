## Performance Report

**Version:** [v0.27.1](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.27.1)

**Commit ID:** [e5271f96fff11c18510e47e33f6ae8abd5118a9a](https://github.com/aws-observability/aws-otel-collector/commit/e5271f96fff11c18510e47e33f6ae8abd5118a9a)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 68.67 | 0.20 | 69.38 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.40 | 0.20 | 68.20 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 69.48 | 0.10 | 69.83 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.63 | 0.20 | 68.31 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 69.10 | 0.20 | 69.69 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 68.02 | 0.10 | 68.42 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.09 | 81.52 | 0.30 | 82.37 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 67.68 | 0.20 | 68.28 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 68.85 | 0.10 | 69.69 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.05 | 0.20 | 68.23 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 67.18 | 0.10 | 67.30 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 69.22 | 0.20 | 69.37 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 68.00 | 0.10 | 68.11 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 69.04 | 0.20 | 70.21 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.03 | 111.20 | 1.80 | 115.89 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 66.52 | 0.10 | 67.62 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 69.61 | 0.20 | 69.84 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 68.80 | 0.20 | 69.53 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 68.32 | 0.20 | 68.32 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.49 | 0.20 | 65.79 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 69.82 | 0.20 | 70.33 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 69.10 | 0.20 | 69.71 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 6.35 | 232.40 | 10.40 | 268.56 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 67.13 | 0.10 | 67.24 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.57 | 80.50 | 3.80 | 81.87 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.03 | 87.88 | 15.60 | 91.11 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.42 | 85.21 | 4.80 | 89.37 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 4.88 | 81.84 | 5.30 | 83.87 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.33 | 144.27 | 4.80 | 193.11 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 4.08 | 82.02 | 4.70 | 83.44 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.89 | 81.49 | 4.40 | 82.95 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.95 | 94.09 | 4.20 | 95.12 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 4.82 | 85.56 | 17.20 | 88.97 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.26 | 85.32 | 20.00 | 88.63 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 24.86 | 167.53 | 39.20 | 203.40 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 28.80 | 88.82 | 29.40 | 93.05 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 27.87 | 83.70 | 28.60 | 86.58 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 24.96 | 735.60 | 35.70 | 1196.86 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 27.45 | 83.67 | 29.90 | 85.91 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 28.97 | 85.31 | 30.90 | 88.20 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 24.05 | 95.92 | 25.60 | 97.17 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 31.52 | 291.03 | 44.40 | 405.75 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 27.53 | 97.75 | 28.80 | 103.83 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 24.01 | 179.63 | 42.00 | 210.66 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 107.47 | 92.54 | 112.60 | 97.01 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 107.23 | 86.40 | 116.19 | 90.05 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 91.37 | 3339.97 | 147.50 | 6034.40 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 84.59 | 84.54 | 91.40 | 86.29 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 119.88 | 16728.94 | 447.60 | 29577.31 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 93.20 | 100.19 | 100.50 | 102.93 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 31.31 | 386.73 | 45.00 | 487.87 |
