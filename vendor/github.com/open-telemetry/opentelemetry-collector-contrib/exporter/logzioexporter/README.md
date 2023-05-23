# Logzio Exporter

| Status                   |                       |
| ------------------------ | --------------------- |
| Stability                | traces [beta]         |
|                          | logs [beta]           |
| Supported pipeline types | traces, logs |
| Distributions            |  [contrib]   |

This exporter supports sending trace and log data to [Logz.io](https://www.logz.io)

### The following configuration options are supported:
Logz.io exporter is utilizing opentelemetry [exporter helper](https://github.com/open-telemetry/opentelemetry-collector/blob/main/exporter/exporterhelper/README.md) for `retry_on_failure`,`sending_queue` and `timeout` settings
- `account_token` (Required): Your logz.io account token for your tracing or logs account.
- `region` Your logz.io account [region code](https://docs.logz.io/user-guide/accounts/account-region.html#available-regions). Defaults to `us`. Required only if your logz.io region is different than US.
- `endpoint` Custom endpoint, mostly used for dev or testing. This will override the region parameter.
- `retry_on_failure` 
    - `enabled` (default = true)
    - `initial_interval`: Time to wait after the first failure before retrying; ignored if `enabled` is `false`  (default = 5s)
    - `max_interval`: Is the upper bound on backoff; ignored if `enabled` is `false` (default = 30s)
    - `max_elapsed_time`: Is the maximum amount of time spent trying to send a batch; ignored if `enabled` is `false` (default = 300s)
- `sending_queue`
    - `enabled` (default = true)
    - `num_consumers`: Number of consumers that dequeue batches; ignored if `enabled` is `false` (default = 10)
    - `queue_size`: Maximum number of batches kept in memory before dropping; ignored if `enabled` is `false`
      User should calculate this as `num_seconds * requests_per_second` where:
        - `num_seconds` is the number of seconds to buffer in case of a backend outage
        - `requests_per_second` is the average number of requests per seconds.
        - default = 5000
- `timeout`: Time to wait per individual attempt to send data to a backend. default = 30s

#### Tracing example:
* We recommend using `batch` processor. Batching helps better compress the data and reduce the number of outgoing connections required to transmit the data.

```yaml
receivers:
  otlp:
    protocols:
      grpc:
        endpoint: "0.0.0.0:4317"
      http:
        endpoint: "0.0.0.0:4318"
  jaeger:
    protocols:
      thrift_compact:
        endpoint: "0.0.0.0:6831"
      thrift_binary:
        endpoint: "0.0.0.0:6832"
      grpc:
        endpoint: "0.0.0.0:14250"
      thrift_http:
        endpoint: "0.0.0.0:14268"
processors:
  batch:
    send_batch_size: 10000
    timeout: 1s
exporters:
  logzio/traces:
    account_token: "LOGZIOtraceTOKEN"
    region: "us"
service:
  pipelines:
    traces:
      receivers: [ otlp,jaeger ]
      processors: [ batch ]
      exporters: [ logzio/traces ]
  telemetry:
    logs:
      level: "debug"
```
#### Logs example:
* We recommend using `batch` processor. Batching helps better compress the data and reduce the number of outgoing connections required to transmit the data.
* We recommend adding `type` attribute to classify your log records
* We recommend adding `resourcedetection` processor to add metadata to your log records

```yaml
receivers:
  filelog:
    include: [ "/private/var/log/*.log" ] # MacOs system logs
    include_file_name: false
    include_file_path: true 
    operators:
      - type: move
        from: attributes["log.file.path"]
        to: attributes["log_file_path"]
    attributes:
      type: <<your-logzio-type>>
processors:
  batch:
    send_batch_size: 10000
    timeout: 1s
  resourcedetection/system:
    detectors: [ "system" ]
    system:
      hostname_sources: [ "os" ]
exporters:
  logzio/logs:
    account_token: "LOGZIOlogsTOKEN"
    region: "us"
service:
  pipelines:
    logs:
      receivers: [filelog]
      processors: [ resourcedetection/system, batch ]
      exporters: [logzio/logs]
  telemetry:
    logs:
      level: "debug"
```
#### Metrics:
In order to use the Prometheus backend you must use the standard prometheusremotewrite exporter as well. The following [regions](https://docs.logz.io/user-guide/accounts/account-region.html#supported-regions-for-prometheus-metrics) are supported and configured as follows. The Logz.io Listener URL for for your region, configured to use port 8052 for http traffic, or port 8053 for https traffic.
Example:
```yaml
exporters:
  prometheusremotewrite:
    endpoint: "https://listener.logz.io:8053"
    headers:
      Authorization: "Bearer LOGZIOprometheusTOKEN"
```

Putting these both together it would look like this in a full configuration:

```yaml
receivers:
  jaeger:
    protocols:
      thrift_http:
        endpoint: "0.0.0.0:14278"

  prometheus:
    config:
      scrape_configs:
      - job_name: 'ratelimiter'
        scrape_interval: 15s
        static_configs:
          - targets: [ "0.0.0.0:8889" ]

exporters:
  logzio/traces:
    account_token: "LOGZIOtraceTOKEN"
    region: "us"

  prometheusremotewrite:
    endpoint: "https://listener.logz.io:8053"
    headers:
      Authorization: "Bearer LOGZIOprometheusTOKEN"

processors:
  batch:
    send_batch_size: 10000
    timeout: 1s
    
service:
  pipelines:
    traces:
      receivers: [jaeger]
      processors: [batch]
      exporters: [logzio/traces]

    metrics:
      receivers: [prometheus]
      exporters: [prometheusremotewrite]
  
  telemetry:
    logs:
      level: debug #activate debug mode
```


[beta]:https://github.com/open-telemetry/opentelemetry-collector#beta
[contrib]:https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib