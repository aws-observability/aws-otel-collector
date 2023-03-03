## Performance Report

**Version:** [v0.27.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.27.0)

**Commit ID:** [7b82fa16ca33cb516de744a5b1a25e6a5049d5fa](https://github.com/aws-observability/aws-otel-collector/commit/7b82fa16ca33cb516de744a5b1a25e6a5049d5fa)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 69.73 | 0.20 | 69.75 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.65 | 0.20 | 68.90 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 68.67 | 0.20 | 68.94 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.97 | 0.20 | 68.35 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 67.95 | 0.10 | 68.35 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 69.17 | 0.20 | 70.25 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.10 | 81.86 | 0.30 | 83.60 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 67.92 | 0.10 | 68.42 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 71.14 | 0.20 | 71.23 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 69.12 | 0.20 | 69.62 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 67.74 | 0.20 | 68.62 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.11 | 0.20 | 68.93 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 69.14 | 0.10 | 69.17 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 69.43 | 0.10 | 69.80 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.06 | 114.35 | 1.80 | 119.53 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 68.51 | 0.10 | 69.37 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 70.21 | 0.10 | 70.69 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 70.02 | 0.10 | 70.53 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.38 | 0.10 | 68.71 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.87 | 0.20 | 68.15 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 66.28 | 0.20 | 66.77 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.39 | 0.20 | 68.92 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 6.21 | 238.07 | 10.00 | 260.23 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 68.32 | 0.10 | 69.05 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.76 | 82.21 | 4.30 | 83.76 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.96 | 85.72 | 15.50 | 87.74 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.73 | 85.31 | 5.10 | 87.51 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 4.47 | 81.91 | 4.70 | 83.80 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.29 | 136.61 | 4.50 | 189.47 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.71 | 80.72 | 4.00 | 82.24 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.51 | 80.24 | 3.70 | 81.78 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.80 | 94.74 | 4.00 | 96.06 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 4.70 | 86.96 | 17.20 | 89.30 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.87 | 84.51 | 20.80 | 87.55 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 26.82 | 155.77 | 42.80 | 197.40 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 30.35 | 88.40 | 31.30 | 92.81 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 29.33 | 83.12 | 31.00 | 85.48 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 27.76 | 724.20 | 40.10 | 1246.72 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 27.16 | 82.49 | 28.00 | 84.18 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 27.69 | 84.06 | 28.20 | 86.63 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 27.48 | 97.73 | 27.90 | 99.42 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 33.56 | 286.11 | 48.10 | 396.94 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 26.28 | 96.87 | 27.40 | 102.33 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 25.23 | 177.20 | 44.00 | 209.60 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 111.70 | 95.68 | 117.10 | 99.37 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 104.27 | 88.27 | 109.39 | 91.49 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 90.78 | 3455.75 | 152.60 | 5772.58 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 93.59 | 84.77 | 104.10 | 87.04 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 110.62 | 17830.03 | 295.09 | 30070.70 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 88.65 | 100.30 | 93.89 | 102.76 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 32.60 | 383.22 | 52.70 | 486.78 |
