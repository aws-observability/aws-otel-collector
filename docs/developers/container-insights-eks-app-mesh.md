# Container Insights with App Mesh

This example shows how to use aws-otel-collector to collect Prometheus metrics of App Mesh applications on EKS, and
export them to CloudWatch.

## Prerequisites

* [Set Up EKS Cluster](setup-eks.md)

## Install App Mesh and Sample Applications

1. [Install App Mesh Controller](https://github.com/aws/eks-charts/tree/master/stable/appmesh-controller#app-mesh-controller)
2. [Install App Mesh Sample Applications](https://github.com/aws/aws-app-mesh-examples/tree/master/walkthroughs/howto-k8s-http-headers) (
   Prerequisites #1 is NOT required in this example).

## Install aws-otel-collector

Please follow [Install aws-otel-collector](container-insight-install-aoc.md) to complete the installation.

## View Your Data

1. View your Container Insights auto-dashboard

   Open the [CloudWatch console](https://console.aws.amazon.com/cloudwatch/). In the AWS Region where your cluster is
   running, choose `Container Insights - Performance monitoring` in the navigation pane. Then
   select `EKS Prometheus AppMesh` and `{{cluster_name}}` from the drop down menu in the main pane. The dashboard shows
   the runtime status of App Mesh applications in the cluster.

2. View Your Metrics

   To see the CloudWatch metrics, choose Metrics in the navigation pane. The metric are in the
   `ContainerInsights/Prometheus` namespace.

3. View your logs

   To see the CloudWatch Logs events, choose Log groups in the navigation pane. The events are in the log group
   `/aws/containerinsights/{{cluster_name}}/prometheus` in the log stream `kubernetes-pod-appmesh-envoy`.

## Known Issues

The configuration of `Summary by pods` pane is not compatible with the
current [CloudWatch embedded metric format](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Embedded_Metric_Format_Specification.html)
, thus pod statistics would not be displayed on the dashboard using aws-otel-collector.

Currently, the filters are set to match CloudWatchMetrics as top level member:

```bash
filter CloudWatchMetrics.0.Namespace=\"ContainerInsights/Prometheus\" |
filter CloudWatchMetrics.0.Metrics.0.Name like /envoy_/
```

In order to fix the format compatibility issue, these two filters should be updated to:

```bash
filter _aws.CloudWatchMetrics.0.Namespace=\"ContainerInsights/Prometheus\" |
filter _aws.CloudWatchMetrics.0.Metrics.0.Name like /envoy_/
```
