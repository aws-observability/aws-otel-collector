## Performance Report

**Version:** [v0.16.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.16.0)

**Commit ID:** [26b6278b054f276424cf33be3cd33f4a25d6c1ec](https://github.com/aws-observability/aws-otel-collector/commit/26b6278b054f276424cf33be3cd33f4a25d6c1ec)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 64.30 | 0.30 | 65.04 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 61.73 | 0.20 | 61.83 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 64.21 | 0.30 | 64.75 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.01 | 0.20 | 62.39 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 61.83 | 0.20 | 61.86 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.53 | 0.20 | 63.72 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.12 | 76.02 | 0.40 | 77.09 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 62.31 | 0.20 | 62.64 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 61.83 | 0.30 | 62.33 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.92 | 0.20 | 62.40 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.65 | 0.20 | 63.84 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.66 | 0.20 | 62.11 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 62.41 | 0.20 | 63.01 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.08 | 0.20 | 63.59 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.42 | 113.22 | 2.80 | 119.21 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 63.50 | 0.20 | 63.77 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.79 | 0.30 | 65.88 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.23 | 0.20 | 63.88 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.54 | 0.20 | 62.83 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.48 | 0.20 | 63.18 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 61.18 | 0.20 | 61.32 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.29 | 0.20 | 63.65 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 8.01 | 264.99 | 15.41 | 297.76 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 61.43 | 0.20 | 61.49 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 5.10 | 144.15 | 5.90 | 211.57 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.36 | 80.18 | 2.80 | 82.49 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 6.29 | 77.43 | 7.10 | 78.20 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.59 | 94.67 | 4.00 | 95.45 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 5.49 | 128.99 | 7.40 | 186.27 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.53 | 73.64 | 3.90 | 74.34 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 5.05 | 74.04 | 5.40 | 74.85 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.61 | 87.84 | 3.90 | 88.08 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 6.11 | 81.23 | 7.10 | 85.57 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 28.89 | 490.39 | 39.81 | 857.53 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.31 | 153.29 | 21.50 | 183.99 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 32.42 | 77.78 | 34.50 | 80.06 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 28.17 | 106.44 | 30.90 | 110.79 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 33.98 | 768.85 | 48.30 | 1241.67 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 27.74 | 74.78 | 29.89 | 75.22 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 32.58 | 78.76 | 42.50 | 79.88 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 28.93 | 89.04 | 36.09 | 89.42 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 29.28 | 451.79 | 36.13 | 550.08 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 38.66 | 690.39 | 52.30 | 1225.64 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 17.28 | 169.16 | 22.82 | 193.38 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 122.77 | 86.44 | 125.73 | 87.56 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 122.42 | 126.38 | 124.07 | 132.06 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 150.25 | 3613.21 | 207.50 | 6268.38 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 124.23 | 79.66 | 127.80 | 80.36 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 165.52 | 13882.20 | 380.74 | 25364.01 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 127.59 | 94.51 | 136.07 | 97.64 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 28.77 | 498.24 | 36.40 | 548.52 |
