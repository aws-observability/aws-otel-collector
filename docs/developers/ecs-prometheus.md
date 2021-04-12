# Collect Prometheus metrics from Amazon ECS tasks

## Overview

Prometheus pull metrics from targets, it has builtin discovery for eks. AWS-OTel-Collector
uses [ecsobserver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/observer/ecsobserver)
to discover targets running on ECS.

## Known issues

See [ecsobserver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/observer/ecsobserver#known-issues)

- There can only be one collector for each cluster. If your ECS cluster is too big, you can use `hashmod` or divide
  discovery targets by service name to shard the workload.

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

You need to configure both `ecsobserver` extension and `prometheusreceiver` for ECS prometheus to work properly. The
config for `ecsobserver` tells it to call ECS API and filter out tasks that expose prometheus metrics, it then generates
a file that `prometheusreceiver` watch for change. The generated files contains private ip and ports for scrape targets.

Example collector configuration can be found
in [examples/ecs/ecs-containerinsight-prometheus.yaml](../../examples/ecs/ecs-containerinsight-prometheus.yaml). It
contains discovery and metrics extraction rule for all
the [container insight workloads](container-insight-ecs-prometheus.md)
. NOTE: You must replace `{{cluster_name}}` in the example, it cannot discover ecs cluster automatically (for now).

Detail explanation for observer config is in
[ecsobserver source](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/observer/ecsobserver)
.

## Deployment

When deploying collector prometheus, you need the following resources

- One ECS cluster, you can create one using [ecs-cli](#create-ecs-cluster-with-ec2-instances)
- IAM roles for tasks
- Collector configuration in [SSM parameter](https://aws-otel.github.io/docs/setup/ecs/config-through-ssm)
- Image url for collector

An example cloudformation that deploys a replica service (count is and [must be 1](#known-issues))
is [deployment-template/ecs/containerinsight-ecs-prometheus-task-cfn](../../deployment-template/ecs/containerinsight-ecs-prometheus-task-cfn.yaml)
.

EC2 bridge network

```bash
export CLUSTER_NAME=aoc-containerinsight-prometheus-example
export COLLECTOR_IMAGE=amazon/aws-otel-collector:latest
export CREATE_IAM_ROLES=True

aws cloudformation create-stack --stack-name AOC-Prometheus-ECS-${CLUSTER_NAME} \
    --template-body file://containerinsight-ecs-prometheus-task-cfn.yaml \
    --parameters ParameterKey=ClusterName,ParameterValue=${CLUSTER_NAME} \
                 ParameterKey=CreateIAMRoles,ParameterValue=${CREATE_IAM_ROLES} \
                 ParameterKey=CollectorImage,ParameterValue=${COLLECTOR_IMAGE} \
    --capabilities CAPABILITY_NAMED_IAM
```

Fargate awsvpc (TODO(pingleig): does not work the cfn template for now).

## Appendix

### Create ECS Cluster with EC2 instances

Download [ecs-cli](https://github.com/aws/amazon-ecs-cli#installing).

```bash
# Create ECS EC2 cluster with 2 EC2 instances

export CLUSTER_NAME=aoc-containerinsight-prometheus-example
ecs-cli up --capability-iam --size 2 --instance-type t2.medium --cluster ${CLUSTER_NAME} --region us-west-2
```