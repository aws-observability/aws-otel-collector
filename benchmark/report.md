## Performance Report

**Version:** [v0.24.1](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.24.1)

**Commit ID:** [ba14e5befdc216cd6b030d9297ae579cc1b60033](https://github.com/aws-observability/aws-otel-collector/commit/ba14e5befdc216cd6b030d9297ae579cc1b60033)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 70.65 | 0.20 | 72.59 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.99 | 0.20 | 66.02 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 65.90 | 0.20 | 66.26 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 66.23 | 0.20 | 66.68 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 66.17 | 0.20 | 66.59 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.44 | 0.20 | 66.65 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.09 | 81.67 | 0.30 | 82.95 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 66.07 | 0.10 | 66.75 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 69.31 | 0.20 | 70.61 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.14 | 0.20 | 67.31 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 65.67 | 0.20 | 66.18 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.86 | 0.20 | 66.10 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 66.73 | 0.20 | 67.44 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.82 | 0.20 | 67.45 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.04 | 111.27 | 1.80 | 116.00 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 66.99 | 0.10 | 67.18 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.18 | 0.20 | 70.84 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.01 | 0.20 | 65.56 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 65.80 | 0.20 | 65.82 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 66.37 | 0.20 | 66.74 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 66.07 | 0.20 | 66.51 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.12 | 0.20 | 67.48 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 5.72 | 237.67 | 9.70 | 262.50 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 65.70 | 0.10 | 66.11 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.04 | 80.16 | 4.30 | 81.82 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.05 | 86.12 | 15.60 | 88.37 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 3.57 | 82.52 | 4.00 | 86.43 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 4.14 | 81.05 | 4.70 | 83.62 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.10 | 141.95 | 4.50 | 193.64 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 4.47 | 80.15 | 4.70 | 81.72 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.41 | 79.74 | 3.60 | 81.84 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 4.13 | 93.67 | 4.40 | 95.67 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.43 | 85.53 | 17.90 | 90.56 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.64 | 83.66 | 20.20 | 86.50 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 25.17 | 149.97 | 43.60 | 173.79 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 29.90 | 91.12 | 30.80 | 100.04 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 28.29 | 80.81 | 28.70 | 83.10 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 25.53 | 738.58 | 37.00 | 1203.82 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 24.74 | 80.98 | 25.20 | 83.12 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 26.95 | 82.71 | 27.50 | 85.54 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 23.84 | 94.75 | 24.20 | 96.89 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 38.61 | 321.60 | 55.20 | 491.52 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 27.11 | 96.40 | 29.10 | 103.69 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 25.09 | 180.25 | 43.50 | 210.75 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 125.31 | 93.46 | 128.11 | 97.60 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 118.44 | 85.45 | 121.30 | 88.81 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 108.69 | 3413.32 | 168.89 | 5996.36 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 101.79 | 83.31 | 103.20 | 86.01 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 106.17 | 12340.59 | 337.99 | 27284.87 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 99.50 | 99.50 | 101.91 | 100.74 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 36.78 | 449.00 | 54.60 | 669.13 |
