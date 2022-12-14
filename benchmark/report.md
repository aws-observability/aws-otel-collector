## Performance Report

**Version:** [v0.24.1](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.24.1)

**Commit ID:** [9bdaa9fec226e72c27235d955fea93022f599cc0](https://github.com/aws-observability/aws-otel-collector/commit/9bdaa9fec226e72c27235d955fea93022f599cc0)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.73 | 0.20 | 70.34 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.27 | 0.20 | 66.58 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.76 | 0.20 | 66.54 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 66.11 | 0.20 | 66.83 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 65.47 | 0.10 | 66.28 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.30 | 0.20 | 66.71 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.09 | 80.41 | 0.30 | 81.68 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 64.71 | 0.10 | 65.59 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.63 | 0.20 | 69.17 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 66.14 | 0.20 | 66.27 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.26 | 0.20 | 67.15 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.98 | 0.20 | 64.52 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 66.27 | 0.20 | 66.67 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.22 | 0.20 | 67.53 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.96 | 108.52 | 1.70 | 114.91 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 65.72 | 0.10 | 66.40 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.91 | 0.20 | 67.00 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.91 | 0.20 | 65.43 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.00 | 0.20 | 66.61 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.53 | 0.20 | 65.99 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 66.33 | 0.20 | 66.63 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 67.66 | 0.20 | 67.78 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 6.07 | 233.85 | 9.40 | 255.57 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 65.09 | 0.10 | 65.41 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.04 | 79.04 | 4.30 | 80.80 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.05 | 85.87 | 15.70 | 88.96 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.19 | 84.79 | 4.50 | 87.38 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 4.23 | 80.55 | 4.70 | 82.48 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.31 | 143.95 | 4.50 | 200.13 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.63 | 80.49 | 4.10 | 82.05 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.59 | 79.67 | 3.90 | 80.85 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.07 | 90.91 | 3.30 | 93.37 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.42 | 84.57 | 17.60 | 91.07 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.97 | 82.79 | 21.00 | 85.40 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 23.96 | 153.54 | 39.80 | 187.74 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 27.63 | 87.98 | 28.10 | 90.02 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 30.03 | 81.90 | 31.90 | 84.42 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 29.06 | 711.33 | 40.50 | 1178.60 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 26.51 | 79.54 | 29.20 | 81.39 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 28.95 | 82.72 | 29.50 | 84.92 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 25.07 | 93.26 | 26.80 | 94.58 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 35.02 | 325.66 | 49.70 | 479.79 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 27.24 | 96.13 | 28.40 | 102.66 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 23.52 | 180.34 | 36.60 | 206.94 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 114.82 | 90.74 | 116.49 | 96.29 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 116.57 | 84.66 | 118.60 | 88.77 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 113.49 | 3507.02 | 175.49 | 6040.63 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 104.51 | 83.10 | 111.70 | 84.79 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 115.69 | 14381.34 | 335.56 | 30415.26 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 98.31 | 97.86 | 103.60 | 99.38 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 33.91 | 428.33 | 55.30 | 573.10 |
