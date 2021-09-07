## Performance Report

**Commit ID:** [8ea305c8cca1b176c888968c3dee84b8001a5ac5](https://github.com/aws-observability/aws-otel-collector/commit/8ea305c8cca1b176c888968c3dee84b8001a5ac5)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.04 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.23 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 58.76 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.53 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.44 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.04 | 60.85 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.64 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 0.15 | 74.39 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 0.13 | 76.64 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 0.54 | 70.50 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.02 | 61.18 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.31 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.96 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.49 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.44 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 60.79 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.05 | 60.22 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 61.58 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 1.37 | 120.44 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 1.21 | 122.43 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 4.71 | 71.25 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 60.17 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 60.44 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 58.83 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.05 | 59.21 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 61.11 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 59.95 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.04 | 60.49 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.04 | 62.86 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | prometheus | m5.2xlarge | 7.35 | 317.89 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | prometheus | m5.2xlarge | 7.27 | 325.66 |
| statsd |  | awsemf, logging | statsd | statsd | m5.2xlarge | 24.34 | 71.20 |
| statsd |  | otlphttp | statsd_mock | statsd | m5.2xlarge | 0.01 | 61.93 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 4.68 | 154.10 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 2.32 | 78.44 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 5.58 | 73.53 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 3.75 | 91.57 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 4.13 | 79.06 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 3.01 | 136.11 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 2.91 | 71.01 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.56 | 71.68 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 4.34 | 74.87 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 3.55 | 86.86 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 6.08 | 80.39 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 29.78 | 560.34 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 18.57 | 149.78 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 32.30 | 78.10 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 26.28 | 104.06 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 30.56 | 75.20 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 28.67 | 743.74 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 26.09 | 73.45 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 31.67 | 76.38 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 33.16 | 79.64 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 25.69 | 87.23 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 31.99 | 471.90 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | xray | m5.2xlarge | 46.48 | 975.44 |
| jaeger | batch | logging, otlphttp | jaeger_mock | jaeger | m5.2xlarge | 18.47 | 171.08 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 113.83 | 87.69 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 122.76 | 167.12 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 130.59 | 84.27 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 126.00 | 3296.86 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 117.34 | 77.42 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 132.93 | 14731.39 |
| otlp | batch | awsxray, logging | otlp_trace | otlp | m5.2xlarge | 152.64 | 15112.16 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 118.19 | 93.64 |
| zipkin | batch | logging, otlphttp | zipkin_mock | zipkin | m5.2xlarge | 34.54 | 530.75 |
