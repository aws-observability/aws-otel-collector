## Performance Report

**Commit ID:** [6306f81c2fa65cce889149df1c5b775e50775b2a](https://github.com/aws-observability/aws-otel-collector/commit/6306f81c2fa65cce889149df1c5b775e50775b2a)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.83 | 0.30 | 64.05 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.99 | 0.20 | 59.20 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.28 | 0.20 | 59.99 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 58.37 | 0.20 | 58.62 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.04 | 58.90 | 0.20 | 59.05 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.56 | 0.20 | 59.27 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.13 | 71.32 | 0.40 | 72.61 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 0.13 | 72.39 | 0.40 | 73.61 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 0.67 | 66.82 | 0.90 | 67.43 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.06 | 58.88 | 0.20 | 59.15 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.11 | 0.40 | 63.34 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.56 | 0.20 | 58.74 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.30 | 0.20 | 59.68 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 59.38 | 0.20 | 59.86 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.05 | 58.15 | 0.20 | 58.25 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.20 | 0.20 | 59.60 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.23 | 114.72 | 2.40 | 120.74 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 1.22 | 114.92 | 2.70 | 119.64 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 5.50 | 67.17 | 6.00 | 67.76 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 59.64 | 0.10 | 59.76 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.48 | 0.30 | 63.53 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.89 | 0.20 | 59.13 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.94 | 0.20 | 60.52 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.20 | 0.20 | 58.46 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.04 | 59.63 | 0.20 | 60.36 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.26 | 0.20 | 60.38 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.26 | 273.35 | 13.70 | 293.58 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 6.82 | 280.98 | 11.30 | 292.55 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 26.39 | 67.91 | 27.69 | 68.26 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 59.35 | 0.20 | 59.56 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 4.56 | 142.60 | 5.90 | 206.99 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.37 | 74.13 | 2.70 | 75.72 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 5.83 | 72.96 | 6.60 | 74.72 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.42 | 88.16 | 3.90 | 88.84 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 2.88 | 126.30 | 4.30 | 179.29 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.39 | 69.30 | 3.60 | 69.71 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.99 | 70.78 | 4.60 | 71.79 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 5.31 | 72.04 | 5.60 | 73.34 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.13 | 83.88 | 3.30 | 84.05 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.43 | 76.17 | 6.60 | 80.14 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 31.35 | 571.19 | 42.29 | 986.86 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 17.23 | 149.72 | 22.60 | 180.76 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 31.61 | 73.73 | 32.31 | 74.86 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 27.60 | 101.18 | 28.50 | 105.62 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 28.28 | 658.23 | 41.92 | 1163.26 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 28.97 | 70.74 | 29.79 | 71.54 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 34.53 | 76.25 | 44.69 | 77.38 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 35.93 | 77.61 | 46.28 | 79.54 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 26.49 | 85.00 | 27.30 | 85.77 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 30.67 | 477.05 | 35.92 | 558.60 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 50.30 | 849.34 | 70.32 | 1485.22 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 15.87 | 167.25 | 23.40 | 201.14 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 131.07 | 82.09 | 136.12 | 83.50 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 118.86 | 123.59 | 120.30 | 131.09 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 123.19 | 3360.09 | 189.65 | 6049.49 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 122.79 | 76.54 | 124.70 | 78.01 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 145.06 | 16417.05 | 478.15 | 28533.93 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 153.77 | 15337.49 | 419.97 | 26340.54 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 112.55 | 90.70 | 113.80 | 93.22 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 28.49 | 511.65 | 36.23 | 557.35 |
