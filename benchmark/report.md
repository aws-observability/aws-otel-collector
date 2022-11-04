## Performance Report

**Version:** [v0.22.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.22.0)

**Commit ID:** [0818a9a5a62ffeba69d8a81c2ddf446ffbbcfbf5](https://github.com/aws-observability/aws-otel-collector/commit/0818a9a5a62ffeba69d8a81c2ddf446ffbbcfbf5)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.50 | 0.10 | 69.75 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.03 | 0.20 | 65.33 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.73 | 0.20 | 65.03 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.60 | 0.20 | 64.00 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 64.19 | 0.20 | 64.52 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.58 | 0.20 | 65.88 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.09 | 79.75 | 0.20 | 80.62 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 64.29 | 0.20 | 64.44 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.42 | 0.20 | 68.83 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.84 | 0.20 | 65.99 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.47 | 0.20 | 65.69 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.10 | 0.10 | 65.80 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 64.61 | 0.20 | 65.14 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.67 | 0.20 | 66.06 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.02 | 105.99 | 1.70 | 112.83 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 63.81 | 0.20 | 64.24 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.68 | 0.20 | 64.97 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.18 | 0.20 | 66.45 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.02 | 0.20 | 65.77 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.26 | 0.20 | 65.70 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 63.97 | 0.20 | 64.09 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 66.77 | 0.20 | 66.95 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 5.50 | 221.40 | 9.00 | 250.52 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 64.14 | 0.10 | 64.40 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.78 | 77.01 | 4.00 | 78.38 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.00 | 84.41 | 15.00 | 87.05 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 3.72 | 82.78 | 4.00 | 85.83 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 4.60 | 78.98 | 5.40 | 80.97 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.48 | 139.51 | 5.10 | 186.82 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.84 | 79.12 | 4.30 | 80.75 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.50 | 78.27 | 5.00 | 80.22 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.19 | 91.48 | 3.80 | 93.13 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.55 | 84.09 | 18.70 | 87.58 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 20.34 | 82.63 | 21.30 | 85.89 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 24.36 | 155.20 | 40.40 | 192.20 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 30.09 | 85.45 | 31.50 | 89.28 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 33.77 | 80.71 | 34.70 | 82.89 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 28.17 | 709.61 | 39.70 | 1248.95 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 26.96 | 79.78 | 28.50 | 81.33 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 28.32 | 81.92 | 29.40 | 84.71 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 25.73 | 92.64 | 26.20 | 94.26 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 40.73 | 327.14 | 59.70 | 478.08 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 27.70 | 94.47 | 29.50 | 102.60 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 24.25 | 182.83 | 42.40 | 211.99 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 124.04 | 94.98 | 125.20 | 102.99 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 125.58 | 83.42 | 127.80 | 85.92 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 112.48 | 3255.42 | 166.09 | 5844.52 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 111.56 | 80.93 | 113.01 | 84.87 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 110.27 | 17128.18 | 328.61 | 29354.08 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 105.65 | 97.90 | 108.00 | 99.90 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 37.42 | 435.39 | 56.30 | 650.65 |
