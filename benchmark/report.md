## Performance Report

**Version:** [v0.16.1](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.16.1)

**Commit ID:** [580c175be701bad7b949c189bdd95de99ad98172](https://github.com/aws-observability/aws-otel-collector/commit/580c175be701bad7b949c189bdd95de99ad98172)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 65.36 | 0.50 | 66.24 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.80 | 0.20 | 62.04 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.31 | 0.20 | 62.84 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.06 | 0.20 | 62.03 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 61.20 | 0.20 | 61.41 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.00 | 0.20 | 63.17 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.12 | 75.69 | 0.40 | 77.26 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 61.53 | 0.20 | 61.83 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 65.49 | 0.40 | 66.22 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.12 | 0.20 | 61.19 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.77 | 0.30 | 62.46 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.92 | 0.20 | 62.73 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 62.91 | 0.20 | 63.32 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.60 | 0.20 | 64.07 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.33 | 112.71 | 3.10 | 117.70 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 61.19 | 0.20 | 62.87 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.60 | 0.50 | 65.70 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.82 | 0.20 | 63.45 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.60 | 0.20 | 62.84 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.31 | 0.20 | 62.56 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.06 | 62.86 | 0.20 | 63.41 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 63.97 | 0.20 | 64.16 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.19 | 271.81 | 14.40 | 299.89 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 61.13 | 0.20 | 61.77 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 6.24 | 154.55 | 7.70 | 217.67 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.43 | 80.19 | 2.80 | 82.60 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 5.21 | 77.57 | 5.80 | 78.12 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.76 | 93.94 | 4.00 | 95.37 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 4.11 | 137.16 | 5.50 | 191.80 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.90 | 73.15 | 4.60 | 73.58 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 6.32 | 75.33 | 6.99 | 76.10 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.61 | 87.64 | 5.70 | 87.95 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 6.55 | 82.58 | 7.79 | 86.11 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 27.12 | 504.52 | 39.80 | 862.79 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.34 | 152.84 | 22.48 | 176.46 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 33.00 | 79.44 | 34.20 | 80.22 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 30.54 | 106.71 | 43.60 | 111.12 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 27.87 | 711.29 | 46.89 | 1179.39 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 29.03 | 74.66 | 32.71 | 75.07 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 34.09 | 79.54 | 48.19 | 81.03 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 55.02 | 90.67 | 62.31 | 91.33 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 31.42 | 457.84 | 36.59 | 539.37 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 37.95 | 747.17 | 51.50 | 1258.19 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 16.65 | 172.99 | 22.10 | 210.23 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 128.44 | 85.36 | 138.52 | 86.62 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 116.08 | 127.73 | 121.38 | 135.06 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 129.66 | 3220.44 | 194.82 | 5778.98 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 122.11 | 80.61 | 123.02 | 81.49 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 153.97 | 16378.63 | 477.93 | 28466.94 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 148.14 | 96.20 | 153.49 | 98.20 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 28.21 | 503.89 | 38.84 | 574.99 |
