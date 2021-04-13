# Collect Prometheus metrics from Amazon ECS tasks

## Overview

Prometheus pull metrics from targets, it has builtin discovery for eks. AWS-OTel-Collector
uses [ecsobserver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/observer/ecsobserver)
to discover targets running on ECS.

## Known issues

See [ecsobserver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/observer/ecsobserver#known-issues)

- There can only be one collector for each cluster. If your ECS cluster is too big, you can use `hashmod` or divide
  discovery targets by service name etc. to shard the workload to multiple collector services, each service has slightly
  different config and only 1 instance.
- `ecsobeserver` does NOT use public ip as scrape target, when exporting targets, it always uses private IP within the
  VPC regardless of launch type and network mode. However, it does expose ec2 public ip as
  label `__meta_ecs_ec2_public_ip` which can be used in relabel config and override `__address__` if you have a valid
  use case and have extra authentication for making the endpoint public.

## Permission

### IAM Policy

Following extra IAM policies are required for ECS prometheus service discovery to work. Existing policy can be found
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

### Network

Make sure your ec2 tasks can reach http endpoints of other tasks using private ip. Depends on how the cluster is
created, you may need to configure security group to allow ingress within current vpc as prometheus **pull** metrics.
The default VPC normally allows ingress within VPC out of box, while some cli tools creates new VPC with fewer rules,
e.g. [ecs-cli](#create-ecs-cluster-with-ec2-instances).

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

## Run sample workload

You can find workloads from [container insight](container-insight-ecs-prometheus.md). You need to make sure the
service/task definition matches the config in example. [JMX](container-insight-ecs-prometheus.md#jmx) is the simplest
workload.

## Appendix

### Create ECS Cluster with EC2 instances

- Download [ecs-cli](https://github.com/aws/amazon-ecs-cli#installing).
- Create new cluster, take note of the security group and subnet it created.
    - if you missed it, you can still find it from cloudformation console, ecs-cli generates cfn and applies it.
- Update security group to allow ingress on all tcp port within the cluster, this makes testing easier, but you might
  want to only allow specific port when using are using non bridge network, bridge network assign random host ports.
- Delete the cluster by either delete the cfn stack or using `ecs-cli down -f -cluster ${CLUSTER_NAME}`

```bash
# Create ECS EC2 cluster with 2 EC2 instances

export CLUSTER_NAME=aoc-containerinsight-prometheus-example
ecs-cli up --capability-iam --size 2 --instance-type t2.medium --cluster ${CLUSTER_NAME} --region us-west-2

# Output is like
# INFO[0000] Using recommended Amazon Linux 2 AMI with ECS Agent 1.51.0 and Docker version 19.03.13-ce
# INFO[0001] Created cluster                               cluster=aoc-containerinsight-prometheus-example region=us-west-2
# INFO[0001] Waiting for your cluster resources to be created...
# INFO[0122] Cloudformation stack status                   stackStatus=CREATE_IN_PROGRESS
# VPC created: vpc-0fcxxxxx
# Security Group created: sg-04xxxxx
# Subnet created: subnet-03xxxx
# Subnet created: subnet-0exxxx
# Cluster creation succeeded.

# Update security group to allow ingress on all ports within security group
SG=sg-04xxxxx
aws ec2 authorize-security-group-ingress --group-id ${SG} --protocol tcp --port 0-65535 --source-group ${SG}
aws ec2 describe-security-groups --group-ids ${SG}
```