## Performance Report

**Version:** [v0.16.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.16.0)

**Commit ID:** [2c6847f8175e08f1700fea50d3caf5607c5b7b6d](https://github.com/aws-observability/aws-otel-collector/commit/2c6847f8175e08f1700fea50d3caf5607c5b7b6d)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 65.97 | 0.40 | 67.46 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.26 | 0.20 | 61.61 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.13 | 0.20 | 61.29 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 60.86 | 0.20 | 61.44 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 62.72 | 0.20 | 62.91 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.06 | 0.20 | 63.55 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.12 | 75.19 | 0.30 | 76.50 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 62.63 | 0.20 | 62.84 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 65.17 | 0.40 | 66.22 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.64 | 0.20 | 64.19 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.09 | 0.30 | 62.55 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.90 | 0.20 | 62.85 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 62.37 | 0.20 | 62.66 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.96 | 0.20 | 64.30 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.29 | 115.19 | 2.40 | 120.89 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 61.92 | 0.20 | 62.42 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 65.17 | 0.50 | 66.15 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.42 | 0.20 | 64.38 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.52 | 0.20 | 63.23 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.94 | 0.20 | 61.58 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 63.10 | 0.20 | 63.42 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.09 | 0.20 | 63.48 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.24 | 270.92 | 13.40 | 300.86 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 61.46 | 0.20 | 61.80 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.59 | 142.49 | 5.70 | 207.95 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.32 | 79.31 | 2.80 | 80.81 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.36 | 78.36 | 4.70 | 79.04 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.27 | 95.83 | 3.60 | 98.82 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.17 | 139.72 | 4.40 | 193.14 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.53 | 73.76 | 3.80 | 74.20 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.13 | 74.52 | 4.50 | 75.45 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.39 | 88.20 | 3.70 | 88.35 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.79 | 81.59 | 6.91 | 86.07 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.93 | 479.30 | 33.54 | 842.49 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.93 | 154.36 | 22.43 | 187.93 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 30.72 | 79.22 | 31.70 | 81.42 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 27.48 | 105.64 | 28.70 | 110.84 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 28.09 | 761.83 | 40.80 | 1295.46 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 25.81 | 74.56 | 26.69 | 75.36 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 34.07 | 79.07 | 44.76 | 80.05 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 26.07 | 90.75 | 27.99 | 91.46 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 30.83 | 459.18 | 36.40 | 529.40 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 39.61 | 666.31 | 56.52 | 1221.69 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.77 | 169.88 | 23.30 | 194.93 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 134.52 | 85.98 | 135.58 | 87.25 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 117.16 | 127.27 | 118.39 | 134.75 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 129.98 | 3660.10 | 199.48 | 6285.22 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 111.45 | 80.75 | 116.80 | 81.55 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 151.21 | 16196.41 | 421.49 | 28242.28 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 102.43 | 96.17 | 106.13 | 98.46 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 28.51 | 509.96 | 36.47 | 561.48 |
