## Performance Report

**Version:** [{{ aoc_version }}](https://github.com/aws-observability/aws-otel-collector/releases/tag/{{ aoc_version }})

**Commit ID:** [{{ commit_id }}](https://github.com/aws-observability/aws-otel-collector/commit/{{ commit_id }})

**Collection Period:** {{ collection_period }} minutes

**Testing AMI:** {{ testing_ami }}

{% for models in models_list %}
### {{ models.data_mode }} (TPS: {{ models.data_rate }})
| Receivers | Processors | Exporters | Test Case | Data Type | Instance Type | Avg CPU Usage (Percent) | Avg Memory Usage (Megabytes) | Max CPU Usage (Percent) | Max Memory Usage (Megabytes) |
|:---------:|:----------:|:---------:|:---------:|:---------:|:-------------:|:-----------------------:|:----------------------------:|:-----------------------:|:----------------------------:|
{% for model in models.models -%}
| {{ model.receivers }} | {{ model.processors }} | {{ model.exporters }} | {{ model.testcase }} | {{ model.dataType }} | {{ model.instanceType }} | {{ model.avgCpu }} | {{ model.avgMem }} | {{ model.maxCpu }} | {{ model.maxMem }} |
{% endfor %}
{%- endfor -%}
