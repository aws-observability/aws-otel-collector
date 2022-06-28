## Performance Report

**Version:** [v0.18.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.18.0)

**Commit ID:** [05dda1444e066bb25a92e2c43edddc0c9f8998da](https://github.com/aws-observability/aws-otel-collector/commit/05dda1444e066bb25a92e2c43edddc0c9f8998da)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 66.67 | 0.20 | 67.92 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.93 | 0.20 | 63.25 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.71 | 0.20 | 63.98 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 62.83 | 0.20 | 62.86 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 62.70 | 0.10 | 63.19 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 64.88 | 0.20 | 65.18 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.09 | 78.03 | 0.30 | 79.72 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock_awsprw | prometheus | m5.2xlarge | 0.11 | 76.96 | 0.30 | 78.86 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 62.00 | 0.20 | 62.37 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.79 | 0.20 | 67.77 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 63.10 | 0.10 | 63.17 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 63.18 | 0.10 | 63.27 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 61.42 | 0.20 | 61.79 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 62.46 | 0.20 | 62.53 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 65.55 | 0.20 | 66.38 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.18 | 109.65 | 2.10 | 113.65 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock_awsprw | prometheus | m5.2xlarge | 1.23 | 111.30 | 2.10 | 114.58 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 62.99 | 0.20 | 63.28 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.55 | 0.20 | 66.27 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.05 | 0.10 | 61.52 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 63.05 | 0.20 | 63.15 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.82 | 0.20 | 63.06 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 63.11 | 0.20 | 63.30 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.63 | 0.20 | 64.79 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 6.74 | 256.62 | 11.40 | 285.84 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock_awsprw | prometheus | m5.2xlarge | 7.40 | 253.83 | 12.50 | 283.48 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 63.09 | 0.10 | 63.34 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.52 | 80.98 | 2.80 | 83.37 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.73 | 77.90 | 5.20 | 80.11 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.29 | 94.22 | 3.70 | 96.94 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.17 | 130.87 | 4.30 | 185.24 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.88 | 76.81 | 4.20 | 78.40 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.64 | 76.02 | 4.00 | 77.73 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.04 | 89.29 | 3.30 | 90.48 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.37 | 79.40 | 5.80 | 84.09 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 19.04 | 146.86 | 26.89 | 189.36 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 31.14 | 77.71 | 33.19 | 79.87 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 30.53 | 108.58 | 31.70 | 113.93 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 28.26 | 718.15 | 39.10 | 1157.57 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 28.54 | 77.70 | 29.11 | 79.15 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 30.63 | 80.44 | 40.21 | 83.30 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 26.57 | 90.96 | 27.41 | 92.70 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 34.43 | 314.45 | 43.59 | 485.45 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 18.42 | 168.26 | 25.80 | 194.59 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 115.61 | 83.42 | 116.82 | 87.07 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 116.85 | 122.38 | 118.86 | 131.90 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 108.56 | 3534.36 | 169.77 | 6095.69 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 115.41 | 78.13 | 117.61 | 80.31 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 100.20 | 20549.76 | 310.79 | 32672.46 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 113.65 | 95.84 | 115.23 | 99.29 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 32.42 | 411.51 | 42.71 | 508.82 |
