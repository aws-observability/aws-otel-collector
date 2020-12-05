## Performance Report

**Commit ID:** [3425b55c09c9916534da816c715f216f34469e19](https://github.com/aws-observability/aws-otel-collector/commit/3425b55c09c9916534da816c715f216f34469e19)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 57.05 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 55.36 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.14 | 62.46 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 55.50 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.09 | 59.79 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 59.46 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | otlp | m5.2xlarge | 0.13 | 63.51 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | otlp | m5.2xlarge | 0.11 | 63.87 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 56.85 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 55.67 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.11 | 61.91 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 57.62 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.09 | 59.09 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 57.79 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | otlp | m5.2xlarge | 1.04 | 94.26 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | otlp | m5.2xlarge | 1.17 | 93.79 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 56.47 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 56.26 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.12 | 62.44 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 56.24 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.08 | 59.44 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 61.12 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | otlp | m5.2xlarge | 6.09 | 237.64 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | otlp | m5.2xlarge | 5.97 | 241.65 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | otlp | m5.2xlarge | 0.05 | 50.60 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 5.43 | 64.95 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 4.99 | 62.54 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 3.67 | 61.48 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 4.45 | 61.43 |
| otlp | batch | awsxray | otlp_trace | otlp | m5.2xlarge | 3.33 | 63.41 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 2.84 | 74.38 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | otlp | m5.2xlarge | 0.05 | 49.59 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 28.98 | 68.87 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 27.36 | 66.47 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 19.50 | 63.02 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 31.56 | 64.78 |
| otlp | batch | awsxray | otlp_trace | otlp | m5.2xlarge | 33.38 | 68.09 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 19.68 | 78.89 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | otlp | m5.2xlarge | 0.04 | 51.64 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 93.87 | 92.18 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 112.95 | 106.08 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 93.48 | 68.75 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 137.48 | 95.23 |
| otlp | batch | awsxray | otlp_trace | otlp | m5.2xlarge | 143.84 | 12177.26 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 90.51 | 91.45 |
