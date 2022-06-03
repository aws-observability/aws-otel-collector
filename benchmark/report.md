## Performance Report

**Version:** [v0.17.1](https://github.com/aws-observability/aws-otel-collector/releases/tag/v0.17.1)

**Commit ID:** [b055a3a024836380ca4fe09b49e6ad62ef5a8008](https://github.com/aws-observability/aws-otel-collector/commit/b055a3a024836380ca4fe09b49e6ad62ef5a8008)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 66.00 | 0.30 | 67.40 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.64 | 0.20 | 61.17 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 63.29 | 0.20 | 63.55 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.41 | 0.20 | 62.60 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 62.24 | 0.20 | 63.09 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.46 | 0.20 | 64.66 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock_awsprw | prometheus | m5.2xlarge | 0.13 | 75.55 | 0.40 | 76.99 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 62.78 | 0.10 | 63.03 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 65.06 | 0.20 | 66.04 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.69 | 0.20 | 62.82 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.59 | 0.20 | 63.27 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.63 | 0.30 | 62.90 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.04 | 62.96 | 0.20 | 63.09 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 64.57 | 0.20 | 65.50 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock_awsprw | prometheus | m5.2xlarge | 1.35 | 111.17 | 3.20 | 119.21 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 63.15 | 0.10 | 63.33 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 62.03 | 0.20 | 62.30 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 64.36 | 0.20 | 64.52 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.70 | 0.20 | 62.09 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.23 | 0.30 | 62.00 |
| otlp | batch | awsemf, logging | otlp_metric_mock | otlp | m5.2xlarge | 0.06 | 63.28 | 0.20 | 63.30 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 63.01 | 0.20 | 63.42 |
| prometheus |  | prometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.00 | 0.00 | 0.00 | 0.00 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock_awsprw | prometheus | m5.2xlarge | 7.47 | 274.22 | 14.90 | 299.94 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 63.13 | 0.20 | 63.82 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 4.66 | 117.71 | 5.50 | 151.68 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 3.08 | 80.00 | 3.41 | 83.25 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.24 | 78.09 | 4.90 | 79.04 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.33 | 93.23 | 3.50 | 94.11 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.20 | 128.58 | 4.60 | 183.03 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 4.92 | 76.10 | 5.60 | 77.28 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.28 | 74.81 | 4.60 | 75.57 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.09 | 86.84 | 3.40 | 87.05 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 6.81 | 81.31 | 7.90 | 85.19 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 25.93 | 327.87 | 37.90 | 525.37 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 21.23 | 159.29 | 26.41 | 184.95 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 40.10 | 77.56 | 42.40 | 79.71 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 26.13 | 105.94 | 27.00 | 110.47 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 28.98 | 638.00 | 39.50 | 1127.33 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 28.99 | 78.00 | 30.21 | 78.60 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 31.46 | 79.17 | 41.59 | 80.22 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 28.05 | 88.92 | 28.99 | 89.46 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 34.55 | 464.26 | 41.14 | 532.17 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver_mock | xray | m5.2xlarge | 36.04 | 457.20 | 47.63 | 740.95 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 21.19 | 172.53 | 27.05 | 203.22 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 124.35 | 86.46 | 126.10 | 87.31 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 109.42 | 129.29 | 110.99 | 135.66 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 112.07 | 3184.48 | 183.02 | 6071.34 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 116.68 | 80.08 | 118.47 | 81.30 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 144.02 | 14249.67 | 476.30 | 25733.26 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 107.57 | 94.87 | 110.96 | 97.71 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 33.52 | 513.40 | 41.32 | 584.53 |
