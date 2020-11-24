## Performance Model for each test cases

  | TeseCase | Data Rate | Instance Type | Avg CPU Usage | Avg Memory Usage | Collection Period | Commit Id  |   |   |   |
  |----------|-----------|---------------|---------------|------------------|-------------------|------------|---|---|---|
{% for performance_model in performance_models %}
  |{{ performance_model.testcase }} | {{ performance_model.data_rate }} | {{ performance_model.instance_type }} | {{ performance_model.avg_cpu_usage }} | {{ performance_model.avg_memory_usage }} | {{ performance_model.collection_period }} | {{ performance_model.commit_id }} |
{% endfor %}
 
