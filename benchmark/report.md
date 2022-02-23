## Performance Report

**Version:** [v0.17.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.17.0)

**Commit ID:** [14833d4af543da709c77cf9dc6827351dbd529b1](https://github.com/aws-observability/aws-otel-collector/commit/14833d4af543da709c77cf9dc6827351dbd529b1)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 65.93 | 0.40 | 67.08 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.46 | 0.20 | 61.55 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.40 | 0.20 | 62.46 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 62.03 | 0.20 | 62.76 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 62.23 | 0.20 | 62.46 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.05 | 0.10 | 64.16 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.14 | 75.80 | 0.50 | 77.44 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 63.25 | 0.20 | 63.36 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.66 | 0.30 | 65.17 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.70 | 0.30 | 63.27 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 63.46 | 0.20 | 64.05 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.03 | 0.20 | 63.48 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 62.18 | 0.30 | 62.91 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 62.53 | 0.20 | 63.25 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.40 | 113.77 | 3.10 | 119.34 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 61.13 | 0.20 | 61.23 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.55 | 0.20 | 61.73 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.85 | 0.20 | 62.91 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.31 | 0.30 | 64.21 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 61.97 | 0.20 | 62.43 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 62.58 | 0.20 | 62.66 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 64.18 | 0.40 | 64.61 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.79 | 273.06 | 14.10 | 299.76 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 62.74 | 0.20 | 63.50 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.87 | 120.14 | 5.70 | 152.24 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.20 | 79.48 | 3.70 | 81.71 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 5.53 | 77.29 | 6.60 | 77.90 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.83 | 92.77 | 6.40 | 93.81 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.28 | 138.13 | 4.90 | 184.50 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 5.04 | 74.95 | 6.00 | 76.12 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.65 | 75.81 | 5.20 | 76.77 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 5.35 | 87.23 | 6.50 | 87.70 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 7.03 | 79.80 | 8.40 | 83.26 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.33 | 306.15 | 33.80 | 505.79 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 20.10 | 156.29 | 26.59 | 187.82 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 42.87 | 77.58 | 66.91 | 79.44 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 28.13 | 107.11 | 29.40 | 111.92 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 27.68 | 651.66 | 40.58 | 1213.69 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 42.03 | 76.93 | 56.89 | 77.42 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 44.86 | 79.10 | 55.38 | 80.31 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 28.89 | 90.45 | 30.50 | 91.15 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 32.97 | 484.15 | 39.75 | 555.67 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 36.77 | 479.52 | 47.02 | 748.63 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 19.68 | 171.53 | 27.90 | 200.01 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 127.61 | 86.40 | 136.61 | 87.51 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 125.71 | 128.28 | 144.96 | 136.19 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 129.31 | 3573.14 | 195.35 | 6135.81 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 116.63 | 80.19 | 126.38 | 81.06 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 150.02 | 13800.29 | 479.50 | 25375.10 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 107.70 | 96.00 | 112.17 | 97.93 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 33.67 | 510.41 | 40.95 | 626.01 |
