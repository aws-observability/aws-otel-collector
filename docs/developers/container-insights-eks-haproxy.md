# Container Insights with HAProxy

This example shows how to use aws-otel-collector to collect Prometheus metrics of HAProxy on EKS, and export them to
CloudWatch.

## Prerequisites

1. [Set Up EKS Cluster](setup-eks.md).

## Install HAProxy

1. Add Helm Repository.

```bash
helm repo add haproxy-ingress https://haproxy-ingress.github.io/charts
helm repo update
```

2. Install via Helm

```bash
# Create namespace to run HAProxy 
kubectl create namespace haproxy-ingress-sample

# Install HAProxy with prometheus annotation
helm install haproxy haproxy-ingress/haproxy-ingress \
--namespace haproxy-ingress-sample \
--set defaultBackend.enabled=true \
--set controller.stats.enabled=true \
--set controller.metrics.enabled=true \
--set-string controller.metrics.service.annotations."prometheus\.io/port"="9101" \
--set-string controller.metrics.service.annotations."prometheus\.io/scrape"="true"
```

**Important**

Check out [Supoorted versions](https://github.com/jcmoraisjr/haproxy-ingress/#use-haproxy-ingress) to install a
compatible version of HAProxy to your EKS Cluster.

## Install aws-otel-collector

Please follow [Install aws-otel-collector](container-insight-install-aoc.md) to complete the installation.

## View Your Data

1. View your Container Insights auto-dashboard

   Open the [CloudWatch console](https://console.aws.amazon.com/cloudwatch/). In the AWS Region where your cluster is
   running, choose `Container Insights - Performance monitoring` in the navigation pane. Then
   select `EKS Prometheus HAProxy` and `{{cluster_name}}` from the drop down menu in the main pane. The dashboard shows
   the runtime status of HAProxy in the cluster.

2. View Your Metrics

   To see the CloudWatch metrics, choose Metrics in the navigation pane. The metric are in the
   `ContainerInsights/Prometheus` namespace.

3. View your logs

   To see the CloudWatch Logs events, choose Log groups in the navigation pane. The events are in the log group
   `/aws/containerinsights/{{cluster_name}}/prometheus` in the log stream `kubernetes-service-endpoints`

