## Performance Report

**Version:** [v0.15.0-1665133649](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.15.0-1665133649)

**Commit ID:** [94d47a077e16882a635d15d9a7e296ac4e0e374d](https://github.com/aws-observability/aws-otel-collector/commit/94d47a077e16882a635d15d9a7e296ac4e0e374d)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.72 | 0.40 | 64.09 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.03 | 0.20 | 58.22 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 58.43 | 0.20 | 58.51 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.28 | 0.20 | 58.47 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.04 | 58.35 | 0.20 | 58.56 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.05 | 0.20 | 60.36 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.13 | 73.12 | 0.40 | 75.09 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 0.13 | 72.68 | 0.40 | 74.23 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 0.73 | 68.70 | 0.90 | 69.04 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 58.83 | 0.20 | 59.25 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 59.80 | 0.20 | 60.19 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.90 | 0.10 | 59.95 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.54 | 0.20 | 59.27 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.96 | 0.20 | 60.25 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.04 | 58.46 | 0.20 | 58.80 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.48 | 0.20 | 60.76 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.34 | 115.50 | 2.80 | 119.30 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 1.29 | 116.69 | 2.50 | 121.10 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 5.27 | 68.27 | 5.80 | 68.78 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 60.34 | 0.10 | 60.72 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.16 | 0.10 | 58.43 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.73 | 0.20 | 59.31 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.22 | 0.20 | 59.52 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.03 | 0.20 | 59.38 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.05 | 59.37 | 0.20 | 59.79 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.50 | 0.20 | 61.00 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 8.26 | 259.11 | 14.80 | 310.40 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 7.76 | 263.53 | 14.10 | 312.07 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 26.77 | 68.71 | 27.80 | 69.11 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 58.11 | 0.20 | 58.41 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 4.93 | 140.78 | 6.20 | 204.10 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.34 | 74.11 | 2.70 | 76.63 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 5.01 | 72.55 | 5.50 | 74.86 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.82 | 90.59 | 4.80 | 91.71 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.13 | 131.69 | 4.40 | 181.81 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.71 | 69.68 | 4.20 | 70.14 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.22 | 70.48 | 4.80 | 71.35 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 4.84 | 72.42 | 5.10 | 73.44 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 2.96 | 83.71 | 3.20 | 83.92 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.56 | 77.71 | 6.60 | 83.02 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 33.61 | 606.89 | 46.88 | 1124.56 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 15.88 | 150.54 | 21.70 | 184.98 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 33.14 | 73.77 | 34.61 | 74.63 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 30.05 | 101.30 | 31.59 | 107.14 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 27.67 | 656.84 | 38.98 | 1157.22 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 26.90 | 72.32 | 27.80 | 73.09 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 34.01 | 75.17 | 43.80 | 76.32 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 35.84 | 78.32 | 47.25 | 80.73 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 28.25 | 86.10 | 29.01 | 86.78 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 30.22 | 467.75 | 35.45 | 524.74 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 54.14 | 974.94 | 69.10 | 1746.01 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.19 | 165.45 | 22.50 | 192.18 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 128.56 | 82.28 | 129.77 | 83.71 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 116.64 | 123.15 | 119.85 | 130.54 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 113.63 | 3205.97 | 168.80 | 5785.35 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 119.30 | 77.31 | 120.68 | 78.47 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 140.25 | 14075.07 | 412.04 | 24696.95 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 174.77 | 12741.63 | 393.70 | 23129.42 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 118.33 | 91.98 | 121.56 | 94.68 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 28.44 | 503.91 | 36.75 | 557.25 |
