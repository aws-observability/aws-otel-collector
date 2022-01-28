## Performance Report

**Version:** [v0.16.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.16.0)

**Commit ID:** [cfe4461de4990be4e0291af87340887f126ab1f4](https://github.com/aws-observability/aws-otel-collector/commit/cfe4461de4990be4e0291af87340887f126ab1f4)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.98 | 0.30 | 66.37 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.09 | 0.20 | 63.24 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.18 | 0.20 | 62.69 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.48 | 0.20 | 62.70 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 63.26 | 0.20 | 63.96 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.31 | 0.20 | 63.88 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.12 | 75.62 | 0.30 | 77.43 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 61.20 | 0.20 | 61.62 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.92 | 0.30 | 67.75 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.14 | 0.20 | 62.82 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.52 | 0.20 | 62.65 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.83 | 0.20 | 63.76 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 62.63 | 0.20 | 63.34 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.61 | 0.20 | 63.09 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.29 | 113.40 | 2.10 | 117.06 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 61.56 | 0.20 | 61.93 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.87 | 0.40 | 65.88 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.25 | 0.20 | 63.01 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.38 | 0.20 | 64.06 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.50 | 0.20 | 62.54 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 61.28 | 0.20 | 61.40 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.88 | 0.20 | 64.01 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.41 | 267.78 | 14.80 | 299.59 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 61.54 | 0.10 | 61.89 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.67 | 145.49 | 5.60 | 212.93 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.33 | 79.18 | 2.80 | 81.23 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.99 | 78.61 | 5.60 | 79.35 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.10 | 94.94 | 3.40 | 95.10 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.03 | 133.82 | 4.30 | 182.80 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.44 | 73.69 | 3.90 | 74.24 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.41 | 75.28 | 4.70 | 76.17 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.27 | 87.50 | 3.50 | 87.78 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.88 | 81.32 | 6.80 | 84.84 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 26.58 | 506.75 | 35.00 | 878.63 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.43 | 153.28 | 21.10 | 183.19 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 31.39 | 79.55 | 32.00 | 80.86 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 27.42 | 105.44 | 28.40 | 110.39 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 28.64 | 683.41 | 40.31 | 1158.60 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 26.42 | 75.14 | 27.10 | 76.01 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 34.33 | 79.65 | 44.52 | 80.88 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 27.69 | 88.46 | 28.31 | 88.90 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 29.86 | 453.52 | 36.58 | 537.00 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 38.10 | 777.23 | 47.91 | 1233.84 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 17.25 | 170.69 | 22.01 | 199.14 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 115.05 | 86.37 | 117.08 | 87.31 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 120.50 | 126.09 | 121.61 | 132.14 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 132.89 | 3218.09 | 196.32 | 5917.88 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 108.33 | 80.28 | 112.52 | 81.06 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 148.61 | 14980.12 | 448.57 | 25942.45 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 120.08 | 95.66 | 121.31 | 97.54 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 28.49 | 511.49 | 36.00 | 551.90 |
