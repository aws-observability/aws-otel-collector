## Performance Report

**Version:** [v0.28.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.28.0)

**Commit ID:** [57d0cc35545a03f8e19b8a5d0247064836019d0e](https://github.com/aws-observability/aws-otel-collector/commit/57d0cc35545a03f8e19b8a5d0247064836019d0e)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 70.87 | 0.20 | 71.76 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.81 | 0.20 | 69.97 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 74.41 | 0.20 | 76.19 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 75.51 | 0.20 | 77.00 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 71.46 | 0.20 | 71.69 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 70.40 | 0.10 | 71.05 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 69.20 | 0.20 | 69.45 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 71.34 | 0.10 | 71.69 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.10 | 86.27 | 0.30 | 86.95 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 69.72 | 0.10 | 69.87 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 72.93 | 0.20 | 73.01 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 70.47 | 0.20 | 70.48 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 73.87 | 0.20 | 74.74 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.15 | 74.30 | 0.30 | 74.59 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 70.14 | 0.10 | 70.28 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 70.83 | 0.20 | 72.01 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 70.87 | 0.20 | 71.20 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 70.70 | 0.10 | 70.86 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.06 | 113.18 | 1.80 | 118.33 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 70.44 | 0.10 | 70.56 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 71.37 | 0.20 | 71.51 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 68.87 | 0.20 | 68.98 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.17 | 74.84 | 0.30 | 75.70 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 76.08 | 0.20 | 77.52 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 71.09 | 0.10 | 72.04 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 69.00 | 0.10 | 69.18 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 72.73 | 0.20 | 73.26 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 71.05 | 0.10 | 71.05 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 5.22 | 251.74 | 9.40 | 270.20 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 70.30 | 0.10 | 70.84 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.08 | 84.11 | 4.30 | 85.25 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.90 | 91.22 | 15.70 | 93.85 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.26 | 90.24 | 4.50 | 91.76 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 30.44 | 140.67 | 41.90 | 178.34 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 5.53 | 88.80 | 5.70 | 89.21 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 4.75 | 85.60 | 4.90 | 87.05 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.19 | 131.93 | 4.20 | 184.77 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.63 | 85.45 | 4.00 | 86.82 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.40 | 86.04 | 3.50 | 86.98 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.54 | 98.11 | 3.70 | 98.33 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 4.54 | 88.97 | 17.00 | 91.57 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.55 | 88.46 | 19.90 | 90.74 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 26.61 | 173.64 | 43.10 | 203.58 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 30.15 | 95.76 | 31.90 | 98.12 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 49.92 | 89.68 | 57.60 | 90.65 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 76.22 | 152.92 | 89.60 | 227.42 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 32.63 | 85.67 | 33.30 | 87.96 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 24.70 | 748.32 | 34.70 | 1185.59 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 25.83 | 85.25 | 26.10 | 87.61 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 28.22 | 86.80 | 30.00 | 88.80 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 25.11 | 99.78 | 26.10 | 100.35 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 33.00 | 283.19 | 48.50 | 387.42 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.91 | 98.64 | 26.90 | 104.35 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 23.92 | 183.27 | 43.00 | 206.84 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 106.88 | 95.95 | 115.90 | 100.99 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 153.29 | 1968.54 | 194.71 | 3834.56 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 143.46 | 2282.48 | 204.09 | 5793.12 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 102.69 | 91.23 | 107.90 | 93.60 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 97.52 | 3411.54 | 154.29 | 5575.87 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 94.20 | 87.44 | 102.30 | 89.20 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 110.23 | 15899.38 | 291.50 | 25595.11 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 90.59 | 101.31 | 100.60 | 102.69 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 34.30 | 373.18 | 52.40 | 462.61 |
