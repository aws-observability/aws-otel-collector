## Performance Report

**Version:** [v0.29.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.29.0)

**Commit ID:** [a267676e2ed3d5163403941a9605e9a3e72b9da0](https://github.com/aws-observability/aws-otel-collector/commit/a267676e2ed3d5163403941a9605e9a3e72b9da0)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 72.22 | 0.20 | 72.38 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 72.46 | 0.20 | 73.24 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 74.48 | 0.20 | 75.21 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 76.80 | 0.20 | 77.84 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 71.60 | 0.10 | 72.19 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 71.11 | 0.10 | 71.51 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 72.28 | 0.20 | 72.66 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 72.27 | 0.20 | 73.35 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.09 | 86.48 | 0.30 | 88.36 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 71.55 | 0.10 | 71.64 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 72.15 | 0.20 | 72.38 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 72.23 | 0.20 | 72.71 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 76.08 | 0.20 | 77.69 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 75.92 | 0.20 | 76.79 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 72.68 | 0.20 | 73.27 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 71.51 | 0.10 | 71.81 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 72.24 | 0.10 | 73.28 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 72.03 | 0.20 | 72.39 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.03 | 113.36 | 1.90 | 118.82 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 72.09 | 0.20 | 72.17 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 74.20 | 0.20 | 74.80 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 70.80 | 0.10 | 71.43 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.16 | 77.28 | 0.30 | 77.82 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.16 | 78.97 | 0.30 | 79.47 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 72.53 | 0.10 | 72.93 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 71.72 | 0.20 | 71.91 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 71.36 | 0.10 | 72.06 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 72.83 | 0.10 | 73.89 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 5.94 | 250.34 | 9.50 | 268.20 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 70.12 | 0.20 | 70.15 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.72 | 86.20 | 4.00 | 87.17 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.07 | 91.66 | 10.10 | 93.70 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.75 | 90.14 | 5.50 | 92.17 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 7.02 | 90.44 | 7.60 | 90.96 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 5.97 | 92.69 | 8.90 | 93.25 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.69 | 87.98 | 3.90 | 89.16 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 4.96 | 107.03 | 5.60 | 125.10 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 2.93 | 144.83 | 4.10 | 189.81 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.92 | 87.26 | 4.30 | 89.05 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.45 | 86.56 | 3.60 | 87.35 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.03 | 98.72 | 3.20 | 99.21 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 4.97 | 91.50 | 17.60 | 94.07 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 18.64 | 90.26 | 19.30 | 92.20 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 26.11 | 158.36 | 43.10 | 194.82 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 29.59 | 95.29 | 35.60 | 98.66 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 74.70 | 154.35 | 88.20 | 224.06 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 49.11 | 91.84 | 59.20 | 92.60 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 25.61 | 88.11 | 26.30 | 89.60 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 47.05 | 142.16 | 51.20 | 146.62 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 26.39 | 738.67 | 38.60 | 1329.49 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 25.36 | 87.22 | 28.30 | 89.55 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 27.67 | 88.93 | 29.20 | 91.39 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 24.57 | 100.75 | 25.60 | 101.56 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 33.88 | 285.69 | 50.90 | 436.72 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 26.19 | 103.23 | 27.60 | 109.65 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 25.10 | 184.26 | 41.80 | 210.41 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 115.88 | 100.85 | 118.80 | 103.92 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 143.60 | 2226.43 | 192.71 | 4613.91 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 134.90 | 3433.02 | 210.91 | 6735.71 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 100.90 | 94.81 | 109.90 | 97.09 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 184.18 | 190.24 | 196.10 | 193.93 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 93.05 | 3376.97 | 158.10 | 6049.51 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 91.42 | 89.68 | 96.90 | 90.95 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 111.64 | 15215.11 | 315.20 | 26748.25 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 85.13 | 103.90 | 89.49 | 106.33 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 33.93 | 384.09 | 50.00 | 519.72 |
