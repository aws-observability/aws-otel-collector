## Performance Report

**Commit ID:** [bde007f411d4c3ff2e321f0acedbeeaee48bbb5c](https://github.com/aws-observability/aws-otel-collector/commit/bde007f411d4c3ff2e321f0acedbeeaee48bbb5c)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 58.14 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.09 | 59.64 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.11 | 63.11 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 56.90 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.08 | 61.82 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 59.27 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | otlp | m5.2xlarge | 0.13 | 65.43 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | otlp | m5.2xlarge | 0.11 | 66.34 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 56.81 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.10 | 58.60 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.13 | 63.70 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.09 | 57.32 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.09 | 61.72 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 61.79 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | otlp | m5.2xlarge | 1.23 | 93.88 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | otlp | m5.2xlarge | 1.09 | 98.64 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 57.39 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.09 | 57.90 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.12 | 62.17 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 57.24 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.08 | 60.58 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 60.66 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | otlp | m5.2xlarge | 6.13 | 240.63 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | otlp | m5.2xlarge | 6.00 | 246.47 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | otlp | m5.2xlarge | 0.05 | 52.14 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 3.78 | 69.34 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 2.52 | 87.27 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 3.90 | 62.78 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 2.51 | 65.03 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.48 | 64.44 |
| otlp | batch | awsxray | otlp_trace | otlp | m5.2xlarge | 4.20 | 64.86 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 2.59 | 76.26 |
| otlp | batch | splunk_hec | splunkhec_exporter_trace_mock | otlp | m5.2xlarge | 3.47 | 71.74 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | otlp | m5.2xlarge | 0.05 | 53.23 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 23.85 | 79.03 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 30.04 | 117.43 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 26.67 | 72.43 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 19.62 | 66.03 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 25.15 | 4554.31 |
| otlp | batch | awsxray | otlp_trace | otlp | m5.2xlarge | 32.44 | 70.65 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 20.34 | 79.01 |
| otlp | batch | splunk_hec | splunkhec_exporter_trace_mock | otlp | m5.2xlarge | 25.98 | 68.33 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | otlp | m5.2xlarge | 0.04 | 51.69 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 98.56 | 161.27 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 127.52 | 135.16 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 116.71 | 139.09 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 88.01 | 80.74 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 102.29 | 23165.27 |
| otlp | batch | awsxray | otlp_trace | otlp | m5.2xlarge | 154.22 | 13748.94 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 107.55 | 92.64 |
| otlp | batch | splunk_hec | splunkhec_exporter_trace_mock | otlp | m5.2xlarge | 108.58 | 88.20 |
