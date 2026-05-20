## Performance Report

**Version:** [v0.48.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.48.0)

**Commit ID:** [c07cff3f351e92e68ab9e66ea65edb6763386199](https://github.com/aws-observability/aws-otel-collector/commit/c07cff3f351e92e68ab9e66ea65edb6763386199)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.54 | 142.55 | 0.70 | 143.76 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.00 | 61.12 | 0.00 | 61.12 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.16 | 124.86 | 0.40 | 127.34 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.21 | 131.53 | 0.50 | 137.19 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.41 | 130.66 | 0.50 | 131.79 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.23 | 136.08 | 0.40 | 139.47 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.14 | 136.61 | 0.40 | 140.64 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 109.27 | 0.10 | 109.28 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 1.31 | 145.61 | 1.60 | 148.22 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.38 | 138.76 | 0.60 | 140.62 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.59 | 145.32 | 0.80 | 148.12 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 1.34 | 139.12 | 1.50 | 140.29 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.84 | 145.77 | 0.90 | 149.91 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.69 | 178.38 | 1.20 | 191.85 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 110.12 | 0.10 | 110.19 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 3.88 | 152.35 | 4.90 | 159.64 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.00 | 98.89 | 0.00 | 98.96 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 1.29 | 164.87 | 2.50 | 339.69 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 2.17 | 149.60 | 2.70 | 154.76 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 5.90 | 148.91 | 6.30 | 151.55 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 3.74 | 149.35 | 4.00 | 154.91 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.47 | 315.40 | 7.90 | 355.07 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 110.22 | 0.10 | 110.24 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.20 | 128.18 | 3.60 | 129.20 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 110.85 | 0.20 | 111.03 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 115.50 | 0.20 | 115.71 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 113.01 | 0.10 | 113.40 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 110.37 | 0.30 | 110.53 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 112.23 | 0.10 | 113.33 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 111.45 | 0.20 | 111.51 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.03 | 110.03 | 0.10 | 110.30 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 112.35 | 0.20 | 112.81 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 23.21 | 131.06 | 30.60 | 133.10 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 112.08 | 0.20 | 112.44 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 116.95 | 0.20 | 116.99 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 110.76 | 0.20 | 111.24 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 111.59 | 0.10 | 112.00 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 112.71 | 0.20 | 113.35 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 110.46 | 0.20 | 110.94 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 111.89 | 0.20 | 112.03 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 112.19 | 0.20 | 112.29 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 96.63 | 140.47 | 107.90 | 144.91 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 110.31 | 0.20 | 110.47 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 114.25 | 0.20 | 115.18 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.00 | 95.33 | 0.00 | 95.33 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 111.50 | 0.20 | 111.55 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 110.50 | 0.20 | 110.76 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 111.17 | 0.10 | 112.12 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 112.53 | 0.10 | 113.10 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 111.52 | 0.20 | 111.82 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 110.70 | 0.20 | 110.88 |
