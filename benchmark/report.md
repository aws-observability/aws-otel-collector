## Performance Report

**Version:** [v0.27.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.27.0)

**Commit ID:** [22cb9c28bbfe78e735a8010dc5b4cc35ca7997a5](https://github.com/aws-observability/aws-otel-collector/commit/22cb9c28bbfe78e735a8010dc5b4cc35ca7997a5)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 67.33 | 0.20 | 67.57 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 68.83 | 0.20 | 69.69 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.35 | 0.20 | 68.30 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 67.97 | 0.20 | 68.26 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 67.80 | 0.20 | 67.94 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.00 | 0.20 | 68.10 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.09 | 83.30 | 0.30 | 84.60 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 67.15 | 0.10 | 67.42 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 68.92 | 0.10 | 69.11 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 68.60 | 0.20 | 68.83 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.32 | 0.20 | 68.79 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.73 | 0.10 | 69.93 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 68.61 | 0.20 | 68.72 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 68.11 | 0.20 | 68.44 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.05 | 112.21 | 1.70 | 118.62 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 67.64 | 0.20 | 68.04 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 69.42 | 0.20 | 70.03 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.51 | 0.10 | 68.65 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 69.27 | 0.10 | 69.62 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 69.51 | 0.10 | 69.71 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 67.52 | 0.20 | 67.81 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 68.75 | 0.20 | 68.88 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 5.79 | 237.22 | 9.50 | 272.00 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 68.52 | 0.10 | 68.68 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.75 | 82.07 | 3.90 | 83.23 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.89 | 88.57 | 15.90 | 90.45 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 3.47 | 86.02 | 3.70 | 88.90 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.71 | 83.90 | 3.90 | 84.99 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.19 | 142.62 | 4.30 | 193.15 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 4.07 | 82.92 | 4.70 | 84.50 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.37 | 83.26 | 4.60 | 84.41 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 4.15 | 94.92 | 4.30 | 95.29 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 4.61 | 85.53 | 17.00 | 87.36 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.21 | 86.38 | 20.10 | 88.47 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 23.96 | 153.60 | 42.70 | 177.20 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 30.77 | 92.41 | 33.10 | 94.45 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 31.16 | 83.78 | 34.20 | 86.17 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 25.47 | 696.49 | 36.70 | 1234.12 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 27.47 | 83.70 | 28.60 | 85.47 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 27.66 | 85.24 | 29.60 | 87.43 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 24.56 | 97.55 | 25.20 | 98.19 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 32.30 | 279.30 | 46.40 | 399.97 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 26.65 | 96.83 | 27.80 | 103.75 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 25.20 | 178.71 | 43.10 | 208.03 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 106.55 | 100.55 | 115.60 | 109.88 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 104.01 | 87.89 | 108.90 | 90.72 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 98.99 | 3172.00 | 155.29 | 5853.22 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 91.14 | 85.80 | 98.91 | 86.99 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 112.55 | 17510.98 | 450.72 | 29223.18 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 89.69 | 99.43 | 96.30 | 101.61 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 32.68 | 378.42 | 52.00 | 445.54 |
