## Performance Report

**Version:** [v0.42.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.42.0)

**Commit ID:** [c37dd16b46f37de1f596f2007b233af051583235](https://github.com/aws-observability/aws-otel-collector/commit/c37dd16b46f37de1f596f2007b233af051583235)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.41 | 128.87 | 0.60 | 129.90 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 2.09 | 134.35 | 2.30 | 135.16 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.18 | 111.73 | 0.40 | 113.01 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.19 | 108.97 | 0.30 | 109.72 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.18 | 117.50 | 0.40 | 118.17 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.37 | 120.23 | 0.60 | 122.54 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.22 | 124.20 | 0.40 | 124.82 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.07 | 120.86 | 0.30 | 123.02 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 93.52 | 0.20 | 96.14 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 1.99 | 137.88 | 2.20 | 140.25 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.46 | 137.82 | 0.70 | 140.03 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 9.31 | 149.11 | 10.30 | 152.29 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.46 | 130.27 | 0.70 | 132.28 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.51 | 144.05 | 0.70 | 146.94 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 1.47 | 136.47 | 1.80 | 140.01 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.77 | 142.79 | 1.00 | 146.36 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.68 | 159.40 | 1.20 | 167.10 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 93.88 | 0.20 | 96.32 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 9.41 | 154.63 | 9.90 | 159.33 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 8.43 | 151.36 | 8.80 | 154.85 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 1.64 | 156.60 | 1.90 | 162.73 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 1.46 | 136.53 | 1.70 | 139.35 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 1.86 | 145.57 | 2.50 | 147.62 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 6.60 | 144.73 | 7.50 | 147.68 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 3.42 | 143.82 | 3.80 | 147.58 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.16 | 275.94 | 7.90 | 318.64 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 93.97 | 0.10 | 96.17 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.88 | 122.75 | 4.20 | 123.36 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.05 | 96.29 | 0.20 | 99.11 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 97.86 | 0.20 | 100.70 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.16 | 100.57 | 0.30 | 104.21 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.08 | 99.24 | 0.30 | 100.27 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 94.20 | 0.20 | 97.08 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 94.66 | 0.20 | 96.91 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 95.05 | 0.30 | 97.56 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 94.62 | 0.20 | 96.67 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 94.19 | 0.20 | 96.64 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 94.93 | 0.20 | 97.04 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 94.76 | 0.20 | 97.64 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 18.02 | 126.72 | 19.30 | 128.62 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 94.58 | 0.20 | 96.65 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.06 | 97.79 | 0.20 | 100.44 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.10 | 99.72 | 0.40 | 101.54 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.06 | 100.86 | 0.30 | 102.85 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 93.90 | 0.20 | 96.98 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 95.22 | 0.20 | 97.55 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 97.64 | 0.20 | 99.91 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 95.25 | 0.20 | 97.17 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 94.20 | 0.20 | 96.87 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 94.66 | 0.20 | 96.85 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 96.26 | 0.20 | 98.67 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 24.37 | 140.73 | 26.10 | 144.19 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 95.17 | 0.20 | 97.46 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.05 | 99.30 | 0.30 | 100.95 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 0.07 | 100.98 | 0.20 | 102.67 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 0.18 | 100.45 | 0.50 | 105.01 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 94.42 | 0.20 | 96.19 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 0.03 | 95.06 | 0.30 | 97.44 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 93.81 | 0.20 | 95.32 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 95.10 | 0.20 | 97.30 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 0.04 | 93.54 | 0.30 | 95.99 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.04 | 95.31 | 0.30 | 97.57 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.05 | 94.19 | 0.20 | 96.37 |
