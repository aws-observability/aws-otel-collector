## Performance Model for each test cases

**Commit ID:** [{{ performance_model.commitId }}](https://github.com/aws-observability/aws-otel-collector/commit/{{ performance_model.commitId }})

**Collection Period:** {{ performance_model.collectionPeriod }}

**Testing AMI:** {{ performance_model.testingAmi }}

  | TeseCase | Data Rate | Instance Type | Avg CPU Usage(Percent) | Avg Memory Usage(Megabytes) |
  |:--------:|:---------:|:-------------:|:----------------------:|:---------------------------:|
{% for performance_model in performance_models %}
  | {{ performance_model.testcase }} | {{ performance_model.dataRate }} | {{ performance_model.instanceType }} | {{ performance_model.avgCpu }} | {{ performance_model.avgMem }} |
{% endfor %}
 
