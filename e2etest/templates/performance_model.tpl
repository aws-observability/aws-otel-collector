## Performance Report

**Commit ID:** [{{ commit_id }}](https://github.com/aws-observability/aws-otel-collector/commit/{{ commit_id }})

**Collection Period:** {{ collection_period }} minutes

**Testing AMI:** {{ testing_ami }}

| Test Case | Data Rate |  Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) |
|:---------:|:---------:|:----------:|:------------:|:-----------------------:|:----------------------------:|
{% for performance_model in performance_models -%}
| {{ performance_model.testcase }} | {{ performance_model.dataRate }} | {{ performance_model.dataType }} | {{ performance_model.instanceType }} | {{ performance_model.avgCpu }} | {{ performance_model.avgMem }} |
{% endfor %}
 
