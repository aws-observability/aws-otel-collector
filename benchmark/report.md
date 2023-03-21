## Performance Report

**Version:** [v0.27.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.27.0)

**Commit ID:** [4ab9f5daec88eca35d6362dfc59f64ae5c37e49e](https://github.com/aws-observability/aws-otel-collector/commit/4ab9f5daec88eca35d6362dfc59f64ae5c37e49e)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 69.75 | 0.20 | 70.39 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.06 | 0.20 | 68.46 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.32 | 0.20 | 68.65 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 66.42 | 0.20 | 67.26 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 68.70 | 0.10 | 68.76 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 69.14 | 0.20 | 69.17 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.09 | 84.22 | 0.30 | 85.68 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.00 | 67.25 | 0.10 | 67.34 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 69.41 | 0.20 | 69.84 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 66.95 | 0.20 | 66.99 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 69.13 | 0.10 | 70.16 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.42 | 0.20 | 68.47 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 68.38 | 0.20 | 68.40 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 68.64 | 0.10 | 69.28 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.97 | 114.67 | 1.70 | 120.19 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 67.76 | 0.20 | 68.20 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 70.86 | 0.20 | 71.56 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 68.69 | 0.10 | 68.71 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 65.37 | 0.20 | 66.14 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.77 | 0.20 | 67.96 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 68.20 | 0.20 | 68.52 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 69.81 | 0.20 | 70.04 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 5.83 | 243.23 | 9.50 | 267.83 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 66.68 | 0.10 | 67.31 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.24 | 83.37 | 4.50 | 84.75 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.99 | 87.84 | 15.60 | 90.53 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.69 | 87.44 | 4.90 | 89.92 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.85 | 84.36 | 4.20 | 85.48 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.75 | 143.01 | 4.90 | 190.66 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 4.03 | 82.49 | 4.50 | 83.61 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.36 | 83.38 | 3.70 | 84.15 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.55 | 96.48 | 3.90 | 96.96 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 4.61 | 85.43 | 17.10 | 87.76 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.72 | 87.34 | 20.60 | 89.79 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 24.95 | 165.70 | 44.50 | 200.93 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 29.48 | 91.45 | 30.70 | 94.17 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 27.89 | 83.20 | 28.90 | 85.38 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 28.97 | 732.76 | 41.50 | 1247.58 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 25.14 | 84.50 | 26.90 | 86.21 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 28.70 | 85.97 | 29.90 | 88.21 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 24.43 | 98.50 | 25.90 | 98.92 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 32.98 | 289.99 | 47.40 | 408.40 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 26.43 | 98.19 | 28.30 | 104.59 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 23.98 | 177.27 | 38.00 | 210.46 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 108.02 | 95.17 | 111.90 | 98.86 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 99.64 | 89.69 | 105.39 | 92.02 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 93.89 | 3268.96 | 158.20 | 5727.72 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 82.81 | 85.77 | 93.51 | 86.85 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 125.05 | 16898.36 | 446.27 | 29506.84 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 93.88 | 100.21 | 106.50 | 102.06 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 34.50 | 392.38 | 47.20 | 500.73 |
