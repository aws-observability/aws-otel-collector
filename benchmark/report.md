## Performance Report

**Version:** [v0.32.2](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.32.2)

**Commit ID:** [a1f7d451718ab8f5a86f7785c327be4db18a9854](https://github.com/aws-observability/aws-otel-collector/commit/a1f7d451718ab8f5a86f7785c327be4db18a9854)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 78.87 | 0.20 | 79.06 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 75.33 | 0.10 | 76.16 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.07 | 80.95 | 0.20 | 81.52 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.16 | 80.96 | 0.30 | 81.64 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 76.19 | 0.10 | 77.08 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 74.45 | 0.10 | 75.20 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 76.06 | 0.20 | 76.45 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 74.25 | 0.20 | 74.58 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.09 | 91.23 | 0.20 | 92.36 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 74.61 | 0.10 | 75.51 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 77.19 | 0.20 | 77.59 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 73.25 | 0.10 | 74.12 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.19 | 80.57 | 0.40 | 82.06 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 80.32 | 0.20 | 80.97 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 74.98 | 0.20 | 75.28 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 75.59 | 0.20 | 75.94 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 73.80 | 0.20 | 74.37 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 75.12 | 0.20 | 75.50 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.94 | 116.67 | 1.70 | 125.89 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 74.74 | 0.20 | 74.86 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 75.32 | 0.20 | 75.51 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 75.16 | 0.20 | 75.90 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 78.28 | 0.20 | 79.53 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 79.40 | 0.20 | 80.72 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 76.10 | 0.10 | 76.18 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 75.26 | 0.10 | 75.63 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 74.63 | 0.10 | 75.13 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 75.92 | 0.10 | 76.38 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 5.53 | 244.83 | 8.90 | 268.45 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 73.11 | 0.20 | 73.44 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.88 | 90.83 | 4.40 | 91.71 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.06 | 98.91 | 15.40 | 102.54 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.56 | 94.79 | 4.80 | 96.31 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 6.52 | 96.48 | 7.50 | 97.28 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 5.63 | 96.13 | 6.10 | 96.46 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.29 | 98.05 | 3.60 | 100.04 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 5.32 | 110.94 | 6.60 | 127.99 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.51 | 124.89 | 4.20 | 130.63 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.83 | 97.83 | 4.70 | 99.75 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.71 | 90.62 | 4.30 | 91.72 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.89 | 102.67 | 4.40 | 103.26 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 4.54 | 98.23 | 17.10 | 102.19 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.27 | 93.52 | 19.80 | 95.44 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 26.93 | 163.93 | 47.30 | 186.49 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 30.28 | 104.12 | 30.80 | 111.35 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 76.18 | 146.82 | 88.30 | 190.87 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 77.54 | 147.90 | 90.10 | 232.92 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 29.35 | 94.86 | 35.20 | 96.92 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 48.25 | 146.27 | 49.80 | 149.87 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 27.58 | 468.68 | 30.40 | 520.61 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 25.86 | 94.76 | 26.40 | 96.48 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 29.76 | 92.62 | 31.00 | 94.58 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 25.05 | 105.45 | 25.90 | 105.98 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 34.46 | 275.22 | 50.50 | 388.83 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 28.21 | 105.53 | 29.50 | 112.09 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 26.06 | 190.84 | 44.30 | 223.19 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 115.18 | 108.71 | 124.00 | 119.00 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 155.04 | 11782.84 | 342.07 | 19734.41 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 167.54 | 12585.78 | 409.89 | 22914.36 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 106.42 | 95.46 | 109.80 | 97.67 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 187.86 | 193.41 | 194.20 | 197.43 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 104.52 | 1986.38 | 118.20 | 2228.36 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 90.31 | 94.44 | 95.30 | 96.44 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 113.75 | 15910.30 | 284.01 | 25571.06 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 88.37 | 107.49 | 95.61 | 108.82 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 34.20 | 379.84 | 51.80 | 445.20 |
