extensions:

receivers:
  otlp:
    protocols:
      grpc:
        endpoint: 0.0.0.0:55680


processors:

exporters:
  logging:
    loglevel: debug
  awsxray:
    local_mode: true
    region: '${region}'
  awsemf:
    region: '${region}'

service:
  pipelines:
    traces:
      receivers: [otlp]
      exporters: [logging, awsxray]
    metrics:
      receivers: [otlp]
      exporters: [logging, awsemf]
