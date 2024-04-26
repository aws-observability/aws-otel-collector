## Performance Report

**Version:** [v0.38.2](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.38.2)

**Commit ID:** [935463f0588eed984c1caa42e7a740b87d0d477e](https://github.com/aws-observability/aws-otel-collector/commit/935463f0588eed984c1caa42e7a740b87d0d477e)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.50 | 95.11 | 0.70 | 96.77 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.20 | 92.71 | 0.40 | 94.59 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.21 | 96.11 | 0.80 | 97.33 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.18 | 87.51 | 0.30 | 89.09 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.21 | 96.57 | 0.30 | 98.94 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.39 | 92.03 | 0.60 | 92.71 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.24 | 97.51 | 0.40 | 99.33 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.09 | 93.07 | 0.30 | 93.67 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 76.40 | 0.10 | 77.43 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 2.15 | 100.04 | 2.30 | 101.40 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 2.62 | 101.42 | 2.80 | 102.67 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 9.20 | 110.38 | 12.10 | 117.89 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.45 | 90.42 | 0.60 | 91.61 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.57 | 99.08 | 0.70 | 101.04 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 1.74 | 95.92 | 1.90 | 96.74 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.84 | 97.33 | 1.00 | 102.54 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.67 | 114.81 | 1.40 | 123.38 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 76.81 | 0.10 | 77.18 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 9.49 | 117.05 | 11.30 | 124.35 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 10.60 | 113.91 | 11.00 | 117.86 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 3.74 | 103.93 | 10.90 | 117.06 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 1.45 | 97.34 | 1.70 | 98.92 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 2.16 | 99.75 | 2.50 | 103.80 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 8.53 | 107.50 | 9.30 | 110.11 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 3.88 | 100.06 | 4.00 | 104.32 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.60 | 228.26 | 7.90 | 257.45 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 76.17 | 0.20 | 76.95 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.95 | 92.38 | 4.30 | 93.54 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 79.11 | 0.20 | 80.28 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 80.22 | 0.20 | 81.17 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 80.81 | 0.20 | 82.72 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.18 | 83.70 | 0.30 | 84.33 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 76.48 | 0.20 | 77.40 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 78.34 | 0.20 | 79.24 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 77.71 | 0.20 | 78.84 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 78.01 | 0.10 | 79.34 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 75.84 | 0.20 | 77.01 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 77.23 | 0.10 | 78.87 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 77.00 | 0.20 | 78.40 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.23 | 95.28 | 20.20 | 96.66 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 77.52 | 0.20 | 78.67 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 78.03 | 0.20 | 78.80 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.17 | 83.71 | 0.30 | 84.50 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 81.65 | 0.20 | 83.59 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 77.25 | 0.20 | 78.25 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 77.04 | 0.10 | 77.82 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 80.39 | 0.20 | 81.56 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 76.12 | 0.20 | 76.69 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.05 | 76.80 | 0.20 | 77.52 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 75.87 | 0.10 | 77.15 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.03 | 78.47 | 0.10 | 78.55 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 26.22 | 107.91 | 27.40 | 111.63 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 77.50 | 0.10 | 78.76 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 79.09 | 0.20 | 79.63 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.07 | 81.18 | 0.20 | 82.60 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.08 | 81.35 | 0.30 | 83.79 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 76.40 | 0.20 | 77.46 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 77.71 | 0.10 | 77.81 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 78.93 | 0.10 | 79.93 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 76.31 | 0.20 | 77.93 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 77.72 | 0.10 | 78.98 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 77.25 | 0.20 | 78.46 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.05 | 77.81 | 0.20 | 78.62 |
