## Performance Report

**Version:** [v0.36.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.36.0)

**Commit ID:** [254cf918f69868e0f6cca40e041ef2cc0f9257f2](https://github.com/aws-observability/aws-otel-collector/commit/254cf918f69868e0f6cca40e041ef2cc0f9257f2)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.39 | 113.64 | 0.50 | 113.67 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.27 | 106.21 | 0.50 | 106.92 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 9.92 | 128.56 | 10.40 | 130.88 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.19 | 107.45 | 0.60 | 110.15 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.16 | 99.01 | 0.60 | 111.43 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.20 | 102.58 | 0.30 | 103.70 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.35 | 105.94 | 0.50 | 107.33 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.22 | 110.09 | 0.50 | 113.30 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.07 | 106.18 | 0.30 | 108.03 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 80.78 | 0.10 | 81.96 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 1.65 | 122.20 | 2.20 | 123.39 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 1.30 | 109.47 | 1.50 | 110.03 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.45 | 110.94 | 0.60 | 111.73 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 9.91 | 132.36 | 10.20 | 134.89 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.40 | 231.71 | 2.80 | 336.04 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.52 | 116.11 | 0.70 | 119.09 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 1.58 | 109.82 | 1.80 | 110.98 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.75 | 115.73 | 0.90 | 117.01 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.64 | 129.85 | 1.30 | 141.02 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 80.81 | 0.10 | 81.68 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 7.58 | 149.76 | 8.20 | 159.34 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 6.33 | 117.59 | 7.00 | 118.63 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 1.56 | 117.40 | 1.80 | 119.66 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 1.58 | 117.75 | 1.80 | 120.28 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 1.72 | 871.61 | 17.40 | 1457.13 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 1.83 | 117.88 | 2.10 | 119.34 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 7.38 | 123.83 | 7.70 | 126.55 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 3.33 | 118.24 | 3.60 | 121.53 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.44 | 251.98 | 7.20 | 289.29 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 79.60 | 0.10 | 80.81 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.57 | 108.41 | 4.00 | 109.17 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 80.74 | 0.20 | 82.01 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 84.51 | 0.20 | 86.43 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 85.74 | 0.30 | 90.93 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 86.77 | 0.30 | 87.51 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 79.70 | 0.10 | 81.10 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 81.88 | 0.10 | 83.38 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 82.41 | 0.20 | 83.26 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 81.88 | 0.20 | 82.60 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.05 | 81.23 | 0.20 | 82.74 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 80.58 | 0.20 | 82.26 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 82.03 | 0.20 | 83.26 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 17.70 | 110.63 | 18.50 | 111.78 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 80.19 | 0.20 | 81.28 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 84.32 | 0.20 | 85.77 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 86.40 | 0.20 | 91.37 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 86.92 | 0.20 | 90.16 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 82.36 | 0.20 | 83.73 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 82.11 | 0.10 | 83.71 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 82.81 | 0.20 | 84.37 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 80.81 | 0.10 | 81.84 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 80.98 | 0.20 | 81.78 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 81.90 | 0.20 | 82.29 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.05 | 80.64 | 0.20 | 81.61 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.67 | 120.98 | 26.90 | 126.13 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 82.72 | 0.20 | 83.63 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 85.55 | 0.20 | 87.21 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.19 | 86.85 | 0.30 | 91.56 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.18 | 87.90 | 0.30 | 92.02 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 81.85 | 0.20 | 83.02 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 82.10 | 0.20 | 83.37 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 81.99 | 0.20 | 83.59 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 81.70 | 0.20 | 83.03 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 79.64 | 0.20 | 81.03 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 80.85 | 0.20 | 82.05 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 83.90 | 0.20 | 84.89 |
