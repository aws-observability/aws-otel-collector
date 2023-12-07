## Performance Report

**Version:** [v0.35.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.35.0)

**Commit ID:** [b16477bafcacbc2c3e75873a22a970cdac69ca12](https://github.com/aws-observability/aws-otel-collector/commit/b16477bafcacbc2c3e75873a22a970cdac69ca12)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 82.64 | 0.20 | 83.85 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.89 | 0.10 | 82.66 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.17 | 86.87 | 0.30 | 90.52 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.15 | 86.73 | 0.30 | 87.03 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.71 | 0.20 | 82.76 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 79.55 | 0.20 | 80.89 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 81.81 | 0.20 | 83.13 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.79 | 0.20 | 81.89 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.06 | 105.64 | 0.30 | 108.49 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 80.33 | 0.10 | 81.45 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 84.61 | 0.20 | 85.93 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 82.54 | 0.20 | 83.83 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 86.12 | 0.20 | 91.38 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 87.89 | 0.20 | 93.01 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 83.33 | 0.20 | 84.62 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 82.06 | 0.20 | 83.49 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 81.97 | 0.20 | 83.07 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 83.37 | 0.20 | 84.74 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.63 | 128.00 | 1.20 | 139.01 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 80.54 | 0.10 | 81.98 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 83.84 | 0.20 | 84.94 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 82.03 | 0.20 | 83.38 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 87.52 | 0.20 | 91.07 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.07 | 88.14 | 0.30 | 92.88 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 82.24 | 0.20 | 83.24 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.60 | 0.20 | 81.73 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 81.64 | 0.10 | 83.11 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 82.70 | 0.20 | 84.07 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.00 | 251.02 | 7.20 | 286.90 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 79.82 | 0.10 | 80.84 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.56 | 106.88 | 4.10 | 107.18 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 1.31 | 112.11 | 1.50 | 113.14 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.83 | 115.52 | 5.10 | 116.58 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 22.88 | 140.57 | 25.00 | 154.70 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 5.37 | 110.32 | 6.30 | 110.51 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.84 | 114.43 | 4.30 | 116.74 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 4.96 | 127.61 | 5.80 | 147.72 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.54 | 136.96 | 4.00 | 146.61 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.40 | 114.13 | 3.90 | 116.73 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.49 | 106.13 | 3.70 | 107.14 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.48 | 119.24 | 4.00 | 119.60 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 2.42 | 112.98 | 2.70 | 116.11 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 17.68 | 109.37 | 18.40 | 110.65 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 10.77 | 113.37 | 11.10 | 119.62 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 30.08 | 120.62 | 30.70 | 127.59 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 41.51 | 111.28 | 43.90 | 111.87 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 61.77 | 140.81 | 64.80 | 150.01 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 27.79 | 113.62 | 28.20 | 115.09 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 46.65 | 165.04 | 48.00 | 167.24 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 27.99 | 453.49 | 30.20 | 536.25 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 27.44 | 112.95 | 29.10 | 115.06 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 29.39 | 111.42 | 29.80 | 112.28 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 27.19 | 119.53 | 30.00 | 119.60 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 22.37 | 142.32 | 23.30 | 167.52 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.74 | 118.81 | 27.00 | 126.20 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 55.50 | 126.37 | 57.20 | 134.52 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 115.46 | 125.04 | 125.91 | 127.41 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 140.71 | 125.03 | 150.50 | 126.15 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 135.43 | 124.71 | 143.70 | 126.47 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 105.73 | 114.99 | 111.60 | 118.44 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 192.82 | 216.13 | 198.09 | 220.50 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 110.47 | 2002.86 | 129.50 | 2250.01 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 101.55 | 112.21 | 119.11 | 113.18 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 135.45 | 13430.81 | 451.51 | 26804.92 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 109.65 | 125.47 | 115.61 | 126.90 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 100.69 | 245.73 | 104.20 | 374.81 |
