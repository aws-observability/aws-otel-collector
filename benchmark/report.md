## Performance Report

**Version:** [v0.24.1](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.24.1)

**Commit ID:** [e9cfb1059ae49444e1211cab96e46ee8fb910a6c](https://github.com/aws-observability/aws-otel-collector/commit/e9cfb1059ae49444e1211cab96e46ee8fb910a6c)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.97 | 0.20 | 66.56 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.45 | 0.20 | 64.70 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.92 | 0.20 | 66.74 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.77 | 0.20 | 67.04 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 66.09 | 0.20 | 67.03 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.65 | 0.20 | 67.19 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.09 | 81.66 | 0.30 | 82.72 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 64.31 | 0.20 | 65.20 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 68.30 | 0.20 | 69.50 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.90 | 0.20 | 65.97 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.62 | 0.20 | 66.96 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.71 | 0.20 | 66.51 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 66.31 | 0.20 | 66.61 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.01 | 0.20 | 67.52 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.95 | 107.70 | 1.60 | 112.71 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 65.49 | 0.10 | 65.85 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.73 | 0.20 | 70.61 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.94 | 0.20 | 65.85 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.84 | 0.20 | 65.26 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.06 | 0.10 | 66.45 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 66.39 | 0.20 | 66.90 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.11 | 0.10 | 66.44 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 5.82 | 219.73 | 9.40 | 254.42 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 66.29 | 0.10 | 66.70 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.70 | 78.25 | 4.00 | 79.74 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.05 | 85.37 | 15.80 | 87.80 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 3.67 | 83.85 | 4.40 | 87.33 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 4.04 | 79.38 | 4.30 | 81.99 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.12 | 142.89 | 4.20 | 185.71 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.87 | 80.22 | 4.20 | 82.33 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.07 | 79.68 | 4.30 | 81.05 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.11 | 91.35 | 3.30 | 93.09 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.19 | 85.13 | 5.60 | 89.26 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.43 | 82.90 | 20.00 | 86.23 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 26.72 | 152.66 | 39.70 | 186.66 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 27.82 | 88.70 | 28.60 | 91.38 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 30.99 | 80.78 | 33.80 | 83.13 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 27.65 | 764.40 | 38.80 | 1268.96 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 26.17 | 78.68 | 26.70 | 81.25 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 28.34 | 82.08 | 28.80 | 84.58 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 24.34 | 93.54 | 25.00 | 95.03 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 39.09 | 318.36 | 55.50 | 507.87 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 26.25 | 97.39 | 27.40 | 104.37 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 24.28 | 180.03 | 39.40 | 210.68 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 120.35 | 95.92 | 122.41 | 98.52 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 116.97 | 83.58 | 119.39 | 87.72 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 113.80 | 3407.20 | 174.40 | 5820.99 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 110.22 | 82.88 | 110.90 | 84.28 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 106.23 | 17299.42 | 300.59 | 27469.30 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 108.99 | 97.82 | 110.50 | 99.86 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 36.49 | 427.76 | 56.70 | 519.19 |
