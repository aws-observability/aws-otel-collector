## Performance Report

**Version:** [v0.43.3](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.43.3)

**Commit ID:** [3a9f6a23632ff047170cf94506d5089fce2f046d](https://github.com/aws-observability/aws-otel-collector/commit/3a9f6a23632ff047170cf94506d5089fce2f046d)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.43 | 111.41 | 0.50 | 111.92 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.11 | 106.90 | 0.30 | 111.62 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.18 | 108.08 | 0.30 | 109.79 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.18 | 102.66 | 0.40 | 104.47 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.20 | 111.59 | 0.40 | 113.99 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.38 | 105.49 | 0.50 | 106.09 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.24 | 111.66 | 0.40 | 114.66 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.08 | 109.64 | 0.40 | 111.34 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 88.06 | 0.10 | 88.24 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 2.11 | 119.70 | 2.30 | 123.26 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.53 | 114.43 | 2.70 | 124.58 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 10.03 | 131.67 | 10.50 | 136.70 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.50 | 117.66 | 1.50 | 145.08 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.54 | 124.74 | 0.80 | 130.10 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 1.58 | 117.30 | 1.80 | 119.63 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.82 | 122.86 | 1.10 | 126.42 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.71 | 140.91 | 1.40 | 151.77 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 92.74 | 0.10 | 93.20 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 9.77 | 136.00 | 10.50 | 141.23 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 8.64 | 133.96 | 16.20 | 167.75 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 1.69 | 123.29 | 1.90 | 127.46 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 1.61 | 120.97 | 1.90 | 130.90 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 1.91 | 124.71 | 2.20 | 128.96 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 7.36 | 126.38 | 8.00 | 129.88 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 3.65 | 122.28 | 4.10 | 127.49 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.40 | 265.71 | 7.70 | 303.42 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 88.17 | 0.10 | 88.57 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.99 | 107.05 | 4.60 | 108.42 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 90.75 | 0.20 | 91.19 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 93.87 | 0.20 | 94.86 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 97.59 | 0.20 | 99.34 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.12 | 96.78 | 0.20 | 97.68 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 91.66 | 0.20 | 92.71 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 93.36 | 0.10 | 93.75 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 90.94 | 0.10 | 91.73 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 89.70 | 0.10 | 89.85 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 92.08 | 0.20 | 92.68 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 90.35 | 0.20 | 90.37 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 90.12 | 0.10 | 91.07 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 19.77 | 110.15 | 20.20 | 112.82 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.05 | 90.60 | 0.10 | 90.77 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 92.14 | 0.30 | 92.73 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.12 | 98.15 | 0.30 | 100.02 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.07 | 97.46 | 0.20 | 98.81 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 91.22 | 0.10 | 91.68 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 93.22 | 0.10 | 93.87 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 90.25 | 0.10 | 91.28 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 91.56 | 0.20 | 91.71 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 91.27 | 0.20 | 91.99 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 90.67 | 0.20 | 91.64 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.03 | 90.26 | 0.20 | 90.42 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 24.91 | 125.76 | 26.10 | 132.06 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 91.64 | 0.20 | 92.61 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 95.88 | 0.20 | 96.89 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 96.46 | 0.20 | 98.04 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 95.54 | 0.20 | 97.07 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 90.48 | 0.20 | 92.32 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.02 | 91.90 | 0.10 | 92.80 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 91.97 | 0.20 | 92.43 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 91.88 | 0.20 | 92.64 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.03 | 90.43 | 0.20 | 91.92 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.03 | 92.40 | 0.20 | 92.83 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.03 | 92.11 | 0.20 | 92.64 |
