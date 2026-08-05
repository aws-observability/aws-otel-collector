## Performance Report

**Version:** [v0.49.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.49.0)

**Commit ID:** [621d414c0f50b5610a91b342cc97f16b96d499b2](https://github.com/aws-observability/aws-otel-collector/commit/621d414c0f50b5610a91b342cc97f16b96d499b2)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.51 | 144.32 | 0.80 | 146.78 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 7.16 | 144.71 | 9.00 | 153.64 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 1.81 | 140.97 | 2.30 | 144.51 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.15 | 120.07 | 0.30 | 122.01 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.19 | 131.32 | 0.50 | 137.37 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.40 | 131.54 | 0.60 | 134.31 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.22 | 131.79 | 0.40 | 137.01 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.15 | 134.40 | 0.60 | 138.24 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 110.93 | 0.10 | 112.03 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 1.14 | 144.79 | 1.50 | 151.39 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.48 | 142.40 | 0.70 | 145.35 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.42 | 141.49 | 0.70 | 146.32 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.34 | 132.98 | 0.60 | 137.53 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.51 | 144.56 | 0.90 | 150.67 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 1.44 | 141.78 | 1.70 | 143.99 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.74 | 143.12 | 1.10 | 148.33 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.72 | 170.68 | 1.50 | 187.92 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 112.13 | 0.10 | 112.23 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 4.32 | 153.32 | 5.30 | 162.01 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 1.29 | 145.13 | 1.80 | 158.67 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 7.59 | 154.60 | 8.20 | 160.35 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 1.13 | 148.96 | 2.60 | 255.25 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 1.83 | 147.19 | 2.50 | 157.02 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 5.41 | 146.14 | 6.50 | 153.26 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 3.30 | 146.40 | 4.00 | 154.26 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 3.87 | 317.19 | 8.90 | 400.50 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 112.46 | 0.10 | 113.48 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 2.40 | 125.95 | 3.20 | 128.85 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 111.26 | 0.20 | 111.49 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 115.16 | 0.20 | 116.11 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.20 | 120.76 | 0.40 | 123.32 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 117.58 | 0.20 | 120.06 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 111.28 | 0.20 | 111.76 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 110.73 | 0.20 | 111.37 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 112.96 | 0.20 | 114.63 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 111.66 | 0.10 | 112.02 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 111.60 | 0.10 | 112.07 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 113.50 | 0.20 | 114.61 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 17.44 | 130.18 | 23.10 | 133.39 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.05 | 112.96 | 0.20 | 113.57 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 114.71 | 0.20 | 115.08 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 118.65 | 0.20 | 120.41 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.20 | 119.64 | 0.30 | 121.82 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 111.33 | 0.10 | 111.40 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 113.18 | 0.20 | 113.83 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 109.74 | 0.10 | 110.28 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 112.03 | 0.20 | 112.76 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 111.08 | 0.20 | 111.46 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 112.75 | 0.20 | 112.90 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 76.57 | 137.73 | 104.20 | 145.20 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 111.76 | 0.20 | 112.29 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 116.27 | 0.30 | 117.49 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.07 | 117.15 | 0.20 | 120.12 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 118.07 | 0.20 | 119.60 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 113.43 | 0.10 | 114.20 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 112.29 | 0.10 | 113.25 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 111.16 | 0.20 | 111.65 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 113.19 | 0.20 | 113.59 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 111.41 | 0.20 | 112.22 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 113.16 | 0.20 | 113.84 |
