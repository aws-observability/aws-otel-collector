## Performance Report

**Commit ID:** [cc8eb315135bdfed89782ee2f1368e5445df5ed3](https://github.com/aws-observability/aws-otel-collector/commit/cc8eb315135bdfed89782ee2f1368e5445df5ed3)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 60.88 | 0.30 | 62.72 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.03 | 0.10 | 58.18 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 57.71 | 0.20 | 57.81 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.44 | 0.20 | 59.11 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.06 | 58.59 | 0.20 | 58.59 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.44 | 0.20 | 58.68 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.12 | 71.67 | 0.40 | 73.13 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 0.13 | 71.87 | 0.40 | 73.41 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 0.58 | 65.12 | 0.80 | 65.75 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 56.75 | 0.20 | 57.09 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.95 | 0.20 | 60.67 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.68 | 0.20 | 59.08 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.60 | 0.20 | 59.51 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 57.21 | 0.20 | 57.45 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.04 | 57.71 | 0.20 | 58.01 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 58.49 | 0.20 | 58.73 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.21 | 113.98 | 3.00 | 119.64 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 1.23 | 113.08 | 2.70 | 118.13 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 5.76 | 65.26 | 6.10 | 66.07 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 58.14 | 0.10 | 58.27 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 59.68 | 0.40 | 61.78 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.41 | 0.20 | 58.41 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 57.88 | 0.20 | 57.90 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 57.75 | 0.20 | 57.90 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.04 | 58.49 | 0.30 | 58.94 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 58.41 | 0.20 | 58.88 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.25 | 267.15 | 11.90 | 293.19 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 7.27 | 274.62 | 13.20 | 309.77 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 26.71 | 66.44 | 28.10 | 66.98 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 57.99 | 0.20 | 58.38 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 5.08 | 163.49 | 6.30 | 222.15 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.30 | 75.09 | 2.70 | 77.10 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 5.00 | 70.75 | 5.40 | 73.04 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.18 | 89.09 | 3.60 | 91.10 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.18 | 129.17 | 4.30 | 177.35 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.14 | 69.58 | 3.60 | 70.39 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.39 | 68.81 | 4.80 | 69.53 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 4.95 | 71.34 | 5.40 | 72.24 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 2.90 | 82.61 | 3.10 | 82.72 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.59 | 77.54 | 7.00 | 82.44 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 29.39 | 714.83 | 42.09 | 1124.94 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.79 | 151.69 | 21.30 | 181.53 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 30.05 | 74.65 | 31.51 | 75.50 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 26.52 | 100.28 | 27.80 | 105.58 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 29.28 | 682.60 | 42.80 | 1221.44 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 26.47 | 70.07 | 27.61 | 71.12 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 34.14 | 75.24 | 43.60 | 76.54 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 37.27 | 78.45 | 46.75 | 86.13 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 25.87 | 84.57 | 27.50 | 85.17 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 28.93 | 475.99 | 35.84 | 533.16 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 46.56 | 942.80 | 67.52 | 1625.97 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.22 | 166.94 | 24.60 | 194.13 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 118.99 | 81.32 | 120.29 | 82.04 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 122.62 | 122.51 | 123.76 | 129.04 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 108.57 | 3216.56 | 177.05 | 5924.95 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 115.66 | 75.76 | 116.59 | 76.71 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 146.75 | 13927.55 | 376.01 | 27553.96 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 144.75 | 17884.46 | 504.26 | 28793.83 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 116.38 | 90.39 | 118.04 | 93.61 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 29.75 | 506.68 | 38.20 | 553.05 |
