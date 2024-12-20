## Performance Report

**Version:** [v0.42.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.42.0)

**Commit ID:** [15a91838c0f16541c64fafd79ca3ae34b2c309f9](https://github.com/aws-observability/aws-otel-collector/commit/15a91838c0f16541c64fafd79ca3ae34b2c309f9)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.41 | 126.03 | 0.60 | 127.86 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.24 | 121.32 | 0.80 | 126.09 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.20 | 120.47 | 0.70 | 122.64 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.18 | 106.97 | 0.40 | 107.95 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.20 | 117.17 | 0.60 | 120.25 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.38 | 121.05 | 0.60 | 122.50 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.23 | 122.92 | 0.60 | 126.21 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.07 | 119.52 | 0.30 | 120.95 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 94.45 | 0.10 | 96.33 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 2.00 | 138.96 | 2.20 | 141.17 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 2.49 | 143.29 | 9.10 | 149.97 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.47 | 136.86 | 0.80 | 138.66 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.49 | 137.22 | 1.40 | 161.42 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.53 | 140.11 | 0.70 | 142.34 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 1.46 | 136.51 | 1.70 | 140.26 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.80 | 142.55 | 1.00 | 144.50 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.65 | 157.99 | 1.30 | 168.19 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 93.93 | 0.10 | 95.47 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 9.45 | 152.89 | 9.90 | 160.84 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 7.45 | 149.63 | 9.00 | 154.35 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 10.38 | 150.38 | 10.90 | 154.87 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 1.50 | 137.52 | 1.70 | 140.29 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 1.84 | 144.24 | 2.10 | 147.40 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 6.47 | 143.35 | 6.70 | 146.33 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 3.35 | 142.50 | 3.60 | 146.81 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.23 | 282.36 | 7.30 | 312.61 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 94.31 | 0.10 | 96.18 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.37 | 124.03 | 3.70 | 124.62 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 94.02 | 0.20 | 95.97 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 98.99 | 0.20 | 101.61 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.07 | 100.80 | 0.40 | 104.77 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.17 | 100.40 | 0.40 | 101.83 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 94.93 | 0.20 | 97.10 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 97.79 | 0.20 | 99.96 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 95.93 | 0.20 | 97.35 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 95.08 | 0.20 | 96.46 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.05 | 96.04 | 0.20 | 98.09 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 96.23 | 0.20 | 98.40 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 96.05 | 0.20 | 97.81 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 17.15 | 126.69 | 18.10 | 127.84 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 95.98 | 0.30 | 97.76 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 99.23 | 0.30 | 101.50 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.19 | 100.63 | 0.40 | 105.37 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 100.91 | 0.20 | 104.80 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 96.42 | 0.20 | 98.35 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 95.73 | 0.30 | 98.14 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 96.74 | 0.20 | 99.31 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 96.27 | 0.30 | 98.39 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 95.20 | 0.20 | 97.23 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 93.59 | 0.30 | 95.47 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 96.84 | 0.20 | 99.02 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 24.01 | 140.39 | 25.90 | 143.83 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.05 | 95.50 | 0.20 | 98.07 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 97.93 | 0.20 | 100.09 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.07 | 101.56 | 0.20 | 105.74 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.10 | 100.08 | 0.30 | 105.47 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 95.10 | 0.20 | 96.99 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.04 | 94.78 | 0.20 | 96.24 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 96.44 | 0.20 | 98.30 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 94.35 | 0.20 | 95.65 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 95.54 | 0.20 | 98.00 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 96.00 | 0.20 | 97.56 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.05 | 94.96 | 0.30 | 97.30 |
