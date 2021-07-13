# Container Insight ECS Prometheus

NOTE: This is doc for developing this feature, for user doc please
check [user guide](https://aws-otel.github.io/docs/getting-started/container-insights/ecs-prometheus).

## Links

- [ecsobserver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/observer/ecsobserver)
  discover ECS tasks
- [Integration test](https://github.com/aws-observability/aws-otel-test-framework/pull/308)

## Quick Start

After building your own image `12346.dkr.ecr.us-west-2.amazonaws.com/aoc:ecssd-0.2` you can
use [this cfn](../../deployment-template/ecs/aws-otel-container-insights-prometheus-ec2-deployment-cfn.yaml)
to launch the collector on ECS EC2 cluster.

```bash
export CLUSTER_NAME=aoc-prometheus-dashboard-1
export CREATE_IAM_ROLES=True
export COLLECTOR_IMAGE=12346.dkr.ecr.us-west-2.amazonaws.com/aoc:ecssd-0.2

aws cloudformation create-stack --stack-name AOC-Prometheus-ECS-${CLUSTER_NAME} \
    --template-body file://aws-otel-container-insights-prometheus-ec2-deployment-cfn.yaml \
    --parameters ParameterKey=ClusterName,ParameterValue=${CLUSTER_NAME} \
                 ParameterKey=CreateIAMRoles,ParameterValue=${CREATE_IAM_ROLES} \
                 ParameterKey=CollectorImage,ParameterValue=${COLLECTOR_IMAGE} \
    --capabilities CAPABILITY_NAMED_IAM
```

It will create the following resource:

- SSM parameter
- IAM roles, `AWSOTelRolePrometheusECS` and `AWSOTelExecutionRolePrometheusECS` by default
- ECS task definition and replica service

If you need to test your image frequently, you need a script to update SSM parameter, push the image, scale service down
to 0 and back to 1. The cloudformation stack is a bit slow for iteration. (The script is an exercise for the reader,
hint: aws cli can't upload ssm parameter value from file name)

## Internal

NOTE: some problems (or problematic solutions...) also apply to (are copied from) Container Insights EKS Prometheus.

To understand the codebase, check README
in [ecsobserver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/observer/ecsobserver)
. You can also
use [cloudwatch-agent](https://github.com/aws/amazon-cloudwatch-agent/tree/master/internal/ecsservicediscovery) as
reference.

### Label, Relabel and Dimension

Labels are key value pairs e.g. `env=prod`. They are called Dimension in CloudWatch. There is no direct translation from
label to dimension because CloudWatch does not support too many dimensions. Metrics declaration allow picking some
labels as dimension. There is also dimension rollup, but we disable it using `NoDimensionRollup`.

For builtin dashboard to work, specific metric dimensions are required. In `ecsobserver`, we export labels
with `__meta_ecs_` prefix (e.g. `__meta_ecs_task_definition_family`), which is different from cloudwatch-agent.
Using `__` prefix is more popular in prometheus's builtin discovery implementations, so we followed that instead when
porting the discovery logic. For getting a dimension like `TaskDefinitionFamily` in CloudWatch we go through two steps:

- prometheus relabel ensures we carry the label down the pipeline, otherwise all `__` are
  dropped. `__meta_ecs_task_definition_family=MyTask` becomes `TaskDefinitionFamily=MyTask`
- emf's metrics declaration picks some label as metric dimension, other labels becomes structured log field. `MyTask`
  becomes a dimension value that will show up in dashboard when you slice and dice.

```yaml
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: "ecssd"
          relabel_configs:  # Relabel here because label with __ prefix will be dropped by receiver.
            - source_labels: [ __meta_ecs_task_definition_family ] # TaskDefinitionFamily
              action: replace
              target_label: TaskDefinitionFamily
              
exporters:
  awsemf:
    metric_declarations:
      - dimensions: [ [ ClusterName, TaskDefinitionFamily, ServiceName ] ] # dimension names are same as our relabeled keys.
        label_matchers:
          - label_names:
              - ServiceName
            regex: '^.*nginx-service$'
        metric_name_selectors:
          - "^nginx_.*$"
```

### job Label

We allow user to specify different names using `job_name` in config. They are NOT exported as `job` and uses the value
from `job_label_name` as exported label key (e.g. `prometheus_job`). Then in processors we use `metricstransform`
processor to rename `promethus_job` back to `job`.

Why don't we just use `job` directly? Short answer is prometheus receiver does not support specifying `job` in discovery
output. We use `file_sd` as the actual discovery implementation to bridge our discovery result, all the targets are
under the job `ecssd` in prometheus config. However, prometheus receiver does not behave exactly like prometheus, it
relies on job name for detecting metrics type. If we export target with job `nginx-prometheus-exporter`, receiver will
look up metadata cache using `nginx-prometheus-exporter` while the only job in cache is `ecssd`, the result is metric
type unknown. The comment in
[this PR](https://github.com/open-telemetry/opentelemetry-collector-contrib/pull/3785#discussion_r654028642)
gives more detail and links to upstream issue.

```yaml
extensions:
  ecs_observer: # extension type is ecs_observer
    # custom name for 'job' so we can rename it back to 'job' using metricstransform processor
    job_label_name: prometheus_job
    result_file: '/etc/ecs_sd_targets.yaml'
    services:
      - name_pattern: '^.*nginx-service$' # NGINX
        metrics_ports:
          - 9113
        job_name: nginx-prometheus-exporter

receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: "ecssd"
          file_sd_configs:
            - files:
                - '/etc/ecs_sd_targets.yaml' # MUST match the file name in ecs_observer.result_file

processors:
  metricstransform:
    transforms:
      - include: ".*" # Rename customized job label back to job
        match_type: regexp
        action: update
        operations:
          - label: prometheus_job # must match the value configured in ecs_observer
            new_label: job
            action: update_label

```

### prom_metric_type Label

`prom_metric_type` is a label only used by CloudWatch builtin dashboards. In order to do that, we changed EMF exporter
to look up resource attributes
and [change output when receiver is prometheus](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/f02e8a03a15a64cd94f0cc5364dc67a9c58343fd/exporter/awsemfexporter/metric_translator.go#L146-L162)
. However, `recevier` is not a default attribute, and we insert it manually using `resource` processor. In another word,
our solution only works when prometheus receiver is the only metrics receiver sending metrics to CloudWatch EMF exporter
in the pipeline.

```yaml
processors:
  resource:
    attributes:
      - key: receiver # Insert receiver: prometheus for CloudWatch EMF Exporter to add prom_metric_type
        value: "prometheus"
        action: insert
```

## Future Work

### Cluster name auto detection

Unlike EKS, ECS has a reliable way to discover current cluster using endpoint provided by ECS agent. We didn't include
it in initial release because we already
have [two components](https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/3188)
with duplicated code for metadata client.

To implement this feature, just check metadata API if user give empty cluster name. Scraping metrics in cluster A using
collector running in cluster B is a valid use case, so we shouldn't override cluster name if user already provide one.
In fact, the collector can run anywhere as long as it can connect to AWS API and ECS tasks.

## Changelog

- 2021-06-23 @pingleig init the doc, ported
  from [#435](https://github.com/aws-observability/aws-otel-collector/pull/435)