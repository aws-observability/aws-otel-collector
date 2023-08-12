## Performance Report

**Version:** [v0.31.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.31.0)

**Commit ID:** [f583d3e1f994ab4ad5ae9f15f5ae2a594c44323d](https://github.com/aws-observability/aws-otel-collector/commit/f583d3e1f994ab4ad5ae9f15f5ae2a594c44323d)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 77.19 | 0.20 | 77.38 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 75.25 | 0.20 | 75.76 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 79.73 | 0.20 | 81.31 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 80.24 | 0.20 | 81.02 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 76.78 | 0.20 | 77.11 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 74.42 | 0.20 | 74.78 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 75.49 | 0.10 | 76.13 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 76.60 | 0.20 | 76.78 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.09 | 91.26 | 0.20 | 92.28 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 74.84 | 0.10 | 75.33 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 78.33 | 0.10 | 78.84 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 75.62 | 0.20 | 76.12 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.16 | 81.31 | 0.30 | 82.26 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.14 | 82.56 | 0.30 | 83.76 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 74.83 | 0.20 | 75.31 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 75.17 | 0.20 | 75.53 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 75.45 | 0.10 | 76.12 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 76.10 | 0.20 | 76.64 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.88 | 118.35 | 1.50 | 122.13 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 75.33 | 0.10 | 75.74 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 76.34 | 0.20 | 76.78 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 75.78 | 0.20 | 76.44 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 79.99 | 0.20 | 80.49 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 81.15 | 0.20 | 81.82 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 73.95 | 0.10 | 74.14 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 74.23 | 0.20 | 74.38 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 75.05 | 0.20 | 75.21 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 75.57 | 0.20 | 75.74 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 5.58 | 234.88 | 8.70 | 280.12 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 72.42 | 0.10 | 73.02 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.80 | 89.82 | 4.10 | 90.59 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.94 | 98.20 | 15.70 | 100.11 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 5.48 | 94.45 | 6.20 | 96.85 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 33.09 | 148.43 | 42.80 | 187.76 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 5.80 | 94.70 | 8.10 | 95.25 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.49 | 98.35 | 3.70 | 101.37 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 5.30 | 112.81 | 6.30 | 127.97 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.21 | 124.04 | 3.60 | 130.38 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.71 | 99.56 | 4.20 | 102.29 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.51 | 91.66 | 5.00 | 92.60 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 4.10 | 102.89 | 4.60 | 103.19 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 4.44 | 98.01 | 16.40 | 103.41 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.24 | 93.82 | 21.90 | 96.38 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 25.69 | 172.89 | 42.50 | 209.81 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 30.20 | 102.97 | 31.60 | 104.91 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 46.93 | 94.96 | 54.50 | 95.99 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 75.46 | 146.32 | 86.89 | 175.53 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 28.58 | 95.06 | 29.70 | 97.40 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 48.44 | 146.47 | 50.10 | 150.77 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 27.36 | 470.73 | 29.90 | 519.19 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 25.87 | 94.57 | 26.40 | 96.58 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 28.16 | 92.78 | 28.70 | 94.58 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 26.39 | 105.26 | 27.00 | 106.10 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 34.65 | 287.60 | 50.60 | 391.76 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 28.36 | 105.51 | 30.60 | 111.75 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 25.20 | 202.23 | 43.30 | 240.15 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 120.67 | 111.91 | 124.80 | 117.42 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 160.53 | 11517.48 | 336.27 | 21207.48 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 165.54 | 13314.46 | 388.79 | 22561.93 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 106.17 | 96.95 | 110.01 | 98.95 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 187.94 | 194.09 | 195.20 | 200.34 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 105.38 | 1963.26 | 121.60 | 2233.09 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 95.97 | 94.48 | 101.70 | 96.78 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 113.55 | 16464.80 | 366.01 | 30815.88 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 89.32 | 108.37 | 95.70 | 110.20 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 33.28 | 394.42 | 51.10 | 495.60 |
