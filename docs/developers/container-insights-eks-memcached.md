# Container Insights with Memcached

This example shows how to use aws-otel-collector to collect Prometheus metrics of memcached on EKS, and export them to
CloudWatch.

## Prerequisites

1. [Set Up EKS Cluster](setup-eks.md).

## Install Memcached 

1. Add Helm Repository.

```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
```

2. Install via Helm

```bash
# Create namespace to run memcached 
kubectl create namespace memcached-sample

# Install memcached with prometheus annotation
helm install my-memcached bitnami/memcached --namespace memcached-sample \
--set metrics.enabled=true \
--set-string serviceAnnotations.prometheus\\.io/port="9150" \
--set-string serviceAnnotations.prometheus\\.io/scrape="true"
```

## Install aws-otel-collector

Please follow [Install aws-otel-collector](container-insight-install-aoc.md) to complete the installation.

## View Your Data

1. View your Container Insights auto-dashboard

   Open the [CloudWatch console](https://console.aws.amazon.com/cloudwatch/). In the AWS Region where your cluster is
   running, choose `Container Insights - Performance monitoring` in the navigation pane. Then
   select `EKS Prometheus Memcached` and `{{cluster_name}}` from the drop down menu in the main pane. The dashboard shows
   the runtime status of memcached in the cluster.

2. View Your Metrics

   To see the CloudWatch metrics, choose Metrics in the navigation pane. The metric are in the
   `ContainerInsights/Prometheus` namespace.

3. View your logs

   To see the CloudWatch Logs events, choose Log groups in the navigation pane. The events are in the log group
   `/aws/containerinsights/{{cluster_name}}/prometheus` in the log stream `kubernetes-service-endpoints`

