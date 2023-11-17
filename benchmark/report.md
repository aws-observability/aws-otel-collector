## Performance Report

**Version:** [v0.35.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.35.0)

**Commit ID:** [c46e130a9a9b199693f280d637d0def0b494b22d](https://github.com/aws-observability/aws-otel-collector/commit/c46e130a9a9b199693f280d637d0def0b494b22d)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 82.47 | 0.20 | 83.59 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 79.31 | 0.20 | 79.94 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.06 | 85.26 | 0.20 | 85.91 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 86.77 | 0.20 | 87.02 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 79.71 | 0.20 | 81.08 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.37 | 0.20 | 81.38 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 79.44 | 0.20 | 80.59 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.38 | 0.20 | 81.38 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.07 | 104.24 | 0.30 | 106.57 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 79.57 | 0.10 | 79.81 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 83.08 | 0.20 | 83.73 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.51 | 0.20 | 82.33 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.07 | 84.89 | 0.20 | 85.08 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.18 | 88.30 | 0.40 | 92.68 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.05 | 0.20 | 80.98 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 79.95 | 0.20 | 80.82 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 80.78 | 0.20 | 81.79 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.50 | 0.20 | 81.21 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.75 | 130.99 | 1.40 | 141.57 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 78.21 | 0.10 | 78.98 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 81.83 | 0.20 | 82.61 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 80.41 | 0.20 | 81.53 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_2_8_1 | otlp | m5.2xlarge | 0.16 | 86.24 | 0.30 | 87.42 |
| otlp | batch | kafka/exporter | kafka_otlp_metric_mock_3_2_0 | otlp | m5.2xlarge | 0.05 | 85.19 | 0.20 | 89.92 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 81.38 | 0.20 | 81.96 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 80.18 | 0.20 | 81.23 |
| otlp | batch | awsemf | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 80.52 | 0.10 | 81.01 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 79.89 | 0.20 | 80.72 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 4.34 | 253.91 | 7.80 | 290.42 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 79.15 | 0.10 | 80.18 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 3.96 | 105.43 | 4.20 | 105.61 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 1.30 | 110.83 | 1.50 | 112.87 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.29 | 114.71 | 5.20 | 116.33 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 19.99 | 137.00 | 25.50 | 156.29 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 23.57 | 139.87 | 25.40 | 153.95 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.46 | 113.54 | 3.60 | 114.84 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 5.04 | 123.38 | 5.70 | 142.68 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.11 | 137.58 | 3.80 | 145.06 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.24 | 116.57 | 3.60 | 117.96 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.49 | 105.96 | 4.40 | 106.35 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.89 | 116.93 | 4.30 | 117.53 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 2.49 | 110.98 | 2.80 | 114.64 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 17.94 | 108.58 | 18.60 | 110.00 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 11.05 | 113.94 | 12.80 | 117.77 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 29.65 | 117.45 | 30.20 | 122.90 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 41.34 | 109.45 | 42.70 | 109.68 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 40.64 | 108.48 | 41.00 | 108.95 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 27.08 | 116.56 | 29.60 | 118.73 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 46.47 | 162.56 | 47.60 | 167.56 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 26.59 | 480.16 | 34.80 | 534.47 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 25.60 | 114.10 | 27.30 | 115.80 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 27.69 | 108.36 | 29.00 | 109.01 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 24.90 | 119.43 | 26.70 | 119.46 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 22.16 | 137.98 | 23.50 | 171.98 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.13 | 119.71 | 26.20 | 127.48 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 55.76 | 122.09 | 60.10 | 132.72 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 109.53 | 121.84 | 114.10 | 126.66 |
| otlp |  | kafka/exporter | kafka_otlp_mock_2_8_1 | otlp | m5.2xlarge | 138.36 | 123.85 | 162.90 | 154.30 |
| otlp |  | kafka/exporter | kafka_otlp_mock_3_2_0 | otlp | m5.2xlarge | 135.06 | 121.64 | 142.60 | 123.45 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 95.04 | 111.48 | 100.00 | 113.74 |
| otlp | groupbytrace, tail_sampling | awsxray | otlp_groupbytrace_tailsampling_mock | otlp | m5.2xlarge | 186.37 | 211.93 | 190.79 | 217.10 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 101.43 | 1988.61 | 125.50 | 2247.20 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 89.52 | 110.21 | 99.39 | 111.43 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 112.84 | 16078.32 | 403.27 | 29201.92 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 91.85 | 124.06 | 103.80 | 126.16 |
| zipkin | batch | otlphttp | zipkin_mock | zipkin | m5.2xlarge | 102.22 | 241.42 | 106.90 | 336.92 |
