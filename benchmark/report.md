## Performance Report

**Version:** [v0.47.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.47.0)

**Commit ID:** [4bdf5c841ea1b654b2cf753e8a35c06488bc5845](https://github.com/aws-observability/aws-otel-collector/commit/4bdf5c841ea1b654b2cf753e8a35c06488bc5845)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.57 | 133.69 | 0.80 | 135.32 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.21 | 119.74 | 0.40 | 121.46 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 9.52 | 140.08 | 10.80 | 143.92 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.19 | 112.86 | 0.30 | 114.06 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.22 | 125.31 | 0.40 | 127.63 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.37 | 122.31 | 0.60 | 123.01 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.16 | 120.75 | 0.50 | 128.20 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.06 | 125.98 | 0.30 | 128.70 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 101.90 | 0.10 | 102.65 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 1.54 | 135.42 | 1.80 | 139.44 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 2.47 | 134.34 | 2.70 | 136.39 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.48 | 130.95 | 0.70 | 132.17 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.42 | 125.35 | 0.60 | 128.67 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.58 | 133.35 | 0.70 | 135.09 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 1.26 | 129.57 | 1.60 | 131.25 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.40 | 125.32 | 1.00 | 136.24 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.61 | 158.22 | 1.10 | 165.40 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 102.21 | 0.10 | 102.65 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 5.35 | 142.23 | 5.90 | 147.99 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 8.28 | 146.57 | 8.80 | 149.21 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 1.60 | 143.10 | 2.20 | 148.16 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 1.33 | 139.23 | 2.40 | 207.32 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 1.24 | 136.27 | 2.50 | 143.00 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 5.40 | 138.59 | 6.70 | 141.28 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 1.28 | 125.39 | 4.20 | 140.85 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 3.74 | 276.39 | 6.80 | 316.31 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 103.60 | 0.10 | 103.92 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 2.21 | 120.85 | 3.50 | 122.00 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 103.61 | 0.20 | 103.70 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 106.32 | 0.20 | 106.63 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.20 | 110.08 | 0.30 | 111.89 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 109.11 | 0.20 | 111.55 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 102.42 | 0.20 | 102.74 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 103.55 | 0.20 | 104.08 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 103.32 | 0.10 | 103.60 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 102.55 | 0.30 | 102.58 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 102.51 | 0.20 | 102.63 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 103.62 | 0.20 | 103.94 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.05 | 101.49 | 0.20 | 101.69 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 18.20 | 120.76 | 26.70 | 123.16 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.03 | 103.73 | 0.20 | 104.17 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 105.80 | 0.20 | 106.25 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.07 | 108.26 | 0.20 | 110.23 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.20 | 112.24 | 0.40 | 114.51 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 103.04 | 0.20 | 103.08 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 103.20 | 0.10 | 103.55 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 103.37 | 0.30 | 103.94 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 104.00 | 0.20 | 104.05 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.03 | 102.80 | 0.20 | 103.02 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 103.00 | 0.20 | 103.26 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.03 | 103.07 | 0.20 | 103.58 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 78.02 | 131.56 | 113.40 | 137.36 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 102.56 | 0.20 | 102.96 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 106.11 | 0.20 | 106.80 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.07 | 110.04 | 0.20 | 110.75 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 110.34 | 0.20 | 112.37 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 101.80 | 0.20 | 101.83 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 104.91 | 0.20 | 105.41 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 104.34 | 0.20 | 104.87 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 103.92 | 0.10 | 104.12 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 102.40 | 0.10 | 102.72 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 102.68 | 0.20 | 102.79 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 103.16 | 0.20 | 103.17 |
