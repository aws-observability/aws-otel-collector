# Collect Prometheus metrics from Amazon ECS tasks

## Overview

Prometheus pull metrics from targets, it has builtin discovery for eks. AWS-OTel-Collector
uses [ecsobserver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/observer/ecsobserver)
to discover targets running on ECS.

## Permission

Following extra policies are required for ECS prometheus service discovery to work. Existing policy can be found
when [setting up IAM role policy](ecs-demo.md#create-ecs-awsotel-iam-policy).

NOTE: `ec2` policy is required for listing ec2 instances to get private IP, it is optional if all your tasks are fargate
or uses `awsvpc` network mode.

```text
ec2:DescribeInstances
ecs:ListTasks
ecs:ListServices
ecs:DescribeContainerInstances
ecs:DescribeServices
ecs:DescribeTasks
ecs:DescribeTaskDefinition
```

## Configuration

Full collector configuration can be found
in [examples/ecs/ecs-containerinsight-prometheus.yaml](../../examples/ecs/ecs-containerinsight-prometheus.yaml). It
contains discovery and metrics extraction rule for all
the [container insight workloads](container-insight-ecs-prometheus.md)
.

NOTE: You must replace `{{cluster_name}}` in the example, it cannot discover ecs cluster automatically (for now).

Detail explanation for observer config is in
its [source](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/observer/ecsobserver)
.

## Deployment

You need to deploy the collector as a replica service (count is 1) and inject configuration
using [ssm parameter](https://aws-otel.github.io/docs/setup/ecs/config-through-ssm).

NOTE: There can only be one instance for each cluster. If your ECS cluster is too big, you can use `hashmod` or divide
discovery targets by service name etc.

Example task definition for running collector on ECS EC2. You need to replace all the `{{VAR}}` with correct value as
in [ecs demo](ecs-demo.md).

```json
{
  "family": "aoc-ecs-sd",
  "taskRoleArn": "{{TaskRole}}",
  "executionRoleArn": "{{TaskExecutionRole}}",
  "networkMode": "bridge",
  "containerDefinitions": [
    {
      "name": "aoc",
      "image": "{{IMAGE}}",
      "secrets": [
        {
          "name": "AOT_CONFIG_CONTENT",
          "valueFrom": "{{SSM}}
        }
      ]
    }
  ],
  "requiresCompatibilities": [
    "EC2"
  ],
  "cpu": "128",
  "memory": "64"
}
```