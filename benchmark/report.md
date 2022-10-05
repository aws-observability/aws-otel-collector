## Performance Report

**Version:** [v0.22.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.22.0)

**Commit ID:** [14a2d975e96a6b3fa712f83577274fe4c4a66d5e](https://github.com/aws-observability/aws-otel-collector/commit/14a2d975e96a6b3fa712f83577274fe4c4a66d5e)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 68.60 | 0.20 | 70.71 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.56 | 0.20 | 66.01 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.37 | 0.20 | 65.82 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.46 | 0.20 | 64.68 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 65.50 | 0.20 | 65.77 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.75 | 0.20 | 66.46 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.09 | 81.17 | 0.30 | 82.40 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 66.30 | 0.20 | 66.64 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.42 | 0.10 | 69.50 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 64.77 | 0.20 | 64.81 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 65.94 | 0.10 | 66.65 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.67 | 0.20 | 65.71 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 64.79 | 0.20 | 65.10 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.54 | 0.20 | 66.62 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.96 | 106.57 | 1.70 | 111.71 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 64.16 | 0.20 | 64.69 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 67.89 | 0.20 | 69.10 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 66.44 | 0.20 | 66.66 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.33 | 0.20 | 65.63 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.59 | 0.20 | 64.95 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 64.86 | 0.20 | 65.51 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.19 | 0.20 | 66.70 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 5.57 | 231.05 | 9.10 | 253.35 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 64.05 | 0.20 | 64.68 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.15 | 77.64 | 5.10 | 79.84 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.93 | 85.19 | 15.00 | 88.06 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 3.72 | 82.65 | 4.00 | 85.51 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 4.67 | 78.55 | 5.20 | 80.31 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.22 | 140.09 | 4.60 | 192.48 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 4.59 | 79.97 | 4.90 | 81.96 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.40 | 78.35 | 5.00 | 80.20 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 4.07 | 92.46 | 4.60 | 93.86 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.44 | 84.74 | 18.40 | 90.13 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.60 | 81.93 | 21.00 | 84.90 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 24.55 | 157.20 | 35.00 | 202.41 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 30.35 | 88.26 | 32.00 | 90.66 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 30.76 | 80.73 | 31.60 | 83.42 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 27.19 | 746.90 | 38.40 | 1245.40 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 26.84 | 79.53 | 27.30 | 80.88 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 28.67 | 83.20 | 29.20 | 86.41 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 25.38 | 93.42 | 25.90 | 95.13 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 41.14 | 298.43 | 58.10 | 479.01 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 26.44 | 93.61 | 28.90 | 100.84 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 23.39 | 184.38 | 30.90 | 213.15 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 124.15 | 92.51 | 132.00 | 95.12 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 130.22 | 83.35 | 131.00 | 86.97 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 119.89 | 3428.98 | 180.31 | 5864.64 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 108.32 | 82.22 | 110.70 | 83.92 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 109.99 | 14418.54 | 252.18 | 28657.59 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 110.09 | 97.69 | 111.20 | 100.34 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 37.76 | 412.57 | 56.80 | 580.67 |
