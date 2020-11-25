## Performance Report

**Commit ID:** [94aa9f400eefd55274d80bde15f69eaf01b79060](https://github.com/aws-observability/aws-otel-collector/commit/94aa9f400eefd55274d80bde15f69eaf01b79060)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux (amazonlinux2)

| Test Case | Data Rate(tps) |  Data Type | Instance Type | Avg CPU Usage (Percent Per Core) | Avg Memory Usage (Megabytes) |
|:---------:|:---------:|:----------:|:------------:|:-----------------------:|:----------------------------:|
|datadog_exporter_metric_mock | 100 | otlp | m5.2xlarge | 0.05 | 56.74 ||
|datadog_exporter_metric_mock | 1000 | otlp | m5.2xlarge | 0.05 | 56.84 ||     
|datadog_exporter_metric_mock | 5000 | otlp | m5.2xlarge | 0.05 | 57.89 || 
|datadog_exporter_trace_mock | 100 | otlp | m5.2xlarge | 81.10 | 713.83 || 
|datadog_exporter_trace_mock | 1000 | otlp | m5.2xlarge | 779.89 | 3013.75 || 
|datadog_exporter_trace_mock | 5000 | otlp | m5.2xlarge | 777.18 | 3308.37 || 
|dynatrace_exporter_metric_mock | 100 | otlp | m5.2xlarge | 0.05 | 55.46 || 
|dynatrace_exporter_metric_mock | 1000 | otlp | m5.2xlarge | 0.05 | 57.19 || 
|dynatrace_exporter_metric_mock | 5000 | otlp | m5.2xlarge | 0.05 | 57.04 || 
|newrelic_exporter_metric_mock | 100 | otlp | m5.2xlarge | 0.04 | 59.12 || 
|newrelic_exporter_metric_mock | 1000 | otlp | m5.2xlarge | 0.03 | 57.92 || 
|newrelic_exporter_metric_mock | 5000 | otlp | m5.2xlarge | 0.04 | 57.94 || 
|newrelic_exporter_trace_mock | 100 | otlp | m5.2xlarge | 23.10 | 65.86 || 
|newrelic_exporter_trace_mock | 1000 | otlp | m5.2xlarge | 386.52 | 4294.13 || 
|newrelic_exporter_trace_mock | 5000 | otlp | m5.2xlarge | 553.62 | 22009.65 || 
|otlp_http_exporter_metric_mock | 100 | otlp | m5.2xlarge | 0.05 | 56.17 || 
|otlp_http_exporter_metric_mock | 1000 | otlp | m5.2xlarge | 0.04 | 56.24 || 
|otlp_http_exporter_metric_mock | 5000 | otlp | m5.2xlarge | 0.04 | 55.01 || 
|otlp_http_exporter_trace_mock | 100 | otlp | m5.2xlarge | 3.78 | 61.27 || 
|otlp_http_exporter_trace_mock | 1000 | otlp | m5.2xlarge | 38.27 | 64.64 || 
|otlp_http_exporter_trace_mock | 5000 | otlp | m5.2xlarge | 165.91 | 161.49 ||
|otlp_metric | 100 | otlp | m5.2xlarge | 0.05 | 60.56 || 
|otlp_metric | 1000 | otlp | m5.2xlarge | 0.05 | 59.99 || 
|otlp_metric | 5000 | otlp | m5.2xlarge | 0.06 | 58.94 || 
|otlp_mock | 100 | otlp | m5.2xlarge | 6.44 | 62.56 || 
|otlp_mock | 1000 | otlp | m5.2xlarge | 102.28 | 468.19 || 
|otlp_mock | 5000 | otlp | m5.2xlarge | 709.48 | 2315.81 || 
|otlp_trace | 100 | otlp | m5.2xlarge | 10.88 | 69.33 || 
|otlp_trace | 1000 | otlp | m5.2xlarge | 133.30 | 142.79 ||
|otlp_trace | 5000 | otlp | m5.2xlarge | 774.70 | 19360.04 || 
|sapm_exporter_trace_mock | 100 | otlp | m5.2xlarge | 6.11 | 75.01 || 
|sapm_exporter_trace_mock | 1000 | otlp | m5.2xlarge | 58.86 | 76.45 || 
|sapm_exporter_trace_mock | 5000 | otlp | m5.2xlarge | 223.49 | 87.23 || 
|xrayreceiver | 100 | otlp | m5.2xlarge | 0.01 | 51.86 || 
|xrayreceiver | 1000 | otlp | m5.2xlarge | 0.01 | 50.17 || 
|xrayreceiver | 5000 | otlp | m5.2xlarge | 0.02 | 49.07 |
 
