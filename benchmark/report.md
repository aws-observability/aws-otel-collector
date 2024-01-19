## Performance Report

**Version:** [v0.37.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.37.0)

**Commit ID:** [9d465949b015ffaea681a5227ebff35398a46bb5](https://github.com/aws-observability/aws-otel-collector/commit/9d465949b015ffaea681a5227ebff35398a46bb5)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.51 | 96.82 | 0.60 | 98.52 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.19 | 96.43 | 0.40 | 98.45 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.18 | 96.12 | 0.30 | 96.94 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.16 | 97.78 | 0.50 | 110.15 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.20 | 99.48 | 0.30 | 102.27 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.38 | 93.76 | 0.50 | 94.04 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.22 | 99.31 | 0.30 | 103.38 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.08 | 95.93 | 0.30 | 97.76 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 78.82 | 0.10 | 79.08 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 2.33 | 101.76 | 2.50 | 103.50 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 2.74 | 101.39 | 3.00 | 102.79 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 11.41 | 115.51 | 11.80 | 119.82 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.45 | 227.41 | 3.50 | 347.89 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.52 | 101.78 | 0.70 | 103.60 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 1.76 | 96.78 | 1.90 | 97.69 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.85 | 102.43 | 1.00 | 105.08 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.75 | 119.51 | 1.40 | 124.52 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 77.43 | 0.10 | 78.48 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 10.77 | 119.96 | 11.20 | 123.96 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 11.17 | 115.20 | 12.20 | 119.05 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 1.75 | 100.23 | 2.00 | 104.25 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 1.75 | 786.54 | 12.70 | 1461.71 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 2.03 | 101.92 | 2.30 | 105.02 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 8.68 | 110.93 | 9.00 | 112.66 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 3.75 | 103.48 | 4.00 | 106.65 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.80 | 239.41 | 7.80 | 254.92 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 77.38 | 0.10 | 78.39 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.64 | 94.36 | 3.80 | 95.15 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 79.56 | 0.20 | 79.68 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 81.68 | 0.20 | 81.86 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 84.74 | 0.20 | 87.06 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 84.77 | 0.20 | 86.56 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 77.47 | 0.20 | 78.41 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 79.19 | 0.20 | 79.43 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 77.31 | 0.20 | 77.47 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 77.73 | 0.10 | 77.73 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 78.78 | 0.10 | 78.99 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 77.00 | 0.20 | 78.75 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 80.27 | 0.10 | 80.29 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.66 | 98.22 | 20.50 | 100.31 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 78.97 | 0.20 | 80.13 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 80.88 | 0.20 | 80.91 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 85.58 | 0.20 | 88.10 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 83.50 | 0.20 | 85.23 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 77.67 | 0.10 | 79.09 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 81.16 | 0.10 | 81.45 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 79.58 | 0.20 | 80.72 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 80.15 | 0.10 | 80.36 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 77.43 | 0.20 | 78.02 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 78.33 | 0.10 | 78.58 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 79.30 | 0.20 | 79.43 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 26.85 | 109.15 | 28.10 | 115.93 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.03 | 80.89 | 0.20 | 81.58 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 83.05 | 0.20 | 83.42 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.16 | 85.28 | 0.30 | 87.11 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.17 | 85.00 | 0.30 | 86.81 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 77.75 | 0.20 | 77.86 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 79.09 | 0.10 | 79.78 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 80.86 | 0.10 | 81.59 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 78.34 | 0.20 | 79.07 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 79.35 | 0.10 | 80.22 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 77.85 | 0.20 | 79.32 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.05 | 76.73 | 0.20 | 78.16 |
