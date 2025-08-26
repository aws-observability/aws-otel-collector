## Performance Report

**Version:** [v0.45.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.45.0)

**Commit ID:** [5f33daacb4f915e472e17ad9a0e643dc7fedb690](https://github.com/aws-observability/aws-otel-collector/commit/5f33daacb4f915e472e17ad9a0e643dc7fedb690)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.67 | 125.79 | 0.80 | 126.33 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 2.20 | 132.77 | 2.50 | 133.96 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.31 | 127.17 | 1.00 | 127.34 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.21 | 112.78 | 0.40 | 114.27 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.19 | 124.29 | 0.40 | 125.88 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.10 | 112.63 | 0.60 | 118.24 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.24 | 123.98 | 0.50 | 126.68 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.08 | 123.18 | 0.30 | 124.99 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 101.45 | 0.10 | 102.19 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 1.76 | 133.59 | 2.00 | 137.65 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.49 | 136.69 | 0.60 | 139.75 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.48 | 136.84 | 0.70 | 139.43 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.47 | 129.49 | 0.60 | 131.43 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.54 | 136.19 | 0.70 | 139.47 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.31 | 119.95 | 1.80 | 131.06 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.97 | 136.65 | 1.20 | 139.46 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.71 | 152.40 | 1.30 | 160.15 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 98.67 | 0.10 | 99.00 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 4.88 | 138.47 | 5.10 | 142.40 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 8.76 | 147.84 | 9.90 | 151.40 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 12.19 | 145.87 | 12.80 | 151.18 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 1.59 | 133.11 | 1.80 | 137.64 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.29 | 121.26 | 2.00 | 140.13 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 6.80 | 138.58 | 7.20 | 142.67 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 3.97 | 135.62 | 4.30 | 140.89 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.18 | 260.85 | 7.00 | 270.63 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 100.28 | 0.20 | 101.35 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 2.70 | 119.40 | 3.30 | 120.31 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 100.63 | 0.10 | 101.44 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.07 | 104.65 | 0.30 | 104.94 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.15 | 108.15 | 0.30 | 109.02 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 106.93 | 0.30 | 109.49 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 100.74 | 0.20 | 101.68 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 101.99 | 0.20 | 102.11 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 103.14 | 0.10 | 103.96 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 101.29 | 0.20 | 102.85 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 100.91 | 0.20 | 101.88 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 100.42 | 0.20 | 101.67 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.06 | 99.88 | 0.20 | 101.00 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 20.84 | 119.55 | 24.90 | 121.78 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.03 | 101.71 | 0.20 | 102.59 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 104.49 | 0.20 | 104.57 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.07 | 107.82 | 0.30 | 108.00 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.25 | 108.11 | 0.40 | 110.56 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 100.72 | 0.10 | 101.76 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.04 | 104.43 | 0.20 | 104.89 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 102.27 | 0.20 | 102.63 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 99.99 | 0.20 | 100.80 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 102.52 | 0.20 | 103.95 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 101.08 | 0.20 | 101.61 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 100.10 | 0.20 | 101.11 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 95.29 | 131.19 | 110.40 | 136.21 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 101.15 | 0.10 | 101.95 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 104.25 | 0.20 | 104.62 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.10 | 108.56 | 0.20 | 111.88 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.11 | 108.59 | 0.30 | 111.17 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 102.36 | 0.30 | 102.97 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 103.66 | 0.10 | 104.33 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 102.60 | 0.20 | 103.42 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 100.77 | 0.20 | 101.57 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 100.59 | 0.20 | 101.81 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 102.60 | 0.30 | 103.22 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 100.43 | 0.20 | 100.59 |
