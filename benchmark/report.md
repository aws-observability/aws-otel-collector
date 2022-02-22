## Performance Report

**Version:** [v0.17.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.17.0)

**Commit ID:** [7d69912415c27c33cc796c4ca834c7c05332476e](https://github.com/aws-observability/aws-otel-collector/commit/7d69912415c27c33cc796c4ca834c7c05332476e)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.26 | 0.20 | 64.35 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.99 | 0.20 | 63.52 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.86 | 0.20 | 62.98 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.57 | 0.20 | 63.09 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 63.14 | 0.20 | 63.28 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.13 | 0.20 | 63.42 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.12 | 75.77 | 0.30 | 77.34 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 61.49 | 0.10 | 62.19 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.98 | 0.30 | 66.89 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 62.11 | 0.20 | 62.35 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 62.02 | 0.30 | 62.19 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.21 | 0.20 | 63.68 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 61.78 | 0.20 | 62.09 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.13 | 0.20 | 64.28 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.50 | 112.28 | 4.70 | 117.99 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 62.49 | 0.20 | 62.76 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 62.07 | 0.30 | 62.42 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.34 | 0.30 | 62.45 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 63.21 | 0.20 | 64.14 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.85 | 0.30 | 62.31 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 62.42 | 0.30 | 62.53 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 63.22 | 0.20 | 63.49 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.71 | 267.21 | 15.30 | 302.99 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 63.02 | 0.30 | 63.93 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.95 | 120.52 | 6.19 | 152.65 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.34 | 79.49 | 3.80 | 81.66 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.89 | 78.48 | 5.30 | 79.02 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.13 | 93.69 | 3.30 | 95.94 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.15 | 141.13 | 4.50 | 187.75 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 5.42 | 74.42 | 6.50 | 75.86 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.68 | 74.46 | 6.20 | 75.27 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 5.14 | 86.95 | 5.90 | 87.22 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 7.64 | 81.91 | 9.10 | 86.85 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 26.37 | 311.05 | 37.59 | 481.65 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 20.20 | 154.70 | 26.21 | 181.87 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 35.81 | 77.65 | 46.89 | 79.42 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 31.72 | 106.02 | 38.41 | 110.87 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 48.27 | 666.24 | 65.49 | 1235.18 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 37.09 | 77.73 | 42.90 | 78.30 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 41.25 | 78.16 | 54.54 | 79.64 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 26.12 | 89.36 | 27.00 | 89.84 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 34.98 | 448.23 | 40.20 | 531.44 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 37.10 | 495.14 | 51.52 | 748.70 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 20.95 | 173.29 | 27.60 | 201.37 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 142.54 | 87.23 | 151.28 | 88.12 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 129.17 | 128.33 | 153.82 | 137.15 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 124.57 | 3536.39 | 190.03 | 6084.35 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 112.73 | 80.33 | 115.00 | 81.10 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 147.81 | 14440.16 | 461.74 | 25028.29 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 112.87 | 96.49 | 114.24 | 98.76 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 30.65 | 523.16 | 41.30 | 568.31 |
