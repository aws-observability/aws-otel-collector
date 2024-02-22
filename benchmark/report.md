## Performance Report

**Version:** [v0.38.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.38.0)

**Commit ID:** [ef72e5f64406fdbcd2fed0f6bd3d3d7a9a4cf8e2](https://github.com/aws-observability/aws-otel-collector/commit/ef72e5f64406fdbcd2fed0f6bd3d3d7a9a4cf8e2)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.51 | 95.08 | 1.30 | 96.86 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 2.18 | 100.43 | 2.40 | 101.97 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 11.12 | 112.29 | 11.60 | 116.50 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.18 | 91.46 | 0.50 | 94.51 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.22 | 96.15 | 0.40 | 99.31 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.38 | 92.88 | 0.50 | 93.32 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.24 | 97.37 | 0.40 | 102.15 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.08 | 93.05 | 0.30 | 94.49 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 76.52 | 0.10 | 77.50 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 2.18 | 98.62 | 2.30 | 99.84 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.48 | 96.67 | 0.60 | 97.81 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.51 | 97.32 | 0.70 | 98.20 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.43 | 92.35 | 0.60 | 93.38 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.57 | 98.28 | 0.70 | 101.88 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 1.77 | 95.43 | 2.00 | 96.33 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.84 | 98.69 | 1.10 | 101.18 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.75 | 116.38 | 1.40 | 124.89 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 77.06 | 0.10 | 77.15 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 10.42 | 116.83 | 10.80 | 121.54 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 10.63 | 113.54 | 11.00 | 117.10 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 1.83 | 99.83 | 2.10 | 103.61 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 1.44 | 98.60 | 2.40 | 131.47 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 2.06 | 100.43 | 2.30 | 104.69 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 8.54 | 107.11 | 9.00 | 109.77 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 3.89 | 100.22 | 4.30 | 104.55 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.56 | 222.68 | 7.70 | 245.56 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 76.69 | 0.20 | 76.99 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.92 | 89.72 | 4.10 | 91.01 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.03 | 77.55 | 0.20 | 77.68 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 78.29 | 0.20 | 78.78 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 83.46 | 0.20 | 84.20 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 81.91 | 0.20 | 83.15 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 78.38 | 0.10 | 79.70 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 77.66 | 0.20 | 77.75 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 78.37 | 0.20 | 78.66 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 76.48 | 0.10 | 77.76 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 78.10 | 0.20 | 79.82 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 76.22 | 0.20 | 77.03 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 77.27 | 0.10 | 78.07 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 18.14 | 96.18 | 18.60 | 98.01 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 76.50 | 0.20 | 76.69 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 80.68 | 0.20 | 82.29 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.16 | 84.98 | 0.40 | 85.52 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 81.79 | 0.20 | 84.05 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 75.98 | 0.10 | 76.08 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 77.17 | 0.20 | 77.38 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 79.16 | 0.20 | 80.17 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 77.08 | 0.20 | 78.16 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 76.86 | 0.20 | 76.88 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 76.18 | 0.10 | 77.44 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.03 | 75.92 | 0.10 | 77.27 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.71 | 109.40 | 27.10 | 113.80 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.03 | 78.54 | 0.10 | 79.86 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 77.54 | 0.20 | 77.99 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 81.49 | 0.20 | 82.80 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.16 | 83.02 | 0.30 | 83.48 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 79.06 | 0.20 | 79.18 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 78.81 | 0.20 | 78.98 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 80.23 | 0.20 | 81.52 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 78.59 | 0.20 | 79.45 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.05 | 75.94 | 0.20 | 75.94 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 77.07 | 0.20 | 77.28 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.03 | 77.21 | 0.20 | 78.35 |
