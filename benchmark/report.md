## Performance Report

**Version:** [v0.19.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.19.0)

**Commit ID:** [e735fd5f6bffa86b996ded74c736c1206c15db8d](https://github.com/aws-observability/aws-otel-collector/commit/e735fd5f6bffa86b996ded74c736c1206c15db8d)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.21 | 0.20 | 68.31 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 61.42 | 0.20 | 61.60 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 62.27 | 0.20 | 63.12 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 62.03 | 0.20 | 62.50 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 62.06 | 0.20 | 62.25 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 62.73 | 0.20 | 63.59 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.11 | 77.85 | 0.30 | 78.93 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock_awsprw | prometheus | m5.2xlarge | 0.11 | 77.53 | 0.30 | 78.25 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 62.78 | 0.20 | 62.81 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 62.88 | 0.20 | 63.73 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 63.04 | 0.20 | 63.05 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 62.50 | 0.10 | 62.86 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 64.04 | 0.20 | 64.52 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 62.62 | 0.20 | 63.47 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.53 | 0.20 | 64.82 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.12 | 110.64 | 1.90 | 114.05 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock_awsprw | prometheus | m5.2xlarge | 1.19 | 108.45 | 2.10 | 112.97 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 62.63 | 0.20 | 62.89 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.43 | 0.20 | 61.80 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.99 | 0.20 | 64.52 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.72 | 0.20 | 61.96 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 62.69 | 0.20 | 62.79 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.03 | 62.62 | 0.20 | 62.92 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.03 | 63.87 | 0.20 | 64.19 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 6.71 | 258.30 | 11.10 | 277.39 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock_awsprw | prometheus | m5.2xlarge | 6.99 | 248.90 | 11.30 | 280.72 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 61.48 | 0.20 | 62.00 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.54 | 79.82 | 2.80 | 82.13 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 5.03 | 78.36 | 5.30 | 80.81 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.00 | 94.68 | 3.30 | 99.22 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 2.99 | 140.21 | 4.10 | 189.72 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 4.25 | 75.88 | 4.50 | 77.73 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.90 | 76.65 | 4.30 | 78.05 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.39 | 90.35 | 4.10 | 91.91 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 5.26 | 81.33 | 5.70 | 83.98 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 19.88 | 148.64 | 26.51 | 187.81 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 29.76 | 78.48 | 31.31 | 82.05 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 28.20 | 107.84 | 29.11 | 113.59 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 30.26 | 766.51 | 41.90 | 1301.65 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 29.48 | 76.42 | 32.01 | 77.91 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 31.35 | 80.88 | 41.90 | 83.73 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 26.18 | 90.75 | 26.71 | 92.87 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 33.81 | 326.18 | 46.21 | 524.31 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 18.85 | 169.48 | 25.00 | 207.22 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 113.93 | 82.90 | 116.63 | 87.06 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 108.63 | 122.71 | 111.70 | 138.15 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 114.57 | 3152.25 | 172.08 | 5671.11 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 104.92 | 78.38 | 110.38 | 80.41 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 99.93 | 19809.94 | 311.97 | 32322.47 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 98.14 | 97.23 | 100.44 | 100.43 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 32.73 | 414.36 | 53.91 | 518.94 |
