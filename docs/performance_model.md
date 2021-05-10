## Performance Report

**Commit ID:** [644834bc304b699c1957f95fa4e451d032288297](https://github.com/aws-observability/aws-otel-collector/commit/644834bc304b699c1957f95fa4e451d032288297)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 59.67 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.10 | 59.87 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 59.95 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 56.19 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 57.92 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.19 | 64.45 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 60.69 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.13 | 65.52 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 0.10 | 66.43 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 0.61 | 62.57 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 54.01 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 59.41 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.09 | 60.87 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.09 | 59.11 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 54.65 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 59.61 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.20 | 63.63 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 59.89 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.11 | 97.99 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 1.10 | 99.93 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 5.01 | 62.67 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 53.76 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 60.70 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.09 | 60.95 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 60.53 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 55.42 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 59.40 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.19 | 63.45 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 59.84 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 5.69 | 257.41 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 5.81 | 256.18 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 24.80 | 63.07 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 55.02 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 4.24 | 65.98 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.60 | 69.25 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 5.26 | 69.77 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.34 | 83.05 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 3.75 | 71.21 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.08 | 114.05 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 2.56 | 64.11 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.86 | 63.65 |
| otlp | batch | awsxray | otlp_trace | otlp | m5.2xlarge | 4.12 | 66.63 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.05 | 77.90 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 3.52 | 68.85 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 29.90 | 73.73 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 18.22 | 141.06 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 27.91 | 68.55 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 25.02 | 117.19 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 28.12 | 67.94 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 24.57 | 608.33 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 23.60 | 64.48 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 31.08 | 67.65 |
| otlp | batch | awsxray | otlp_trace | otlp | m5.2xlarge | 27.96 | 69.92 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 25.23 | 79.60 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 24.18 | 229.66 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 43.27 | 88.29 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 19.13 | 162.38 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 99.45 | 80.88 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 127.83 | 124.34 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 106.55 | 76.40 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 99.80 | 2857.27 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 111.60 | 69.40 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 120.44 | 12223.77 |
| otlp | batch | awsxray | otlp_trace | otlp | m5.2xlarge | 125.73 | 8093.93 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 96.63 | 82.18 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 23.96 | 334.67 |
