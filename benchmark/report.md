## Performance Report

**Version:** [v0.31.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.31.0)

**Commit ID:** [375889537354a80d2393e697e3f510889805204d](https://github.com/aws-observability/aws-otel-collector/commit/375889537354a80d2393e697e3f510889805204d)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 76.04 | 0.10 | 76.18 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 74.12 | 0.10 | 74.34 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.04 | 78.21 | 0.10 | 78.28 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 78.76 | 0.20 | 79.23 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 73.22 | 0.20 | 73.28 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 74.71 | 0.20 | 74.93 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 74.02 | 0.10 | 74.07 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 73.97 | 0.10 | 74.19 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.09 | 91.14 | 0.30 | 92.25 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 73.53 | 0.10 | 73.83 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 75.07 | 0.20 | 75.46 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 73.48 | 0.20 | 74.01 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.16 | 79.15 | 0.30 | 79.89 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.07 | 77.12 | 0.20 | 78.09 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 74.11 | 0.20 | 74.22 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 73.80 | 0.20 | 74.49 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 74.82 | 0.20 | 75.28 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 73.25 | 0.20 | 73.47 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.98 | 118.98 | 1.60 | 124.22 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 73.77 | 0.10 | 74.47 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 76.42 | 0.20 | 76.59 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 73.15 | 0.20 | 73.80 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 77.92 | 0.20 | 78.55 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.17 | 78.19 | 0.30 | 79.22 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 75.05 | 0.10 | 75.85 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 71.66 | 0.20 | 72.07 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 72.04 | 0.20 | 72.39 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 74.88 | 0.20 | 75.22 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 6.19 | 251.70 | 9.90 | 272.42 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 73.38 | 0.10 | 73.70 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.68 | 89.19 | 3.90 | 90.52 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.77 | 97.56 | 15.50 | 100.18 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.66 | 93.29 | 5.10 | 94.52 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 6.42 | 94.22 | 7.90 | 94.60 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 32.77 | 137.83 | 45.00 | 175.05 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.49 | 97.05 | 3.70 | 99.45 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 5.11 | 109.55 | 5.90 | 127.77 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 4.05 | 124.42 | 4.80 | 129.43 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.43 | 97.73 | 4.00 | 100.48 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.65 | 89.37 | 4.40 | 90.01 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.60 | 100.93 | 4.10 | 101.28 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 4.33 | 95.44 | 15.70 | 99.46 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 18.41 | 92.87 | 19.10 | 94.95 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 26.18 | 162.04 | 44.70 | 194.42 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 32.13 | 98.96 | 34.50 | 100.63 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 73.95 | 154.96 | 88.49 | 195.01 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 50.57 | 99.00 | 82.90 | 161.85 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 27.05 | 93.58 | 27.60 | 95.44 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 51.83 | 146.60 | 54.40 | 150.94 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 27.01 | 466.59 | 31.00 | 518.31 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 24.85 | 93.93 | 25.30 | 96.09 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 28.70 | 91.35 | 30.50 | 94.83 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 25.45 | 103.74 | 26.00 | 104.41 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 34.64 | 287.10 | 48.70 | 417.39 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 26.12 | 105.41 | 27.60 | 112.21 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 25.93 | 186.08 | 41.50 | 221.67 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 111.66 | 104.22 | 121.40 | 108.21 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 189.86 | 12990.09 | 350.21 | 22199.98 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 155.06 | 10854.39 | 321.88 | 19645.31 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 102.19 | 93.82 | 110.00 | 96.37 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 189.79 | 192.37 | 200.90 | 196.28 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 103.22 | 1980.52 | 119.60 | 2234.14 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 96.09 | 92.91 | 106.00 | 95.16 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 106.00 | 14739.75 | 328.51 | 27452.81 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 83.56 | 105.47 | 89.10 | 107.62 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 34.38 | 394.95 | 60.00 | 500.31 |
