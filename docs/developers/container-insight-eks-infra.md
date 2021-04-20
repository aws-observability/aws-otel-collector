# Overview

When deployed as a daemon set in an EKS cluster, the AWS OpenTelemetry Collector can be used to collect infrastructure metrics for many resources in the cluster, such as CPU, memory, disk, and network. This document describes how to install aws-otel-collector into your existing EKS cluster to collect infrastructure metrics from your cluster. 

## Prerequisites

* [See AWS public doc to set up EKS cluster](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Container-Insights-prerequisites.html)

## Install aws-otel-collector

1. Set up variables to export metrics of your EKS cluster to the region where the logs should be published to.

```
export CLUSTER_NAME=<eks-cluster-name>
export AWS_REGION=<aws-region>
```

2. Deploy aws-otel-collector as a daemon set. An example config template can be found [here](../../deployment-template/eks/otel-container-insight-infra.yaml).
    * Replace `{{region}}` with the name of the region where the logs are published (e.g. `us-west-2`).
    * Replace `{{cluster_name}}` with the actual eks cluster name.

```bash
cat otel-container-insight-infra.yaml |
sed "s/{{region}}/$AWS_REGION/g" | 
sed "s/{{cluster_name}}/$CLUSTER_NAME/g" |
kubectl apply -f - 
```

3. Verify the collector is running

```bash
kubectl get pods -l name=aws-otel-eks-ci -n aws-otel-eks
```

## View Your Data
1. View CloudWatch auto-dashboards for Container Insights

   Open the [CloudWatch console](https://console.aws.amazon.com/cloudwatch/). In the AWS Region where your cluster is
   running, choose `Container Insights - Performance monitoring` in the navigation pane. Use the drop-down boxes near the top to select the type of resource to view, as well as the specific cluster name. The following resource types are available: `EKS Clusters`, `EKS Namespaces`, `EKS Nodes`, `EKS Services`, and `EKS Pods`. 

2. View Your Metrics

   To see the CloudWatch metrics, choose Metrics in the navigation pane. The metric are in the
   `ContainerInsights` namespace.

3. View your logs

   To see the CloudWatch Logs events, choose Log groups in the navigation pane. The events are in the log group
   `/aws/containerinsights/{{cluster_name}}/performance`. 
