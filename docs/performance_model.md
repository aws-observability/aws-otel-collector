## Performance Report

**Commit ID:** [721bc3b0e6c7b12d0aabeaaae87cbc653cd67d38](https://github.com/aws-observability/aws-otel-collector/commit/721bc3b0e6c7b12d0aabeaaae87cbc653cd67d38)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 55.16 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.08 | 59.19 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | otlp | m5.2xlarge | 0.14 | 62.13 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | otlp | m5.2xlarge | 0.10 | 62.22 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 54.14 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.08 | 58.51 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | otlp | m5.2xlarge | 1.22 | 90.34 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | otlp | m5.2xlarge | 1.08 | 94.77 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 53.93 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.09 | 58.75 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | otlp | m5.2xlarge | 6.14 | 233.69 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | otlp | m5.2xlarge | 6.11 | 241.32 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | otlp | m5.2xlarge | 0.04 | 48.08 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 2.58 | 61.86 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.54 | 60.17 |
| otlp | batch | awsxray | otlp_trace | otlp | m5.2xlarge | 3.61 | 61.87 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | splunk_hec | splunkhec_exporter_trace_mock | otlp | m5.2xlarge | 0.00 | 0.00 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | otlp | m5.2xlarge | 0.06 | 49.22 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 20.73 | 62.70 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 24.64 | 4336.50 |
| otlp | batch | awsxray | otlp_trace | otlp | m5.2xlarge | 32.83 | 67.57 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | splunk_hec | splunkhec_exporter_trace_mock | otlp | m5.2xlarge | 0.00 | 0.00 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | otlp | m5.2xlarge | 0.05 | 49.12 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 97.06 | 78.82 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 122.69 | 21910.63 |
| otlp | batch | awsxray | otlp_trace | otlp | m5.2xlarge | 151.14 | 17449.98 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
| otlp | batch | splunk_hec | splunkhec_exporter_trace_mock | otlp | m5.2xlarge | 0.00 | 0.00 |
