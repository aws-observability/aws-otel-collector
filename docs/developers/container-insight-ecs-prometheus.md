# Container Insight ECS Prometheus

Table of Content

<!-- TODO(pingleig): generates the toc -->

## Overview

Container insight provides auto dashboard for common workload like [app mesh](#app-mesh), [jmx](#jmx), [nginx](#nginx),
[nginx plus](#nginx-plus).

## Usage

First deploy the collector following instructions for [ecs-prometheus](ecs-prometheus.md). Then deploy one
of [the workloads](#workloads).

## Known issues

[Nginx](#nginx) and [Nginx Plus](#nginx-plus) auto dashboard will show other non nginx services. You can create your own
dashboard to filter service name and ignore unrelated services.

## Workloads

The auto dashboard supports following four workloads out of box. NOTE: you
MUST [enable ECS container insight](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/deploy-container-insights-ECS-cluster.html)
. Otherwise, the dashboard is empty even if there are valid metrics.

### App Mesh

Setting up Nginx workload is same
as [CloudWatch Agent App Mesh](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/ContainerInsights-Prometheus-Sample-Workloads-ECS-appmesh.html)
. The example configuration is for
the [color apps](https://github.com/aws/aws-app-mesh-examples/tree/master/examples/apps/colorapp#app-mesh-walkthrough-deploy-the-color-app-on-ecs)
. If you want to try other app mesh samples you need to modify the discovery configuration, relabel, metrics selection
accordingly.

### JMX

Setting up Java JMX workload is same
as [CloudWatch Agent JMX](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/ContainerInsights-Prometheus-Sample-Workloads-ECS-javajmx.html)
. You can find the source for sample jmx image
in [our test framework](https://github.com/aws-observability/aws-otel-test-framework/tree/terraform/sample-apps/jmx).

An example cloudformation that deploys a replica service on ECS EC2
is [deployment-template/ecs/containerinsight-ecs-prometheus-workload-jmx-cfn](../../deployment-template/ecs/containerinsight-ecs-prometheus-workload-jmx-cfn.yaml)
NOTE: the roles are created when [deploying collector](ecs-prometheus.md#deployment).

```bash
export CLUSTER_NAME=aoc-containerinsight-prometheus-example
export WORKLOAD_IMAGE=we-dont-have-public-workload-image-yet
export TASK_ROLE_ARN=AWSOTelRolePrometheusECS
export TASK_EXECUTION_ROLE_ARN=AWSOTelExecutionRolePrometheusECS

aws cloudformation create-stack --stack-name AOC-Prometheus-ECS-${CLUSTER_NAME}-workload-jmx \
    --template-body file://containerinsight-ecs-prometheus-workload-jmx-cfn.yaml \
    --parameters ParameterKey=ClusterName,ParameterValue=${CLUSTER_NAME} \
                 ParameterKey=TaskRoleArn,ParameterValue=${TASK_ROLE_ARN} \
                 ParameterKey=ExecutionRoleArn,ParameterValue=${TASK_EXECUTION_ROLE_ARN} \
                 ParameterKey=WorkloadImage,ParameterValue=${WORKLOAD_IMAGE}
```

### Nginx

Setting up Nginx workload is same
as [ClouWatch Agent Nginx ECS](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/ContainerInsights-Prometheus-Setup-nginx-ecs.html)

### Nginx Plus

Setting up Nginx plus workload is same
as [CloudWatch Agent Nginx Plus ECS](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/ContainerInsights-Prometheus-Setup-nginx-plus-ecs.html)
. You need to register a 30 days nginx plus trial to build the docker image.