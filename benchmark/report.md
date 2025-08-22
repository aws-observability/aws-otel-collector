## Performance Report

**Version:** [v0.45.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.45.0)

**Commit ID:** [3355ddd5d6cc3117464f376f640a8aef9040cce6](https://github.com/aws-observability/aws-otel-collector/commit/3355ddd5d6cc3117464f376f640a8aef9040cce6)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 111.64 | 0.40 | 121.67 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.27 | 126.69 | 0.40 | 127.47 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.20 | 121.34 | 0.40 | 123.40 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.12 | 112.45 | 0.40 | 116.97 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.16 | 117.10 | 0.40 | 123.41 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.47 | 120.18 | 0.60 | 120.85 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.31 | 123.95 | 0.50 | 127.17 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.09 | 121.52 | 0.30 | 123.76 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 100.66 | 0.10 | 101.19 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 1.47 | 131.33 | 1.70 | 136.26 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 2.80 | 139.07 | 3.00 | 141.79 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.49 | 135.63 | 0.70 | 139.35 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.32 | 122.16 | 1.10 | 140.76 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.29 | 132.32 | 0.80 | 142.91 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 1.59 | 131.94 | 1.90 | 135.74 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.82 | 136.86 | 1.00 | 141.51 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.71 | 151.73 | 1.30 | 159.18 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 102.48 | 0.10 | 103.33 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.52 | 122.34 | 5.30 | 137.67 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 9.26 | 148.12 | 10.40 | 151.52 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 10.85 | 146.54 | 12.00 | 150.46 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.82 | 126.58 | 2.20 | 137.95 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 2.26 | 139.33 | 2.60 | 142.11 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 6.72 | 140.60 | 7.10 | 144.57 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 3.71 | 135.63 | 4.10 | 141.02 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.07 | 261.79 | 6.80 | 275.18 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 102.65 | 0.10 | 104.03 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.35 | 119.68 | 3.60 | 120.90 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.03 | 101.60 | 0.20 | 101.81 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.08 | 105.07 | 0.20 | 105.65 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.08 | 107.88 | 0.20 | 109.84 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.16 | 108.97 | 0.30 | 109.26 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 101.47 | 0.20 | 102.16 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 100.39 | 0.20 | 100.98 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 102.29 | 0.20 | 103.12 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 102.96 | 0.20 | 103.06 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 101.76 | 0.20 | 102.37 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 100.68 | 0.20 | 101.64 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 98.66 | 0.20 | 99.28 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 21.91 | 119.37 | 28.00 | 121.41 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.06 | 101.96 | 0.20 | 102.99 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 105.84 | 0.20 | 106.15 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 107.43 | 0.20 | 109.51 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.09 | 105.86 | 0.20 | 108.96 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 99.02 | 0.20 | 100.01 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.04 | 102.15 | 0.20 | 103.16 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 101.09 | 0.10 | 101.88 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 100.15 | 0.20 | 101.23 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 100.88 | 0.10 | 101.73 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 102.24 | 0.20 | 103.40 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.03 | 98.93 | 0.10 | 99.12 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 96.38 | 130.75 | 110.60 | 135.74 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 100.86 | 0.20 | 101.86 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 102.55 | 0.20 | 103.64 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.24 | 109.13 | 0.40 | 112.14 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.14 | 109.43 | 0.40 | 111.19 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 102.92 | 0.20 | 102.98 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 102.46 | 0.20 | 102.48 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 102.60 | 0.20 | 103.13 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 100.12 | 0.20 | 100.79 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 101.11 | 0.20 | 102.22 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 101.63 | 0.10 | 102.25 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 101.45 | 0.20 | 102.26 |
