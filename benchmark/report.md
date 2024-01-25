## Performance Report

**Version:** [v0.37.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.37.0)

**Commit ID:** [4043ec75b4079b6c984784a7e0eee9524ea52757](https://github.com/aws-observability/aws-otel-collector/commit/4043ec75b4079b6c984784a7e0eee9524ea52757)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.51 | 97.16 | 0.70 | 98.94 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.20 | 98.86 | 0.60 | 99.59 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 10.80 | 115.57 | 11.60 | 120.30 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.18 | 90.51 | 0.40 | 91.30 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.18 | 98.74 | 0.30 | 100.51 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.40 | 94.28 | 0.50 | 94.85 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.22 | 100.17 | 0.40 | 101.97 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.08 | 93.37 | 0.20 | 95.20 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 77.50 | 0.10 | 79.28 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 2.33 | 101.57 | 2.50 | 102.97 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 2.65 | 102.65 | 2.80 | 103.66 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.49 | 97.71 | 0.60 | 98.44 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.43 | 93.82 | 0.60 | 94.43 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.56 | 101.26 | 0.80 | 103.39 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 1.79 | 97.30 | 1.90 | 98.09 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.86 | 100.69 | 1.10 | 103.45 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.74 | 118.11 | 1.30 | 124.15 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 79.26 | 0.10 | 79.33 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 10.75 | 118.84 | 11.70 | 124.59 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 10.56 | 116.29 | 10.90 | 118.82 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 1.80 | 101.39 | 2.10 | 104.93 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 1.46 | 97.93 | 1.70 | 100.10 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 2.01 | 102.35 | 2.20 | 105.14 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 8.50 | 108.19 | 9.30 | 110.31 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 3.80 | 100.85 | 4.30 | 105.78 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.54 | 240.76 | 7.50 | 261.11 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 77.95 | 0.10 | 78.97 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.95 | 93.30 | 4.30 | 94.23 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 80.65 | 0.20 | 81.81 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 82.97 | 0.20 | 83.08 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.18 | 84.76 | 0.30 | 86.47 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 83.49 | 0.20 | 84.22 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 78.57 | 0.20 | 78.62 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.02 | 78.67 | 0.20 | 78.79 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 79.60 | 0.10 | 80.67 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 78.15 | 0.10 | 78.63 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 78.93 | 0.20 | 79.76 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 78.42 | 0.20 | 78.57 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 77.32 | 0.20 | 77.64 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.90 | 96.40 | 20.80 | 98.50 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 78.10 | 0.10 | 78.19 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 79.86 | 0.20 | 79.96 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 82.13 | 0.20 | 84.42 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.17 | 84.37 | 0.30 | 86.59 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 78.40 | 0.20 | 79.42 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 79.91 | 0.20 | 79.92 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 79.65 | 0.20 | 80.36 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 77.28 | 0.20 | 77.61 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.03 | 79.18 | 0.20 | 79.44 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 77.86 | 0.10 | 78.22 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.03 | 79.17 | 0.10 | 79.25 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 27.08 | 108.16 | 29.10 | 114.46 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 80.00 | 0.20 | 81.56 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 82.61 | 0.20 | 82.96 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 84.13 | 0.20 | 86.17 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 84.91 | 0.20 | 87.26 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 77.45 | 0.20 | 78.29 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 79.92 | 0.20 | 79.97 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 80.72 | 0.10 | 81.66 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 78.57 | 0.10 | 79.60 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 78.62 | 0.20 | 79.35 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 78.52 | 0.20 | 80.07 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 80.56 | 0.20 | 81.53 |
