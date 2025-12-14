## Performance Report

**Version:** [v0.46.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.46.0)

**Commit ID:** [d00c5d966c0c9f9fcb0d2b465051d0cb7339a2a7](https://github.com/aws-observability/aws-otel-collector/commit/d00c5d966c0c9f9fcb0d2b465051d0cb7339a2a7)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.19 | 125.11 | 0.90 | 134.32 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 108.53 | 0.40 | 110.58 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.10 | 118.62 | 0.50 | 119.54 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.44 | 124.51 | 0.60 | 125.93 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.23 | 136.57 | 0.70 | 153.67 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.09 | 127.30 | 1.70 | 128.52 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 102.73 | 0.10 | 103.52 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.28 | 123.05 | 1.70 | 137.69 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 112.74 | 0.50 | 131.10 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.20 | 147.79 | 1.00 | 151.57 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 1.59 | 130.09 | 1.90 | 131.53 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.80 | 281.27 | 4.30 | 399.15 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.97 | 156.22 | 3.10 | 162.03 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 103.27 | 0.20 | 103.82 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 5.31 | 164.30 | 6.00 | 185.39 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.00 | 71.11 | 0.00 | 71.11 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.43 | 134.25 | 3.80 | 280.34 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 2.48 | 961.18 | 20.70 | 1536.02 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 6.07 | 142.28 | 6.60 | 147.09 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 3.45 | 1100.72 | 22.70 | 1755.76 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.75 | 264.22 | 12.20 | 297.50 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 101.95 | 0.20 | 102.51 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.45 | 120.47 | 4.00 | 122.18 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.06 | 104.98 | 0.20 | 105.62 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.08 | 107.64 | 0.20 | 107.76 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.00 | 51.79 | 0.00 | 51.79 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 103.04 | 0.20 | 103.12 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 104.45 | 0.20 | 104.87 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 103.71 | 0.10 | 104.86 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 104.40 | 0.20 | 104.53 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.06 | 103.38 | 0.20 | 103.42 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 102.22 | 0.20 | 102.76 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 104.15 | 0.20 | 104.54 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 28.27 | 121.87 | 33.30 | 123.91 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.06 | 102.49 | 0.20 | 102.65 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.08 | 105.35 | 0.20 | 105.35 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.00 | 96.29 | 0.00 | 96.29 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.00 | 62.85 | 0.00 | 62.85 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 103.32 | 0.20 | 103.70 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 105.64 | 0.20 | 105.75 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.07 | 101.52 | 0.20 | 102.32 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 105.09 | 0.20 | 105.88 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.06 | 103.66 | 0.20 | 103.68 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 102.96 | 0.20 | 103.61 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.03 | 102.86 | 0.20 | 104.09 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 114.96 | 132.95 | 131.91 | 140.13 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.06 | 104.62 | 0.20 | 105.39 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.08 | 107.59 | 0.20 | 107.81 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.00 | 97.87 | 0.00 | 97.87 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 101.10 | 0.20 | 101.40 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 104.12 | 0.20 | 104.26 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 104.14 | 0.20 | 104.40 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 104.15 | 0.20 | 104.24 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.05 | 105.11 | 0.30 | 105.45 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 104.24 | 0.20 | 105.07 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.06 | 104.51 | 0.20 | 104.80 |
