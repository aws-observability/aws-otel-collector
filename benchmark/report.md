## Performance Report

**Version:** [v0.43.3](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.43.3)

**Commit ID:** [2a9454e541d7b44199cf4ebd80f706a637065098](https://github.com/aws-observability/aws-otel-collector/commit/2a9454e541d7b44199cf4ebd80f706a637065098)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.46 | 111.00 | 0.60 | 112.32 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 1.42 | 117.25 | 2.30 | 118.97 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 8.02 | 127.72 | 10.90 | 134.79 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.19 | 104.05 | 0.40 | 105.95 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.22 | 111.70 | 0.30 | 113.33 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.39 | 107.45 | 0.50 | 108.40 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.24 | 112.92 | 0.40 | 116.11 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.08 | 110.24 | 0.40 | 113.49 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 90.27 | 0.20 | 90.65 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 2.11 | 120.44 | 2.30 | 122.52 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.34 | 120.40 | 0.70 | 122.46 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.48 | 121.69 | 0.70 | 123.13 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.46 | 112.76 | 0.60 | 116.09 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.56 | 125.08 | 0.70 | 130.17 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 1.62 | 117.75 | 1.80 | 120.65 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.85 | 124.66 | 1.10 | 128.09 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.71 | 141.03 | 1.40 | 149.51 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 90.78 | 0.10 | 91.37 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 9.90 | 136.32 | 10.50 | 141.02 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 8.94 | 131.74 | 9.40 | 134.69 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 1.73 | 124.76 | 2.00 | 129.82 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 1.63 | 129.72 | 5.00 | 299.86 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 1.93 | 122.91 | 2.10 | 127.97 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 7.01 | 127.21 | 7.30 | 130.67 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 3.75 | 123.90 | 4.00 | 128.04 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 3.87 | 260.04 | 7.60 | 287.54 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 88.96 | 0.10 | 89.31 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.09 | 107.38 | 4.30 | 108.78 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 88.57 | 0.20 | 89.23 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 96.00 | 0.20 | 97.23 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 96.60 | 0.20 | 98.53 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 96.64 | 0.20 | 98.02 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 90.92 | 0.20 | 91.57 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 91.00 | 0.10 | 91.88 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 90.97 | 0.20 | 91.69 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 90.01 | 0.10 | 90.24 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.03 | 91.15 | 0.20 | 91.78 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 89.11 | 0.10 | 89.82 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 89.46 | 0.20 | 90.51 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 18.35 | 112.06 | 19.00 | 113.57 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 91.63 | 0.20 | 91.89 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 90.94 | 0.20 | 91.37 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.11 | 97.05 | 0.20 | 98.64 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.12 | 96.52 | 0.30 | 97.98 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 90.30 | 0.20 | 91.12 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 92.71 | 0.20 | 94.01 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 92.22 | 0.20 | 92.92 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 91.99 | 0.20 | 92.79 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.05 | 91.32 | 0.20 | 92.20 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 90.51 | 0.20 | 91.54 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 91.66 | 0.20 | 91.69 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 26.36 | 126.01 | 27.80 | 133.54 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 89.07 | 0.20 | 89.77 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 95.81 | 0.20 | 96.91 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 96.28 | 0.20 | 98.14 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 96.99 | 0.20 | 98.66 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 90.22 | 0.20 | 91.29 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 92.17 | 0.20 | 93.21 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 91.58 | 0.20 | 92.20 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 90.96 | 0.20 | 91.71 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 90.12 | 0.20 | 90.32 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 91.29 | 0.10 | 92.34 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 92.28 | 0.10 | 92.80 |
