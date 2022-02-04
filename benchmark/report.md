## Performance Report

**Version:** [v0.15.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.15.0)

**Commit ID:** [c24acacb322cde52d015881dec5086831d9ed5ab](https://github.com/aws-observability/aws-otel-collector/commit/c24acacb322cde52d015881dec5086831d9ed5ab)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.84 | 0.20 | 62.70 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.60 | 0.20 | 59.08 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 57.87 | 0.20 | 58.09 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.51 | 0.20 | 58.57 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 59.97 | 0.20 | 60.28 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.85 | 0.20 | 60.34 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.14 | 72.99 | 0.40 | 74.18 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 60.42 | 0.20 | 61.19 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.00 | 0.20 | 62.83 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.59 | 0.20 | 60.09 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.17 | 0.20 | 60.32 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 59.26 | 0.20 | 59.62 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 60.30 | 0.20 | 60.78 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.73 | 0.30 | 60.96 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.47 | 108.47 | 3.00 | 115.54 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 59.57 | 0.20 | 60.05 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.17 | 0.30 | 63.23 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.13 | 0.20 | 59.62 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.03 | 0.20 | 60.62 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.78 | 0.20 | 59.09 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 58.30 | 0.20 | 58.41 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.59 | 0.20 | 62.02 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.76 | 255.06 | 14.80 | 283.53 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 59.22 | 0.10 | 59.73 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.58 | 154.90 | 5.80 | 209.75 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.46 | 75.09 | 2.80 | 77.12 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 5.10 | 72.71 | 5.60 | 74.72 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.42 | 90.38 | 3.80 | 90.72 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.57 | 132.35 | 5.30 | 188.79 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 2.99 | 70.36 | 3.20 | 70.80 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.03 | 71.59 | 4.40 | 72.63 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.28 | 85.12 | 3.50 | 85.68 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.54 | 79.38 | 6.80 | 84.07 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.57 | 494.13 | 34.18 | 862.87 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.78 | 150.03 | 23.50 | 179.86 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 31.88 | 75.37 | 32.80 | 76.78 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 26.69 | 102.51 | 27.29 | 106.68 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 28.41 | 665.55 | 41.91 | 1225.94 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 26.24 | 71.88 | 27.00 | 72.80 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 33.69 | 76.88 | 43.41 | 78.46 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 27.93 | 87.40 | 28.70 | 87.94 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 29.97 | 465.68 | 36.88 | 553.76 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 36.80 | 703.24 | 48.58 | 1218.67 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 17.12 | 167.10 | 22.34 | 207.08 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 126.59 | 81.83 | 133.02 | 83.47 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 114.44 | 124.58 | 116.00 | 131.08 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 121.32 | 3153.14 | 185.31 | 5713.75 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 118.30 | 76.88 | 120.03 | 78.18 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 144.85 | 16136.62 | 448.54 | 28658.21 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 107.65 | 91.01 | 114.22 | 94.00 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 30.38 | 507.57 | 37.80 | 572.54 |
