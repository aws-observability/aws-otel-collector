## Performance Report

**Commit ID:** [363ca924481559670bc0553055941f0be85a6805](https://github.com/aws-observability/aws-otel-collector/commit/363ca924481559670bc0553055941f0be85a6805)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux


### Metric (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 57.64 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.09 | 59.64 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.12 | 61.78 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 53.65 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 56.94 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.08 | 59.63 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 62.02 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | otlp | m5.2xlarge | 0.12 | 65.08 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | otlp | m5.2xlarge | 0.11 | 65.35 |

### Metric (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 57.02 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.09 | 59.85 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.12 | 63.02 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 52.26 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.06 | 57.49 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.08 | 61.51 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 60.47 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | otlp | m5.2xlarge | 1.21 | 93.18 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | otlp | m5.2xlarge | 1.19 | 97.26 |

### Metric (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| otlp | batch | datadog | datadog_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 56.87 |
| otlp | batch | dynatrace | dynatrace_exporter_metric_mock | otlp | m5.2xlarge | 0.09 | 58.49 |
| otlp | batch | logging, newrelic | newrelic_exporter_metric_mock | otlp | m5.2xlarge | 0.13 | 64.08 |
| otlp | batch | otlp | otlp_grpc_exporter_metric_mock | otlp | m5.2xlarge | 0.08 | 53.42 |
| otlp | batch | otlphttp | otlp_http_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 57.48 |
| otlp | batch | awsemf, logging | otlp_metric | otlp | m5.2xlarge | 0.08 | 61.02 |
| otlp | batch | signalfx | signalfx_exporter_metric_mock | otlp | m5.2xlarge | 0.07 | 58.36 |
| prometheus |  | awsprometheusremotewrite | prometheus_mock | otlp | m5.2xlarge | 6.06 | 233.73 |
| prometheus |  | awsprometheusremotewrite | prometheus_static | otlp | m5.2xlarge | 6.13 | 241.65 |

### Trace (TPS: 100)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | otlp | m5.2xlarge | 0.06 | 53.12 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 5.45 | 67.03 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 2.76 | 83.37 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 4.39 | 63.13 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 2.61 | 162.32 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 2.40 | 62.20 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 3.92 | 62.39 |
| otlp | batch | awsxray | otlp_trace | otlp | m5.2xlarge | 3.61 | 65.42 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 2.36 | 76.03 |
| otlp | batch | splunk_hec | splunkhec_exporter_trace_mock | otlp | m5.2xlarge | 3.80 | 70.82 |

### Trace (TPS: 1000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | otlp | m5.2xlarge | 0.05 | 51.40 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 26.12 | 71.88 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 29.54 | 119.02 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 28.32 | 67.41 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 20.58 | 1074.97 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 20.46 | 63.43 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 35.48 | 69.38 |
| otlp | batch | awsxray | otlp_trace | otlp | m5.2xlarge | 32.66 | 71.27 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 22.27 | 78.60 |
| otlp | batch | splunk_hec | splunkhec_exporter_trace_mock | otlp | m5.2xlarge | 25.60 | 66.98 |

### Trace (TPS: 5000)
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|
| awsxray | batch | awsxray | xrayreceiver | otlp | m5.2xlarge | 0.05 | 51.18 |
| otlp | batch | datadog | datadog_exporter_trace_mock | otlp | m5.2xlarge | 104.95 | 91.53 |
| otlp | batch | logzio | logzio_exporter_trace_mock | otlp | m5.2xlarge | 121.17 | 132.31 |
| otlp | batch | newrelic | newrelic_exporter_trace_mock | otlp | m5.2xlarge | 114.74 | 108.73 |
| otlp | batch | otlp | otlp_grpc_exporter_trace_mock | otlp | m5.2xlarge | 88.78 | 5452.06 |
| otlp | batch | otlphttp | otlp_http_exporter_trace_mock | otlp | m5.2xlarge | 106.61 | 73.82 |
| otlp | batch | awsxray | otlp_mock | otlp | m5.2xlarge | 138.14 | 15793.72 |
| otlp | batch | awsxray | otlp_trace | otlp | m5.2xlarge | 155.67 | 14288.25 |
| otlp | batch | sapm | sapm_exporter_trace_mock | otlp | m5.2xlarge | 104.43 | 92.06 |
| otlp | batch | splunk_hec | splunkhec_exporter_trace_mock | otlp | m5.2xlarge | 117.92 | 80.70 |
