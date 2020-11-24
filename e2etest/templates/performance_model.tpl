## Performance Model for each test cases

  | TeseCase | Data Rate | Instance Type | Avg CPU Usage | Avg Memory Usage | Collection Period | Commit Id  | Testing AMi  |   |   |
  |----------|-----------|---------------|---------------|------------------|-------------------|------------|---|---|---|
{% for performance_model in performance_models %}
  |{{ performance_model.testcase }} | {{ performance_model.dataRate }} | {{ performance_model.instanceType }} | {{ performance_model.avgCpu }} | {{ performance_model.avgMem }} | {{ performance_model.collectionPeriod }} | {{ performance_model.commitId }} | {{ performance_model.testingAmi }} | |
{% endfor %}
 
