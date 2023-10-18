## Performance Report

**Version:** [v0.34.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.34.0)

**Commit ID:** [112425152e287bd99ba7240ff151d8537b76deaf](https://github.com/aws-observability/aws-otel-collector/commit/112425152e287bd99ba7240ff151d8537b76deaf)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 82.39 | 0.20 | 83.15 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 79.47 | 0.30 | 80.45 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 85.61 | 0.30 | 87.31 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.16 | 85.61 | 0.40 | 90.33 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.39 | 0.20 | 81.53 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 82.14 | 0.10 | 83.34 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 80.75 | 0.20 | 81.92 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 83.47 | 0.20 | 83.91 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.07 | 105.42 | 0.30 | 107.47 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 79.38 | 0.10 | 80.17 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 83.20 | 0.20 | 84.58 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.07 | 0.20 | 81.66 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 87.45 | 0.30 | 87.73 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 85.72 | 0.20 | 90.77 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.12 | 0.20 | 80.96 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 79.52 | 0.20 | 80.63 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 80.73 | 0.20 | 81.67 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.46 | 0.20 | 82.31 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.67 | 131.28 | 1.10 | 141.66 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 80.63 | 0.10 | 81.81 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 83.28 | 0.10 | 84.58 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 79.26 | 0.20 | 80.13 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.17 | 86.69 | 0.40 | 90.26 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.07 | 85.46 | 0.20 | 89.34 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.66 | 0.20 | 82.54 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.72 | 0.20 | 81.83 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 83.70 | 0.20 | 84.80 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.63 | 0.20 | 82.69 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.25 | 252.99 | 7.30 | 288.44 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 81.16 | 0.20 | 81.94 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.42 | 105.66 | 3.80 | 106.10 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.69 | 116.82 | 15.00 | 119.77 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.11 | 114.04 | 4.40 | 115.75 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 5.43 | 115.76 | 5.70 | 115.95 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 6.23 | 110.55 | 7.20 | 110.67 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.50 | 113.91 | 3.70 | 115.54 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 5.94 | 124.37 | 6.50 | 142.11 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 2.99 | 140.30 | 3.50 | 144.74 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.94 | 112.65 | 4.30 | 114.32 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.38 | 106.52 | 4.30 | 106.61 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 2.87 | 116.88 | 3.10 | 117.22 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 3.77 | 115.26 | 16.30 | 118.83 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 17.48 | 109.45 | 18.20 | 111.50 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 23.15 | 193.60 | 39.40 | 220.57 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 28.48 | 118.88 | 28.80 | 125.76 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 72.60 | 173.53 | 83.40 | 203.43 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 50.43 | 117.00 | 82.81 | 175.38 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 27.59 | 113.10 | 28.80 | 115.95 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 45.71 | 163.92 | 47.00 | 164.54 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 25.60 | 513.18 | 29.70 | 533.82 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 25.18 | 114.60 | 25.70 | 116.54 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 26.35 | 109.21 | 26.60 | 110.72 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 24.89 | 120.46 | 25.70 | 120.46 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 32.72 | 295.48 | 45.90 | 437.28 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.29 | 119.87 | 26.70 | 127.66 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 23.24 | 214.34 | 36.70 | 245.46 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 104.69 | 123.83 | 108.70 | 128.50 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 163.36 | 10534.43 | 358.89 | 18684.87 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 158.52 | 12729.12 | 381.69 | 20965.45 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 95.37 | 111.10 | 101.70 | 112.92 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 174.37 | 215.08 | 186.40 | 219.27 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 102.89 | 1995.10 | 119.70 | 2248.04 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 85.94 | 110.00 | 90.90 | 111.93 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 126.14 | 16011.15 | 326.51 | 27706.48 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 96.08 | 122.43 | 103.60 | 123.70 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 32.50 | 404.82 | 50.40 | 496.66 |
