## Performance Report

**Version:** [v0.19.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.19.0)

**Commit ID:** [14b0e892278c1733fff92a2f50e79a86517edcd0](https://github.com/aws-observability/aws-otel-collector/commit/14b0e892278c1733fff92a2f50e79a86517edcd0)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 66.51 | 0.20 | 67.57 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 61.71 | 0.20 | 62.14 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.48 | 0.20 | 64.35 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.87 | 0.20 | 64.06 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 62.90 | 0.20 | 63.38 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.78 | 0.20 | 64.34 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.10 | 78.08 | 0.30 | 78.79 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock_awsprw | prometheus | m5.2xlarge | 0.11 | 78.63 | 0.30 | 80.00 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 64.66 | 0.10 | 65.48 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 65.27 | 0.20 | 65.86 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 64.23 | 0.20 | 64.86 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 63.46 | 0.10 | 63.69 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 63.12 | 0.20 | 63.38 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 64.54 | 0.20 | 65.04 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.63 | 0.20 | 64.86 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.14 | 109.99 | 2.00 | 114.77 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock_awsprw | prometheus | m5.2xlarge | 1.30 | 112.29 | 2.10 | 116.65 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 62.78 | 0.20 | 63.28 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.26 | 0.20 | 63.76 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.18 | 0.20 | 64.54 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 63.19 | 0.20 | 63.71 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 64.18 | 0.20 | 64.61 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 62.71 | 0.20 | 62.92 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.54 | 0.20 | 65.99 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.12 | 251.76 | 11.00 | 274.44 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock_awsprw | prometheus | m5.2xlarge | 7.48 | 254.81 | 11.60 | 280.81 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 61.51 | 0.20 | 61.69 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.15 | 83.58 | 15.40 | 87.11 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 3.60 | 79.46 | 3.80 | 83.08 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 4.37 | 76.79 | 4.50 | 78.95 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 2.97 | 135.43 | 4.10 | 184.62 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.80 | 77.28 | 4.00 | 79.41 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.69 | 75.33 | 4.00 | 76.89 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.01 | 89.81 | 3.30 | 91.54 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 6.18 | 81.76 | 19.70 | 86.09 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 26.30 | 149.01 | 42.50 | 193.71 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 29.77 | 85.30 | 30.80 | 88.51 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 31.41 | 77.14 | 32.20 | 79.02 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 26.36 | 749.61 | 37.30 | 1221.14 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 26.56 | 76.51 | 27.20 | 77.89 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 31.30 | 80.43 | 42.30 | 83.63 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 29.15 | 91.16 | 30.30 | 93.25 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 41.51 | 319.63 | 61.80 | 469.19 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 25.79 | 176.66 | 40.60 | 219.55 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 129.03 | 88.68 | 136.40 | 95.03 |
| otlp | batch | logzio/traces | logzio_exporter_trace_mock | otlp | m5.2xlarge | 127.17 | 82.15 | 134.89 | 85.00 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 121.56 | 3128.58 | 184.52 | 5837.60 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 122.80 | 78.76 | 124.00 | 80.38 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 98.81 | 22539.23 | 337.22 | 32665.69 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 114.63 | 97.43 | 116.60 | 100.09 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 40.27 | 431.22 | 59.00 | 630.42 |
