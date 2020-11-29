## Performance Report

**Commit ID:** [78064b3b1bbfa539056c583010d7c45f416e5972](https://github.com/aws-observability/aws-otel-collector/commit/78064b3b1bbfa539056c583010d7c45f416e5972)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux

| Test Case | Data Rate |  Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:---------:|:----------:|:------------:|:-----------------------:|:----------------------------:|
| datadog_exporter_metric_mock | 100 | otlp | m5.2xlarge | 0.04 | 57.79 |
| datadog_exporter_metric_mock | 1000 | otlp | m5.2xlarge | 0.04 | 56.85 |
| datadog_exporter_metric_mock | 5000 | otlp | m5.2xlarge | 0.06 | 57.45 |
| datadog_exporter_trace_mock | 100 | otlp | m5.2xlarge | 83.20 | 749.80 |
| datadog_exporter_trace_mock | 1000 | otlp | m5.2xlarge | 780.12 | 3223.17 |
| datadog_exporter_trace_mock | 5000 | otlp | m5.2xlarge | 777.09 | 3738.08 |
| dynatrace_exporter_metric_mock | 100 | otlp | m5.2xlarge | 0.05 | 57.69 |
| dynatrace_exporter_metric_mock | 1000 | otlp | m5.2xlarge | 0.05 | 55.57 |
| dynatrace_exporter_metric_mock | 5000 | otlp | m5.2xlarge | 0.05 | 57.85 |
| newrelic_exporter_metric_mock | 100 | otlp | m5.2xlarge | 0.03 | 57.90 |
| newrelic_exporter_metric_mock | 1000 | otlp | m5.2xlarge | 0.03 | 55.95 |
| newrelic_exporter_metric_mock | 5000 | otlp | m5.2xlarge | 0.04 | 57.76 |
| newrelic_exporter_trace_mock | 100 | otlp | m5.2xlarge | 24.15 | 64.56 |
| newrelic_exporter_trace_mock | 1000 | otlp | m5.2xlarge | 405.12 | 4643.12 |
| newrelic_exporter_trace_mock | 5000 | otlp | m5.2xlarge | 573.91 | 21837.42 |
| otlp_http_exporter_metric_mock | 100 | otlp | m5.2xlarge | 0.04 | 55.93 |
| otlp_http_exporter_metric_mock | 1000 | otlp | m5.2xlarge | 0.03 | 55.71 |
| otlp_http_exporter_metric_mock | 5000 | otlp | m5.2xlarge | 0.04 | 54.99 |
| otlp_http_exporter_trace_mock | 100 | otlp | m5.2xlarge | 4.02 | 61.64 |
| otlp_http_exporter_trace_mock | 1000 | otlp | m5.2xlarge | 37.23 | 65.36 |
| otlp_http_exporter_trace_mock | 5000 | otlp | m5.2xlarge | 166.33 | 161.98 |
| otlp_metric | 100 | otlp | m5.2xlarge | 0.06 | 59.85 |
| otlp_metric | 1000 | otlp | m5.2xlarge | 0.05 | 59.51 |
| otlp_metric | 5000 | otlp | m5.2xlarge | 0.05 | 60.83 |
| otlp_mock | 100 | otlp | m5.2xlarge | 6.66 | 62.96 |
| otlp_mock | 1000 | otlp | m5.2xlarge | 102.49 | 481.45 |
| otlp_mock | 5000 | otlp | m5.2xlarge | 779.60 | 2549.13 |
| otlp_trace | 100 | otlp | m5.2xlarge | 14.08 | 75.31 |
| otlp_trace | 1000 | otlp | m5.2xlarge | 144.69 | 187.35 |
| otlp_trace | 5000 | otlp | m5.2xlarge | 771.28 | 20664.85 |
| sapm_exporter_trace_mock | 100 | otlp | m5.2xlarge | 5.71 | 76.21 |
| sapm_exporter_trace_mock | 1000 | otlp | m5.2xlarge | 52.73 | 77.25 |
| sapm_exporter_trace_mock | 5000 | otlp | m5.2xlarge | 221.01 | 86.32 |
| signalfx_exporter_metric_mock | 100 | otlp | m5.2xlarge | 0.04 | 61.85 |
| signalfx_exporter_metric_mock | 1000 | otlp | m5.2xlarge | 0.04 | 61.61 |
| signalfx_exporter_metric_mock | 5000 | otlp | m5.2xlarge | 0.04 | 59.78 |
| xrayreceiver | 100 | otlp | m5.2xlarge | 0.02 | 51.24 |
| xrayreceiver | 1000 | otlp | m5.2xlarge | 0.01 | 50.80 |
| xrayreceiver | 5000 | otlp | m5.2xlarge | 0.02 | 52.06 |

 