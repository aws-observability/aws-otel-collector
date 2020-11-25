## Performance Report

**Commit ID:** [94aa9f400eefd55274d80bde15f69eaf01b79060](https://github.com/aws-observability/aws-otel-collector/commit/94aa9f400eefd55274d80bde15f69eaf01b79060)

**Collection Period:** 10 minutes

**Testing AMI:** soaking_linux (amazonlinux2)


### Table 1 Trace (TPS: 100)

| Test Case | Data Rate(tps) |  Data Type | Instance Type | Avg CPU Usage (Percent Per Core) | Avg Memory Usage (Megabytes) |
|:---------:|:---------:|:----------:|:------------:|:-----------------------:|:----------------------------:|
|datadog_exporter_trace_mock (OTLP -> datadog exporter) | 100 | otlp | m5.2xlarge | 81.10 | 713.83 || 
|newrelic_exporter_trace_mock (OTLP -> newRelic exporter)  | 100 | otlp | m5.2xlarge | 23.10 | 65.86 || 
|otlp_http_exporter_trace_mock (OTLP -> http exporter) | 100 | otlp | m5.2xlarge | 3.78 | 61.27 || 
|otlp_mock (OTLP -> xray exporter)| 100 | otlp | m5.2xlarge | 6.44 | 62.56 || 
|otlp_trace (OTLP -> xray exporter true backend)| 100 | otlp | m5.2xlarge | 10.88 | 69.33 || 
|sapm_exporter_trace_mock (OTLP -> sapm exporter) | 100 | otlp | m5.2xlarge | 6.11 | 75.01 || 

### Table 2 Trace (TPS: 1000)

| Test Case | Data Rate(tps) |  Data Type | Instance Type | Avg CPU Usage (Percent Per Core) | Avg Memory Usage (Megabytes) |
|:---------:|:---------:|:----------:|:------------:|:-----------------------:|:----------------------------:|
|datadog_exporter_trace_mock (OTLP -> datadog exporter) | 1000 | otlp | m5.2xlarge | 779.89 | 3013.75 || 
|newrelic_exporter_trace_mock (OTLP -> newRelic exporter)| 1000 | otlp | m5.2xlarge | 386.52 | 4294.13 ||
|otlp_http_exporter_trace_mock (OTLP -> http exporter)| 1000 | otlp | m5.2xlarge | 38.27 | 64.64 ||
|otlp_mock (OTLP -> xray exporter) | 1000 | otlp | m5.2xlarge | 102.28 | 468.19 || 
|otlp_trace (OTLP -> xray exporter true backend)| 1000 | otlp | m5.2xlarge | 133.30 | 142.79 ||
|sapm_exporter_trace_mock (OTLP -> sapm exporter)| 1000 | otlp | m5.2xlarge | 58.86 | 76.45 ||

### Table 3 Trace (TPS: 5000)

| Test Case | Data Rate(tps) |  Data Type | Instance Type | Avg CPU Usage (Percent Per Core) | Avg Memory Usage (Megabytes) |
|:---------:|:---------:|:----------:|:------------:|:-----------------------:|:----------------------------:|
|datadog_exporter_trace_mock (OTLP -> datadog exporter)| 5000 | otlp | m5.2xlarge | 777.18 | 3308.37 || 
|newrelic_exporter_trace_mock (OTLP -> newRelic exporter)| 5000 | otlp | m5.2xlarge | 553.62 | 22009.65 || 
|otlp_http_exporter_trace_mock (OTLP -> http exporter)| 5000 | otlp | m5.2xlarge | 165.91 | 161.49 ||
|otlp_mock (OTLP -> xray exporter) | 5000 | otlp | m5.2xlarge | 709.48 | 2315.81 || 
|otlp_trace (OTLP -> xray exporter true backend) | 5000 | otlp | m5.2xlarge | 774.70 | 19360.04 || 
|sapm_exporter_trace_mock (OTLP -> sapm exporter) | 5000 | otlp | m5.2xlarge | 223.49 | 87.23 || 

### Table 4 Metric (TPS: 100)
| Test Case | Data Rate(tps) |  Data Type | Instance Type | Avg CPU Usage (Percent Per Core) | Avg Memory Usage (Megabytes) |
|:---------:|:---------:|:----------:|:------------:|:-----------------------:|:----------------------------:|
|datadog_exporter_metric_mock (OTLP -> datadog exporter)| 100 | otlp | m5.2xlarge | 0.05 | 56.74 ||
|dynatrace_exporter_metric_mock (OTLP -> dynatrace exporter)| 100 | otlp | m5.2xlarge | 0.05 | 55.46 || 
|newrelic_exporter_metric_mock (OTLP -> newRelic exporter)| 100 | otlp | m5.2xlarge | 0.04 | 59.12 ||
|otlp_http_exporter_metric_mock (OTLP -> http exporter)| 100 | otlp | m5.2xlarge | 0.05 | 56.17 ||
|otlp_emf_metric (OTLP -> emf exporter true backend)| 100 | otlp | m5.2xlarge | 0.05 | 60.56 || 

### Table 5 Metric (TPS: 1000)
| Test Case | Data Rate(tps) |  Data Type | Instance Type | Avg CPU Usage (Percent Per Core) | Avg Memory Usage (Megabytes) |
|:---------:|:---------:|:----------:|:------------:|:-----------------------:|:----------------------------:|
|datadog_exporter_metric_mock (OTLP -> datadog exporter)| 1000 | otlp | m5.2xlarge | 0.05 | 56.84 ||
|dynatrace_exporter_metric_mock (OTLP -> dynatrace exporter)| 1000 | otlp | m5.2xlarge | 0.05 | 57.19 || 
|newrelic_exporter_metric_mock (OTLP -> newRelic exporter)| 1000 | otlp | m5.2xlarge | 0.03 | 57.92 || 
|otlp_http_exporter_metric_mock (OTLP -> http exporter) | 1000 | otlp | m5.2xlarge | 0.04 | 56.24 || 
|otlp_emf_metric (OTLP -> emf exporter true backend) | 1000 | otlp | m5.2xlarge | 0.05 | 59.99 || 

### Table 6 Metric (TPS: 5000)
| Test Case | Data Rate(tps) |  Data Type | Instance Type | Avg CPU Usage (Percent Per Core) | Avg Memory Usage (Megabytes) |
|:---------:|:---------:|:----------:|:------------:|:-----------------------:|:----------------------------:|
|datadog_exporter_metric_mock (OTLP -> datadog exporter)| 5000 | otlp | m5.2xlarge | 0.05 | 57.89 ||
|dynatrace_exporter_metric_mock (OTLP -> dynatrace exporter)| 5000 | otlp | m5.2xlarge | 0.05 | 57.04 || 
|newrelic_exporter_metric_mock (OTLP -> newRelic exporter)| 5000 | otlp | m5.2xlarge | 0.04 | 57.94 || 
|otlp_http_exporter_metric_mock (OTLP -> http exporter)| 5000 | otlp | m5.2xlarge | 0.04 | 55.01 ||
|otlp_emf_metric (OTLP -> emf exporter true backend) | 5000 | otlp | m5.2xlarge | 0.06 | 58.94 ||


 
