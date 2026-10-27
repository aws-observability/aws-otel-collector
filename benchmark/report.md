## Performance Report

**Version:** [v0.51.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.51.0)

**Commit ID:** [daa8f56a927136a9b1e09d4223cbf0085fd40a7e](https://github.com/aws-observability/aws-otel-collector/commit/daa8f56a927136a9b1e09d4223cbf0085fd40a7e)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.52 | 140.44 | 0.70 | 144.42 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 7.55 | 143.54 | 9.00 | 150.01 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 1.68 | 137.81 | 2.40 | 141.43 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 111.62 | 0.30 | 113.64 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.20 | 130.42 | 0.40 | 135.39 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.42 | 130.40 | 0.70 | 133.22 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.23 | 133.46 | 0.60 | 140.59 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.15 | 130.57 | 0.60 | 136.22 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.00 | 108.15 | 0.10 | 108.87 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 1.19 | 143.05 | 1.60 | 149.41 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.42 | 138.38 | 0.60 | 142.25 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.36 | 139.21 | 0.60 | 143.68 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.37 | 135.05 | 0.70 | 157.18 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.52 | 141.83 | 0.80 | 148.88 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 1.40 | 138.07 | 1.70 | 140.08 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.75 | 144.11 | 0.90 | 148.03 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.74 | 176.41 | 1.80 | 189.14 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 107.31 | 0.10 | 107.80 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 4.12 | 148.99 | 4.90 | 158.15 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 1.54 | 149.48 | 2.00 | 154.39 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 7.17 | 149.31 | 8.20 | 156.70 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 1.30 | 156.33 | 2.40 | 297.62 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 2.06 | 148.75 | 2.30 | 154.74 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 5.42 | 143.77 | 6.30 | 150.55 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 3.22 | 145.07 | 3.70 | 153.76 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.21 | 301.69 | 9.50 | 343.34 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 108.30 | 0.10 | 108.69 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 2.44 | 124.36 | 3.10 | 126.57 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.05 | 108.04 | 0.20 | 108.48 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 112.77 | 0.20 | 113.98 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 113.70 | 0.20 | 115.09 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.07 | 114.22 | 0.20 | 117.56 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 109.76 | 0.20 | 110.70 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 108.50 | 0.10 | 108.74 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 107.09 | 0.20 | 107.48 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 109.06 | 0.20 | 109.28 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.05 | 108.89 | 0.20 | 109.82 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 107.35 | 0.20 | 108.04 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 17.48 | 125.31 | 24.20 | 128.66 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 109.62 | 0.10 | 109.67 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 112.79 | 0.20 | 113.97 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.21 | 116.48 | 0.40 | 119.55 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.20 | 118.36 | 0.40 | 119.79 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 108.30 | 0.10 | 109.57 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 109.65 | 0.10 | 109.99 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 110.33 | 0.20 | 111.10 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 109.36 | 0.20 | 109.60 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.05 | 107.28 | 0.20 | 107.39 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.05 | 108.02 | 0.30 | 108.63 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 73.41 | 136.62 | 97.80 | 144.81 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 108.43 | 0.20 | 109.15 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 112.94 | 0.20 | 113.26 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.07 | 115.50 | 0.20 | 118.32 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.07 | 115.73 | 0.20 | 117.59 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 108.79 | 0.20 | 109.27 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 109.83 | 0.10 | 109.98 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 108.42 | 0.20 | 108.81 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 109.22 | 0.20 | 110.31 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 108.37 | 0.20 | 108.48 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.06 | 107.25 | 0.20 | 107.41 |
