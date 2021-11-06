## Performance Report

**Commit ID:** [cc8eb315135bdfed89782ee2f1368e5445df5ed3](https://github.com/aws-observability/aws-otel-collector/commit/cc8eb315135bdfed89782ee2f1368e5445df5ed3)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 59.42 | 0.20 | 59.67 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 58.32 | 0.20 | 59.40 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 58.88 | 0.20 | 58.92 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.46 | 0.20 | 58.52 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.05 | 57.04 | 0.20 | 57.09 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 57.92 | 0.20 | 58.15 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.14 | 71.90 | 0.40 | 73.64 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 0.12 | 70.75 | 0.30 | 72.40 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 0.60 | 64.88 | 0.90 | 65.58 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 57.89 | 0.20 | 58.16 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 56.51 | 0.20 | 56.81 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.16 | 0.20 | 58.25 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.45 | 0.20 | 59.66 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 57.75 | 0.20 | 58.82 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.04 | 59.12 | 0.20 | 59.15 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.61 | 0.20 | 58.92 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.30 | 113.04 | 2.40 | 118.71 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 1.18 | 116.09 | 2.00 | 121.34 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 5.30 | 65.93 | 5.70 | 66.61 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 57.55 | 0.10 | 57.59 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 58.79 | 0.20 | 59.18 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 56.56 | 0.20 | 56.65 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.63 | 0.20 | 59.11 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 57.55 | 0.30 | 57.80 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.04 | 57.20 | 0.20 | 57.49 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.31 | 0.20 | 58.75 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 6.81 | 278.01 | 11.50 | 296.99 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 6.95 | 271.90 | 12.40 | 306.56 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 26.26 | 67.12 | 27.30 | 67.50 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 57.27 | 0.20 | 57.28 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 4.74 | 155.20 | 6.30 | 218.20 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.26 | 73.60 | 2.50 | 75.60 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 6.56 | 72.22 | 7.10 | 73.65 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.31 | 87.80 | 3.70 | 88.47 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.11 | 131.00 | 4.40 | 183.70 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.03 | 68.44 | 3.30 | 68.99 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.14 | 69.83 | 4.50 | 70.59 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 4.15 | 70.87 | 4.30 | 71.86 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.79 | 82.19 | 4.00 | 82.89 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.74 | 76.47 | 6.90 | 81.48 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 35.50 | 717.08 | 46.59 | 1203.32 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.97 | 146.43 | 23.82 | 175.82 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 32.26 | 70.64 | 33.20 | 72.47 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 26.90 | 99.74 | 28.00 | 105.03 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 27.10 | 663.79 | 39.10 | 1166.34 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 26.18 | 70.36 | 27.30 | 70.94 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 34.82 | 74.18 | 44.40 | 75.50 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 36.45 | 76.03 | 45.17 | 77.58 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 25.28 | 84.29 | 26.70 | 85.04 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 29.12 | 486.07 | 36.10 | 592.26 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 55.40 | 1097.65 | 73.60 | 1819.27 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 17.77 | 164.87 | 23.21 | 195.72 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 129.29 | 79.84 | 130.41 | 80.71 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 122.21 | 121.72 | 123.62 | 131.11 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 126.04 | 3306.36 | 191.04 | 6143.54 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 120.28 | 74.08 | 121.07 | 75.90 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 153.24 | 15939.13 | 489.64 | 27066.59 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 170.07 | 12617.76 | 445.54 | 22220.33 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 113.44 | 89.18 | 115.24 | 91.77 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 28.95 | 495.76 | 36.65 | 537.85 |
