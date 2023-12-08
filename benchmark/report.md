## Performance Report

**Version:** [v0.36.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.36.0)

**Commit ID:** [443dfecfef58e785914aab1745638764f349bb5e](https://github.com/aws-observability/aws-otel-collector/commit/443dfecfef58e785914aab1745638764f349bb5e)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 84.03 | 0.20 | 85.66 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.50 | 0.20 | 81.42 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.05 | 87.20 | 0.30 | 91.41 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 88.51 | 0.20 | 92.04 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 83.42 | 0.20 | 84.75 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.10 | 0.20 | 82.62 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 80.55 | 0.10 | 81.95 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.98 | 0.20 | 82.16 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.07 | 105.45 | 0.30 | 108.30 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 80.38 | 0.10 | 81.67 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 84.55 | 0.20 | 85.54 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.34 | 0.10 | 82.70 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.17 | 86.81 | 0.30 | 92.21 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 85.93 | 0.30 | 90.73 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 82.12 | 0.20 | 83.39 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.55 | 0.20 | 82.78 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 80.88 | 0.10 | 82.37 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.57 | 0.20 | 82.85 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.67 | 132.27 | 1.30 | 140.85 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 78.76 | 0.20 | 79.98 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 84.42 | 0.20 | 85.33 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.20 | 0.20 | 82.48 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 86.37 | 0.20 | 92.52 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.16 | 87.53 | 0.30 | 88.67 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 82.05 | 0.20 | 83.16 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 81.13 | 0.20 | 81.91 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 80.85 | 0.20 | 82.21 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 80.82 | 0.20 | 82.23 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.27 | 250.24 | 7.50 | 294.49 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 81.11 | 0.10 | 82.37 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.59 | 107.04 | 3.80 | 107.52 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 1.32 | 111.71 | 1.50 | 113.30 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.22 | 114.41 | 4.60 | 115.55 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 4.92 | 110.32 | 5.10 | 110.65 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 23.49 | 140.92 | 25.20 | 154.30 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.89 | 115.10 | 4.10 | 117.49 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 5.37 | 127.30 | 5.90 | 145.25 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.49 | 136.34 | 3.90 | 145.90 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.99 | 113.49 | 4.40 | 115.81 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.58 | 106.40 | 3.70 | 106.77 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.33 | 118.86 | 3.70 | 119.05 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 2.46 | 113.80 | 2.80 | 117.06 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 18.00 | 110.52 | 18.50 | 111.71 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 10.79 | 115.85 | 11.70 | 122.66 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 29.98 | 117.55 | 30.70 | 121.02 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 42.39 | 111.59 | 43.20 | 112.31 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 44.57 | 116.68 | 61.50 | 146.63 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 28.07 | 112.66 | 28.40 | 114.72 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 48.57 | 165.29 | 50.50 | 169.87 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 27.84 | 454.46 | 29.80 | 533.78 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 27.35 | 113.07 | 32.40 | 115.48 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 32.14 | 110.26 | 32.70 | 111.27 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 27.47 | 121.15 | 30.00 | 121.23 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 22.01 | 136.61 | 23.30 | 165.33 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.52 | 120.19 | 26.40 | 126.97 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 55.49 | 127.66 | 58.40 | 141.96 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 111.14 | 122.00 | 115.50 | 125.78 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 144.24 | 152.42 | 155.71 | 160.43 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 130.28 | 123.81 | 146.90 | 125.24 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 106.96 | 113.84 | 113.39 | 117.14 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 186.17 | 214.02 | 192.50 | 218.79 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 110.39 | 2006.13 | 120.40 | 2248.55 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 101.02 | 112.25 | 115.50 | 113.42 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 125.52 | 12802.97 | 375.47 | 22949.94 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 100.26 | 124.69 | 107.80 | 128.27 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 102.96 | 249.12 | 106.70 | 327.35 |
