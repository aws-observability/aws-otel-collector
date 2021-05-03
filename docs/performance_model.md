## Performance Report

**Commit ID:** [644834bc304b699c1957f95fa4e451d032288297](https://github.com/aws-observability/aws-otel-collector/commit/644834bc304b699c1957f95fa4e451d032288297)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 58.40 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.11 | 60.70 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 59.45 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 56.12 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 58.54 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.20 | 64.40 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 58.72 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.12 | 66.65 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 0.10 | 65.85 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 0.55 | 63.98 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 53.42 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 59.46 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.10 | 61.64 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.09 | 61.44 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 55.12 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 59.65 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.19 | 63.76 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 59.50 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.23 | 96.79 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 1.11 | 100.40 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 4.96 | 62.33 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 54.41 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 58.87 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.09 | 59.66 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 60.32 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 55.91 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 60.00 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.19 | 64.54 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 60.13 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 5.70 | 258.86 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 5.99 | 258.59 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 24.22 | 62.94 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 54.37 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 4.27 | 66.01 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.05 | 53.85 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 4.94 | 70.06 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.60 | 83.51 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 3.75 | 69.99 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 2.88 | 124.58 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.04 | 65.35 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.69 | 64.48 |
| otlp | batch | awsxray | otlp_trace | otlp | m5.2xlarge | 3.80 | 66.67 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.79 | 76.93 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 53.79 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 26.98 | 74.04 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 54.56 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 29.16 | 68.68 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 25.88 | 117.45 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 26.92 | 67.16 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 24.79 | 663.47 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 22.86 | 65.82 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 27.31 | 67.23 |
| otlp | batch | awsxray | otlp_trace | otlp | m5.2xlarge | 28.76 | 69.97 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 23.01 | 80.44 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 55.50 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 39.26 | 88.00 |
| jaeger | batch | otlphttp | jaeger_mock | jaeger | m5.2xlarge | 0.04 | 55.61 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 112.14 | 82.37 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 118.66 | 124.35 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 106.41 | 76.60 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 102.22 | 3085.89 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 96.60 | 68.64 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 118.25 | 13641.93 |
| otlp | batch | awsxray | otlp_trace | otlp | m5.2xlarge | 117.66 | 12073.02 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 95.26 | 82.36 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 0.04 | 55.47 |
