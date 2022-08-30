## Performance Report

**Version:** [v0.21.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.21.0)

**Commit ID:** [bf025291620c9eb0600a16ec462110b6e83f7956](https://github.com/aws-observability/aws-otel-collector/commit/bf025291620c9eb0600a16ec462110b6e83f7956)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 64.35 | 0.10 | 65.01 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.18 | 0.20 | 63.33 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.94 | 0.20 | 64.48 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 64.09 | 0.20 | 64.59 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 63.58 | 0.20 | 64.19 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 65.02 | 0.20 | 65.38 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.10 | 78.77 | 0.30 | 80.20 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 63.13 | 0.10 | 63.54 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 67.63 | 0.20 | 69.79 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.15 | 0.20 | 63.63 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.24 | 0.20 | 64.58 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 63.60 | 0.10 | 63.81 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 63.14 | 0.20 | 63.45 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.65 | 0.20 | 65.05 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.14 | 112.48 | 1.90 | 116.53 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 62.30 | 0.20 | 63.37 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 63.33 | 0.20 | 63.42 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.93 | 0.10 | 65.85 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 63.56 | 0.20 | 63.71 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.03 | 0.20 | 63.23 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 64.20 | 0.20 | 64.98 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.37 | 0.20 | 65.72 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 6.83 | 245.21 | 10.60 | 273.87 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 63.28 | 0.20 | 63.75 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.16 | 82.33 | 15.50 | 86.07 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 3.82 | 81.56 | 4.00 | 84.58 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 4.08 | 79.06 | 4.30 | 80.56 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 2.91 | 137.26 | 3.90 | 184.61 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.92 | 76.72 | 4.10 | 78.69 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.93 | 76.22 | 4.20 | 77.57 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.30 | 89.93 | 3.90 | 91.60 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 6.01 | 81.80 | 19.30 | 85.90 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 26.62 | 155.95 | 40.10 | 185.93 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 29.84 | 83.47 | 30.30 | 87.09 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 30.01 | 79.64 | 30.70 | 81.76 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 27.33 | 710.62 | 37.80 | 1270.82 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 27.31 | 78.05 | 28.20 | 80.11 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 31.47 | 80.67 | 41.80 | 84.07 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 26.37 | 91.22 | 27.20 | 92.75 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 41.22 | 310.06 | 58.20 | 429.27 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 27.05 | 176.58 | 44.60 | 206.79 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 120.83 | 85.96 | 132.91 | 89.99 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 124.34 | 82.38 | 133.99 | 86.09 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 105.61 | 3392.73 | 170.41 | 5591.39 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 101.57 | 80.29 | 102.61 | 81.67 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 103.04 | 20453.26 | 326.92 | 32670.13 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 111.30 | 97.38 | 126.70 | 100.39 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 42.41 | 445.69 | 70.81 | 621.90 |
