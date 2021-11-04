## Performance Report

**Commit ID:** [cc8eb315135bdfed89782ee2f1368e5445df5ed3](https://github.com/aws-observability/aws-otel-collector/commit/cc8eb315135bdfed89782ee2f1368e5445df5ed3)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.96 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 57.81 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.07 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 58.39 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.04 | 57.57 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.37 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.15 | 71.10 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 0.12 | 71.39 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 0.63 | 66.63 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 58.20 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.65 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.51 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.71 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.29 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.04 | 58.39 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 58.56 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.24 | 114.84 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 1.26 | 114.34 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 5.59 | 66.23 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 57.58 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 59.98 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 57.90 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 57.09 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.62 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.04 | 58.77 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.90 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.06 | 270.36 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 7.20 | 261.11 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 26.20 | 66.60 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 57.78 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 5.08 | 158.07 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.36 | 73.95 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 6.86 | 71.89 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.35 | 87.96 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 2.89 | 134.93 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.14 | 69.04 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.26 | 69.27 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 4.17 | 71.68 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.36 | 83.27 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.64 | 75.51 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 33.57 | 715.24 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.21 | 146.14 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 32.22 | 71.31 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 26.45 | 100.44 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 27.42 | 753.05 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 26.33 | 70.31 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 33.17 | 74.40 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 35.79 | 78.60 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 25.75 | 85.11 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 31.03 | 480.01 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 59.57 | 1094.24 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 15.92 | 165.71 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 120.42 | 79.61 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 121.16 | 121.48 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 130.82 | 3581.64 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 107.02 | 74.65 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 151.03 | 15339.10 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 168.70 | 13998.66 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 101.99 | 90.22 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 29.42 | 498.96 |
