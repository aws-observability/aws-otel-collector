## Performance Report

**Commit ID:** [7495b4b3f4943caf14f7dc22f8d12ca0f22d423c](https://github.com/aws-observability/aws-otel-collector/commit/7495b4b3f4943caf14f7dc22f8d12ca0f22d423c)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux

| Test Case | Data Rate |  Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:---------:|:----------:|:------------:|:-----------------------:|:----------------------------:|
| datadog_exporter_metric_mock | 100 | otlp | m5.2xlarge | 0.05 | 57.04 |
| datadog_exporter_metric_mock | 1000 | otlp | m5.2xlarge | 0.04 | 58.33 |
| datadog_exporter_metric_mock | 5000 | otlp | m5.2xlarge | 0.05 | 54.84 |
| datadog_exporter_trace_mock | 100 | otlp | m5.2xlarge | 3.84 | 67.66 |
| datadog_exporter_trace_mock | 1000 | otlp | m5.2xlarge | 8.58 | 66.98 |
| datadog_exporter_trace_mock | 5000 | otlp | m5.2xlarge | 21.82 | 70.14 |
| dynatrace_exporter_metric_mock | 100 | otlp | m5.2xlarge | 0.06 | 57.52 |
| dynatrace_exporter_metric_mock | 1000 | otlp | m5.2xlarge | 0.05 | 56.31 |
| dynatrace_exporter_metric_mock | 5000 | otlp | m5.2xlarge | 0.05 | 57.04 |
| newrelic_exporter_metric_mock | 100 | otlp | m5.2xlarge | 0.04 | 57.22 |
| newrelic_exporter_metric_mock | 1000 | otlp | m5.2xlarge | 0.04 | 59.18 |
| newrelic_exporter_metric_mock | 5000 | otlp | m5.2xlarge | 0.04 | 57.79 |
| newrelic_exporter_trace_mock | 100 | otlp | m5.2xlarge | 2.74 | 61.71 |
| newrelic_exporter_trace_mock | 1000 | otlp | m5.2xlarge | 8.18 | 65.24 |
| newrelic_exporter_trace_mock | 5000 | otlp | m5.2xlarge | 28.06 | 72.28 |
| otlp_http_exporter_metric_mock | 100 | otlp | m5.2xlarge | 0.04 | 55.51 |
| otlp_http_exporter_metric_mock | 1000 | otlp | m5.2xlarge | 0.04 | 56.00 |
| otlp_http_exporter_metric_mock | 5000 | otlp | m5.2xlarge | 0.04 | 57.00 |
| otlp_http_exporter_trace_mock | 100 | otlp | m5.2xlarge | 0.58 | 60.51 |
| otlp_http_exporter_trace_mock | 1000 | otlp | m5.2xlarge | 0.90 | 61.51 |
| otlp_http_exporter_trace_mock | 5000 | otlp | m5.2xlarge | 2.11 | 62.35 |
| otlp_metric | 100 | otlp | m5.2xlarge | 0.05 | 60.10 |
| otlp_metric | 1000 | otlp | m5.2xlarge | 0.06 | 59.93 |
| otlp_metric | 5000 | otlp | m5.2xlarge | 0.05 | 60.83 |
| otlp_mock | 100 | otlp | m5.2xlarge | 1.82 | 62.59 |
| otlp_mock | 1000 | otlp | m5.2xlarge | 12.37 | 64.02 |
| otlp_mock | 5000 | otlp | m5.2xlarge | 53.65 | 67.99 |
| otlp_trace | 100 | otlp | m5.2xlarge | 1.72 | 63.61 |
| otlp_trace | 1000 | otlp | m5.2xlarge | 12.03 | 65.36 |
| otlp_trace | 5000 | otlp | m5.2xlarge | 29.92 | 67.24 |
| sapm_exporter_trace_mock | 100 | otlp | m5.2xlarge | 0.82 | 73.77 |
| sapm_exporter_trace_mock | 1000 | otlp | m5.2xlarge | 1.96 | 75.62 |
| sapm_exporter_trace_mock | 5000 | otlp | m5.2xlarge | 6.31 | 77.33 |
| signalfx_exporter_metric_mock | 100 | otlp | m5.2xlarge | 0.05 | 59.87 |
| signalfx_exporter_metric_mock | 1000 | otlp | m5.2xlarge | 0.04 | 60.28 |
| signalfx_exporter_metric_mock | 5000 | otlp | m5.2xlarge | 0.05 | 58.16 |
| xrayreceiver | 100 | otlp | m5.2xlarge | 0.02 | 52.04 |
| xrayreceiver | 1000 | otlp | m5.2xlarge | 0.02 | 51.69 |
| xrayreceiver | 5000 | otlp | m5.2xlarge | 0.02 | 51.11 |

 