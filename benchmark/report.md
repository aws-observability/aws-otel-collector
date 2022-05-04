## Performance Report

**Version:** [v0.17.0](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.17.0)

**Commit ID:** [c07da1ee7ac0d9c8651ecb2d053ae7c29f1d49d5](https://github.com/aws-observability/aws-otel-collector/commit/c07da1ee7ac0d9c8651ecb2d053ae7c29f1d49d5)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 65.89 | 0.20 | 66.34 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.17 | 0.20 | 64.48 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.25 | 0.20 | 63.39 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.25 | 0.30 | 64.32 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.06 | 63.41 | 0.30 | 63.86 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.83 | 0.20 | 64.95 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.13 | 76.92 | 0.40 | 77.89 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock_awsprw | prometheus | m5.2xlarge | 0.13 | 75.77 | 0.40 | 77.35 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 62.54 | 0.10 | 62.71 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 66.29 | 0.30 | 67.28 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.85 | 0.20 | 62.88 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.21 | 0.20 | 64.82 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.62 | 0.20 | 63.80 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.05 | 62.84 | 0.20 | 63.16 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.93 | 0.20 | 64.25 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.41 | 113.84 | 3.20 | 119.85 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock_awsprw | prometheus | m5.2xlarge | 1.41 | 113.65 | 3.10 | 117.87 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 62.49 | 0.20 | 62.50 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 66.66 | 0.30 | 67.93 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.52 | 0.20 | 63.91 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 64.01 | 0.20 | 64.47 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.84 | 0.20 | 65.69 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.06 | 65.10 | 0.20 | 65.43 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.37 | 69.08 | 2.40 | 71.37 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.47 | 274.44 | 15.00 | 295.92 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock_awsprw | prometheus | m5.2xlarge | 7.25 | 277.36 | 13.11 | 303.41 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 62.54 | 0.20 | 62.83 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.76 | 119.65 | 5.80 | 153.90 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.28 | 81.09 | 3.90 | 83.58 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 5.05 | 79.25 | 5.70 | 79.81 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.66 | 95.76 | 4.00 | 97.81 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.74 | 136.29 | 5.20 | 181.41 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 5.06 | 75.27 | 5.70 | 77.23 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.23 | 74.99 | 4.60 | 76.03 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.89 | 88.86 | 4.50 | 88.99 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 7.26 | 81.87 | 8.40 | 85.03 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.75 | 327.95 | 34.19 | 497.07 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 19.75 | 154.75 | 25.21 | 182.16 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 54.11 | 80.41 | 58.23 | 81.67 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 27.02 | 105.28 | 27.81 | 110.30 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 33.64 | 689.46 | 62.09 | 1169.34 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 31.06 | 77.34 | 33.40 | 77.89 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 34.84 | 80.21 | 45.54 | 81.50 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 35.37 | 90.86 | 55.62 | 91.82 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 33.78 | 467.92 | 40.58 | 569.64 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 0.08 | 73.48 | 1.40 | 83.97 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 20.02 | 177.99 | 30.20 | 207.95 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 137.13 | 86.75 | 153.17 | 87.85 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 119.17 | 129.13 | 131.83 | 137.98 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 138.14 | 3169.43 | 191.97 | 5864.03 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 108.17 | 80.75 | 115.98 | 81.53 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 146.67 | 14146.21 | 457.67 | 26160.35 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 106.57 | 96.22 | 113.75 | 99.01 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 32.75 | 526.40 | 40.70 | 576.69 |
