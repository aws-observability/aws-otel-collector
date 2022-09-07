## Performance Report

**Version:** [v0.21.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.21.0)

**Commit ID:** [ead98b319b36e5a12cbe53ac79242efe60472d78](https://github.com/aws-observability/aws-otel-collector/commit/ead98b319b36e5a12cbe53ac79242efe60472d78)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 64.43 | 0.20 | 64.63 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 64.23 | 0.20 | 64.45 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 65.04 | 0.20 | 65.28 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 63.60 | 0.20 | 63.98 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 64.20 | 0.20 | 64.51 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 66.07 | 0.20 | 66.47 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.10 | 78.58 | 0.20 | 80.22 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 63.86 | 0.20 | 64.41 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 65.80 | 0.10 | 66.53 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 64.47 | 0.20 | 65.08 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.01 | 0.10 | 64.46 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 62.39 | 0.20 | 62.98 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 64.87 | 0.20 | 64.97 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.12 | 0.20 | 66.14 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.10 | 109.89 | 1.80 | 113.95 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 64.16 | 0.20 | 64.61 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 67.55 | 0.20 | 68.44 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 63.21 | 0.20 | 63.60 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.65 | 0.20 | 65.01 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.50 | 0.20 | 63.82 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 63.13 | 0.20 | 63.32 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 65.30 | 0.20 | 65.41 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 6.53 | 244.85 | 10.70 | 273.53 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 62.96 | 0.10 | 63.26 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.99 | 77.09 | 4.30 | 78.57 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.94 | 82.54 | 15.50 | 84.68 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 3.64 | 80.46 | 4.00 | 82.73 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.96 | 78.07 | 4.30 | 80.00 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.17 | 136.57 | 4.40 | 187.86 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.85 | 77.86 | 4.50 | 79.46 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.54 | 77.47 | 3.90 | 78.79 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.24 | 90.02 | 3.60 | 91.32 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.65 | 84.16 | 19.00 | 89.06 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 20.00 | 81.33 | 20.60 | 85.14 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 25.90 | 158.62 | 42.60 | 187.94 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 29.45 | 92.16 | 30.20 | 96.80 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 29.83 | 79.63 | 30.40 | 81.71 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 26.80 | 708.83 | 37.20 | 1226.71 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 26.35 | 79.08 | 26.80 | 80.24 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 30.36 | 80.22 | 31.00 | 83.35 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 25.45 | 91.73 | 26.50 | 93.59 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 40.56 | 331.97 | 60.60 | 517.75 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 27.76 | 94.32 | 29.40 | 101.25 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 24.84 | 180.88 | 42.30 | 211.55 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 127.81 | 90.20 | 129.80 | 101.31 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 120.51 | 82.80 | 125.40 | 86.15 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 109.94 | 3067.57 | 171.11 | 5642.21 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 114.23 | 80.86 | 115.80 | 82.47 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 98.51 | 21387.25 | 275.30 | 32665.55 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 117.79 | 97.97 | 119.00 | 99.98 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 39.04 | 440.45 | 55.90 | 628.11 |
