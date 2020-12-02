## Performance Report

**Commit ID:** [7495b4b3f4943caf14f7dc22f8d12ca0f22d423c](https://github.com/aws-observability/aws-otel-collector/commit/7495b4b3f4943caf14f7dc22f8d12ca0f22d423c)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux

| Test Case | Data Rate |  Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:---------:|:----------:|:------------:|:-----------------------:|:----------------------------:|
| datadog_exporter_metric_mock | 100 | otlp | m5.2xlarge | 0.06 | 57.40 |
| datadog_exporter_metric_mock | 1000 | otlp | m5.2xlarge | 0.06 | 58.12 |
| datadog_exporter_metric_mock | 5000 | otlp | m5.2xlarge | 0.04 | 57.33 |
| datadog_exporter_trace_mock | 100 | otlp | m5.2xlarge | 35.27 | 2145.47 |
| datadog_exporter_trace_mock | 1000 | otlp | m5.2xlarge | 640.30 | 24562.07 |
| datadog_exporter_trace_mock | 5000 | otlp | m5.2xlarge | 383.42 | 22261.06 |
| dynatrace_exporter_metric_mock | 100 | otlp | m5.2xlarge | 0.05 | 57.30 |
| dynatrace_exporter_metric_mock | 1000 | otlp | m5.2xlarge | 0.05 | 57.78 |
| dynatrace_exporter_metric_mock | 5000 | otlp | m5.2xlarge | 0.05 | 56.76 |
| newrelic_exporter_metric_mock | 100 | otlp | m5.2xlarge | 0.04 | 57.92 |
| newrelic_exporter_metric_mock | 1000 | otlp | m5.2xlarge | 0.04 | 59.03 |
| newrelic_exporter_metric_mock | 5000 | otlp | m5.2xlarge | 0.04 | 58.55 |
| newrelic_exporter_trace_mock | 100 | otlp | m5.2xlarge | 26.93 | 64.15 |
| newrelic_exporter_trace_mock | 1000 | otlp | m5.2xlarge | 427.04 | 4085.90 |
| newrelic_exporter_trace_mock | 5000 | otlp | m5.2xlarge | 663.86 | 18212.25 |
| otlp_http_exporter_metric_mock | 100 | otlp | m5.2xlarge | 0.03 | 56.29 |
| otlp_http_exporter_metric_mock | 1000 | otlp | m5.2xlarge | 0.04 | 55.96 |
| otlp_http_exporter_metric_mock | 5000 | otlp | m5.2xlarge | 0.04 | 54.79 |
| otlp_http_exporter_trace_mock | 100 | otlp | m5.2xlarge | 4.03 | 61.80 |
| otlp_http_exporter_trace_mock | 1000 | otlp | m5.2xlarge | 40.53 | 70.18 |
| otlp_http_exporter_trace_mock | 5000 | otlp | m5.2xlarge | 162.69 | 166.07 |
| otlp_metric | 100 | otlp | m5.2xlarge | 0.05 | 60.70 |
| otlp_metric | 1000 | otlp | m5.2xlarge | 0.05 | 60.14 |
| otlp_metric | 5000 | otlp | m5.2xlarge | 0.05 | 59.33 |
| otlp_mock | 100 | otlp | m5.2xlarge | 7.59 | 62.92 |
| otlp_mock | 1000 | otlp | m5.2xlarge | 105.75 | 468.61 |
| otlp_mock | 5000 | otlp | m5.2xlarge | 776.60 | 2894.44 |
| otlp_trace | 100 | otlp | m5.2xlarge | 11.51 | 70.26 |
| otlp_trace | 1000 | otlp | m5.2xlarge | 132.46 | 150.91 |
| otlp_trace | 5000 | otlp | m5.2xlarge | 765.22 | 19315.02 |
| sapm_exporter_trace_mock | 100 | otlp | m5.2xlarge | 5.31 | 74.98 |
| sapm_exporter_trace_mock | 1000 | otlp | m5.2xlarge | 56.06 | 76.69 |
| sapm_exporter_trace_mock | 5000 | otlp | m5.2xlarge | 228.54 | 87.55 |
| signalfx_exporter_metric_mock | 100 | otlp | m5.2xlarge | 0.06 | 59.37 |
| signalfx_exporter_metric_mock | 1000 | otlp | m5.2xlarge | 0.05 | 60.80 |
| signalfx_exporter_metric_mock | 5000 | otlp | m5.2xlarge | 0.04 | 61.41 |
| xrayreceiver | 100 | otlp | m5.2xlarge | 0.02 | 50.69 |
| xrayreceiver | 1000 | otlp | m5.2xlarge | 0.02 | 51.67 |
| xrayreceiver | 5000 | otlp | m5.2xlarge | 0.01 | 52.03 |

 