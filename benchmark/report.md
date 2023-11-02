## Performance Report

**Version:** [v0.35.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.35.0)

**Commit ID:** [f7f5e40f7c2a49e75bbe8fd3096620848d6eb5af](https://github.com/aws-observability/aws-otel-collector/commit/f7f5e40f7c2a49e75bbe8fd3096620848d6eb5af)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 81.42 | 0.20 | 82.60 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 79.58 | 0.20 | 80.28 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.16 | 85.69 | 0.30 | 86.34 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 85.42 | 0.30 | 85.91 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 80.32 | 0.20 | 81.59 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 82.98 | 0.20 | 83.87 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 80.48 | 0.20 | 81.33 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 79.68 | 0.20 | 80.68 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.07 | 104.53 | 0.20 | 107.32 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 79.67 | 0.10 | 80.63 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 82.13 | 0.20 | 83.51 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 79.92 | 0.20 | 80.72 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 86.26 | 0.20 | 86.86 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 83.13 | 0.20 | 83.73 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.06 | 0.10 | 82.13 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 79.89 | 0.20 | 80.78 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 81.31 | 0.20 | 81.89 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 79.43 | 0.20 | 80.14 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.66 | 129.47 | 1.40 | 140.05 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 79.72 | 0.10 | 80.38 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 84.06 | 0.20 | 85.18 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.97 | 0.20 | 81.89 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 85.03 | 0.20 | 85.33 |
| otlp | batch | kafka/exporter, logging | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.18 | 85.31 | 0.40 | 86.45 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.75 | 0.20 | 81.95 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 79.76 | 0.20 | 80.36 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 80.48 | 0.10 | 81.44 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.15 | 0.20 | 82.19 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.29 | 250.80 | 7.40 | 280.78 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 78.99 | 0.10 | 79.84 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.55 | 106.14 | 4.00 | 107.10 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.76 | 118.67 | 15.50 | 121.46 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.69 | 113.63 | 5.10 | 115.41 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 30.36 | 164.88 | 39.70 | 219.77 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 32.40 | 172.74 | 42.70 | 200.79 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.41 | 115.22 | 3.70 | 117.28 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 5.98 | 126.23 | 6.60 | 144.81 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.28 | 138.08 | 3.80 | 145.20 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 4.41 | 113.33 | 4.70 | 115.16 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.60 | 104.73 | 3.80 | 105.78 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.24 | 116.60 | 3.50 | 117.20 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 3.96 | 113.66 | 16.00 | 117.09 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 17.79 | 113.66 | 18.50 | 115.03 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 25.03 | 204.03 | 44.00 | 229.14 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 30.04 | 118.49 | 30.80 | 121.66 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 48.22 | 110.40 | 55.40 | 111.38 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 49.71 | 111.65 | 57.50 | 112.20 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 28.14 | 112.79 | 29.10 | 114.63 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 47.01 | 164.28 | 48.50 | 168.97 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 26.70 | 483.18 | 33.90 | 534.53 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 25.19 | 111.95 | 26.00 | 114.34 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 28.14 | 113.06 | 29.30 | 113.77 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 24.76 | 118.96 | 25.30 | 119.43 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 32.19 | 311.27 | 45.80 | 405.29 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.24 | 118.47 | 26.90 | 124.91 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 23.27 | 214.43 | 40.00 | 237.34 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 112.72 | 121.36 | 121.19 | 126.20 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 190.70 | 7921.92 | 331.58 | 14399.65 |
| otlp |  | kafka/exporter, logging | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 186.09 | 7405.30 | 315.79 | 13660.23 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 101.32 | 112.66 | 110.60 | 115.39 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 188.13 | 211.77 | 196.10 | 216.06 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 102.00 | 2002.93 | 122.90 | 2252.57 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 92.02 | 111.86 | 96.10 | 113.68 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 117.64 | 18471.69 | 415.11 | 30986.26 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 97.73 | 124.46 | 106.10 | 125.85 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 32.44 | 399.46 | 53.60 | 476.39 |
