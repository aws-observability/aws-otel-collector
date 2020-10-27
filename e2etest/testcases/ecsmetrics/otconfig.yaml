receivers:
  awsecscontainermetrics:
processors:
  filter:
    metrics:
      include:
        match_type: strict
        metric_names:
          - ecs.task.memory.reserved
          - ecs.task.memory.utilized
          - ecs.task.cpu.reserved
          - ecs.task.cpu.utilized
          - ecs.task.network.rate.rx
          - ecs.task.network.rate.tx
          - ecs.task.storage.read_bytes
          - ecs.task.storage.write_bytes
  metricstransform:
    transforms:
      - metric_name: ecs.task.memory.reserved
        action: update
        new_name: ecs.task.memory.reserved_${testing_id}
      - metric_name: ecs.task.memory.utilized
        action: update
        new_name: ecs.task.memory.utilized_${testing_id}
      - metric_name: ecs.task.cpu.reserved
        action: update
        new_name: ecs.task.cpu.reserved_${testing_id}
      - metric_name: ecs.task.cpu.utilized
        action: update
        new_name: ecs.task.cpu.utilized_${testing_id}
      - metric_name: ecs.task.network.rate.rx
        action: update
        new_name: ecs.task.network.rate.rx_${testing_id}
      - metric_name: ecs.task.network.rate.tx
        action: update
        new_name: ecs.task.network.rate.tx_${testing_id}
      - metric_name: ecs.task.storage.read_bytes
        action: update
        new_name: ecs.task.storage.read_bytes_${testing_id}
      - metric_name: ecs.task.storage.write_bytes
        action: update
        new_name: ecs.task.storage.write_bytes_${testing_id}
exporters:
  logging:
    loglevel: debug
  awsemf:
    namespace: '${otel_service_namespace}/${otel_service_name}'
    region: '${region}'

service:
  pipelines:
    metrics:
      receivers: [awsecscontainermetrics]
      processors: [filter, metricstransform]
      exporters: [logging, awsemf]
