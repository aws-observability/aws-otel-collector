## Performance Report

**Version:** [v0.44.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.44.0)

**Commit ID:** [6887f273fa4244e16a05cc0d06620b3aeaac7293](https://github.com/aws-observability/aws-otel-collector/commit/6887f273fa4244e16a05cc0d06620b3aeaac7293)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.55 | 126.51 | 0.70 | 127.18 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.25 | 127.71 | 0.70 | 127.98 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.28 | 126.72 | 0.50 | 127.17 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.28 | 115.14 | 0.40 | 116.78 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.28 | 122.62 | 0.40 | 123.93 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.07 | 110.44 | 0.60 | 119.03 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.24 | 124.38 | 0.40 | 127.57 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.08 | 123.45 | 0.30 | 125.69 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 100.91 | 0.10 | 101.87 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 1.62 | 133.06 | 1.90 | 137.68 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 2.42 | 138.12 | 2.80 | 142.58 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 9.97 | 144.43 | 10.60 | 147.37 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.59 | 129.19 | 0.80 | 132.90 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.55 | 137.47 | 0.70 | 141.82 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.12 | 115.45 | 1.70 | 130.17 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.96 | 136.52 | 1.10 | 140.01 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.67 | 152.85 | 1.30 | 159.82 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 103.02 | 0.10 | 103.94 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 5.29 | 139.20 | 6.10 | 143.88 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 9.16 | 145.81 | 9.80 | 150.27 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 1.75 | 138.10 | 1.90 | 144.61 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 1.93 | 138.88 | 4.60 | 269.77 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.42 | 121.51 | 2.00 | 140.28 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 1.22 | 128.13 | 7.40 | 141.17 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 3.81 | 134.65 | 4.10 | 138.38 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.21 | 258.78 | 7.30 | 274.72 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 101.35 | 0.10 | 102.56 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 0.14 | 110.98 | 2.50 | 118.00 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 102.77 | 0.10 | 103.67 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 103.62 | 0.20 | 104.19 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.09 | 107.35 | 0.20 | 109.00 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 107.31 | 0.30 | 110.17 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 101.09 | 0.20 | 101.80 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 104.42 | 0.20 | 105.37 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 104.28 | 0.20 | 105.50 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 100.66 | 0.20 | 101.68 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 101.82 | 0.20 | 102.49 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 101.64 | 0.20 | 101.83 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.06 | 101.29 | 0.20 | 102.12 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 1.45 | 113.23 | 25.10 | 121.95 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 101.96 | 0.20 | 102.88 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 102.08 | 0.20 | 103.14 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.24 | 108.25 | 0.60 | 110.94 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.09 | 108.61 | 0.20 | 111.39 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 100.21 | 0.20 | 101.44 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 104.11 | 0.20 | 104.37 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 101.83 | 0.30 | 102.74 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 102.10 | 0.30 | 102.65 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.06 | 101.91 | 0.20 | 103.19 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 100.54 | 0.20 | 101.07 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.03 | 101.68 | 0.20 | 102.14 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.26 | 115.53 | 28.60 | 134.17 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.03 | 103.18 | 0.20 | 104.26 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.08 | 105.26 | 0.20 | 105.26 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.09 | 108.14 | 0.30 | 109.76 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.25 | 109.24 | 0.50 | 111.62 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 100.60 | 0.20 | 100.94 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 103.80 | 0.10 | 104.66 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 103.23 | 0.20 | 103.95 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 101.60 | 0.20 | 102.57 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.06 | 100.56 | 0.20 | 101.90 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 102.40 | 0.20 | 103.69 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 99.95 | 0.10 | 101.18 |
