## Performance Report

**Version:** [v0.16.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.16.0)

**Commit ID:** [373e6d60ccf0b8122fdc800efad59b5a464b5384](https://github.com/aws-observability/aws-otel-collector/commit/373e6d60ccf0b8122fdc800efad59b5a464b5384)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 65.54 | 0.30 | 66.81 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.44 | 0.20 | 63.28 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.61 | 0.20 | 62.09 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.37 | 0.20 | 62.24 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 61.46 | 0.20 | 61.62 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.76 | 0.10 | 64.07 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.13 | 74.86 | 0.40 | 75.69 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 62.49 | 0.20 | 63.24 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 65.25 | 0.40 | 67.15 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.32 | 0.20 | 62.45 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.11 | 0.20 | 63.78 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.50 | 0.20 | 63.99 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 61.89 | 0.20 | 62.16 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.04 | 0.20 | 63.11 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.23 | 114.74 | 2.90 | 119.64 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 61.27 | 0.10 | 62.88 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.93 | 0.20 | 66.23 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.28 | 0.20 | 62.95 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.53 | 0.20 | 62.30 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.86 | 0.20 | 61.92 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 61.70 | 0.20 | 62.14 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.48 | 0.20 | 62.50 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.13 | 272.54 | 12.30 | 301.60 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 60.98 | 0.20 | 61.49 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.72 | 151.05 | 6.00 | 210.56 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.30 | 79.00 | 2.60 | 81.08 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.62 | 78.04 | 5.10 | 78.72 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.25 | 93.83 | 3.60 | 95.03 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.27 | 135.01 | 4.60 | 190.66 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.38 | 72.76 | 3.70 | 73.19 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.40 | 76.19 | 4.60 | 77.53 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.15 | 87.97 | 3.40 | 88.48 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.80 | 82.70 | 6.90 | 86.56 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 26.02 | 547.35 | 35.60 | 892.73 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.16 | 155.28 | 21.80 | 182.80 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 31.31 | 79.16 | 32.49 | 80.29 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 28.10 | 106.04 | 29.60 | 110.53 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 30.19 | 687.39 | 42.81 | 1238.46 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 27.33 | 75.02 | 29.10 | 75.89 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 32.96 | 79.11 | 43.02 | 80.58 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 26.38 | 89.84 | 27.00 | 90.36 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 31.15 | 468.49 | 36.90 | 535.94 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 36.90 | 726.96 | 48.69 | 1257.87 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 17.26 | 172.25 | 23.36 | 202.76 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 116.84 | 86.20 | 119.60 | 87.24 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 123.16 | 127.14 | 125.13 | 133.87 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 129.65 | 3285.97 | 191.54 | 6097.71 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 108.53 | 79.67 | 113.69 | 80.78 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 139.93 | 15812.21 | 427.98 | 28395.57 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 109.74 | 96.32 | 110.96 | 98.77 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 30.09 | 501.32 | 35.93 | 558.81 |
