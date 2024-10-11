## Performance Report

**Version:** [v0.41.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.41.0)

**Commit ID:** [17b196536510c3263699fced0fda9ce4092166d8](https://github.com/aws-observability/aws-otel-collector/commit/17b196536510c3263699fced0fda9ce4092166d8)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.43 | 121.36 | 1.10 | 122.47 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 3.87 | 128.22 | 10.40 | 142.16 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 9.52 | 139.65 | 9.90 | 144.51 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.18 | 103.76 | 0.30 | 105.32 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.18 | 112.40 | 0.50 | 114.39 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.37 | 114.94 | 0.60 | 117.26 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.22 | 119.85 | 0.50 | 121.61 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.08 | 116.49 | 0.30 | 119.51 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 92.00 | 0.20 | 93.85 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 1.94 | 132.77 | 2.10 | 134.19 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 2.02 | 133.70 | 10.80 | 145.38 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.46 | 132.45 | 0.70 | 133.45 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.47 | 126.39 | 0.70 | 129.01 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.58 | 134.88 | 0.80 | 140.98 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 1.57 | 131.00 | 1.80 | 133.82 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.76 | 138.86 | 1.00 | 142.51 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.75 | 144.05 | 1.20 | 152.97 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 92.16 | 0.20 | 94.04 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 9.52 | 149.83 | 10.00 | 155.74 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 3.80 | 140.02 | 10.00 | 150.18 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 1.58 | 138.33 | 1.80 | 141.91 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 1.50 | 133.33 | 1.80 | 136.81 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 1.87 | 140.69 | 2.20 | 142.92 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 7.10 | 138.52 | 7.40 | 140.62 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 3.40 | 137.58 | 3.70 | 141.95 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.66 | 256.04 | 7.90 | 287.52 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 91.33 | 0.20 | 92.80 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.93 | 117.89 | 4.20 | 118.58 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 91.44 | 0.20 | 93.01 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 94.63 | 0.30 | 94.99 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.18 | 98.19 | 0.40 | 102.35 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.12 | 99.91 | 0.40 | 102.99 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 91.17 | 0.20 | 93.66 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 91.00 | 0.20 | 94.25 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 90.68 | 0.20 | 92.32 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 92.86 | 0.20 | 95.63 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.05 | 91.32 | 0.20 | 94.09 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 91.95 | 0.20 | 94.40 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 93.00 | 0.20 | 95.96 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 18.18 | 121.75 | 19.00 | 123.49 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 91.91 | 0.20 | 94.38 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 94.19 | 0.20 | 97.16 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 96.68 | 0.30 | 98.41 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 96.85 | 0.30 | 101.11 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 90.25 | 0.20 | 93.49 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 93.85 | 0.20 | 96.51 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 91.87 | 0.30 | 93.31 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 91.67 | 0.30 | 94.14 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 90.58 | 0.20 | 93.18 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 92.09 | 0.20 | 94.81 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 92.64 | 0.20 | 95.23 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.21 | 134.75 | 26.20 | 137.57 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 92.47 | 0.20 | 94.54 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 93.83 | 0.20 | 96.97 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.09 | 96.72 | 0.30 | 100.42 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.19 | 98.60 | 0.40 | 100.64 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 89.92 | 0.20 | 92.87 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 92.08 | 0.20 | 92.19 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 94.30 | 0.20 | 96.70 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 91.55 | 0.20 | 94.47 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 91.04 | 0.20 | 93.72 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 89.53 | 0.20 | 91.20 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 91.72 | 0.20 | 94.36 |
