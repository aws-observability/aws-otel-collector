# Container Insights with Nginx

This example shows how to use aws-otel-collector to collect Prometheus metrics of App Mesh applications on EKS, and
export them to CloudWatch.

## Prerequisites

1. [Set Up EKS Cluster](setup-eks.md).

## Install Nginx

1. Add Helm Repository.

```bash
helm repo add nginx-stable https://helm.nginx.com/stable
helm repo update
```

2. Install via Helm

```bash
# Create namespace to run nginx
kubectl create namespace nginx-ingress-sample

# Install nginx with prometheus annotation
helm install my-nginx ingress-nginx/ingress-nginx \
--namespace nginx-ingress-sample \
--set controller.metrics.enabled=true \
--set-string controller.metrics.service.annotations."prometheus\.io/port"="10254" \
--set-string controller.metrics.service.annotations."prometheus\.io/scrape"="true"
```

3. Get service external IP

   It might take some time for the service external IP to be assigned. When it's ready, the value can be extracted by
   the following command.

```bash
kubectl get service my-nginx-ingress-nginx-controller -nnginx-ingress-sample | awk {'print $4'}
```

## Install Nginx Sample Applications
1. Set `EXTERNAL_IP` variable
```bash
export EXTERNAL_IP=$(kubectl get service my-nginx-ingress-nginx-controller -nnginx-ingress-sample --no-headers | awk {'print $4'})
```
2. Create sample applications to generate Nginx traffic
```bash
SAMPLE_TRAFFIC_NAMESPACE=nginx-sample-traffic
curl https://raw.githubusercontent.com/aws-samples/amazon-cloudwatch-container-insights/master/k8s-deployment-manifest-templates/deployment-mode/service/cwagent-prometheus/sample_traffic/nginx-traffic/nginx-traffic-sample.yaml | 
sed "s/{{external_ip}}/$EXTERNAL_IP/g" | 
sed "s/{{namespace}}/$SAMPLE_TRAFFIC_NAMESPACE/g" | 
kubectl apply -f -
```

## Install aws-otel-collector
Please follow [Install aws-otel-collector](container-insight-install-aoc.md) to complete the installation.

## View Your Data
1. View your Container Insights auto-dashboard

   Open the [CloudWatch console](https://console.aws.amazon.com/cloudwatch/). In the AWS Region where your cluster is
   running, choose `Container Insights - Performance monitoring` in the navigation pane. Then
   select `EKS Prometheus Nginx` and `{{cluster_name}}` from the drop down menu in the main pane. The dashboard shows
   the runtime status of Nginx applications in the cluster.

2. View Your Metrics

   To see the CloudWatch metrics, choose Metrics in the navigation pane. The metric are in the
   `ContainerInsights/Prometheus` namespace.

3. View your logs

   To see the CloudWatch Logs events, choose Log groups in the navigation pane. The events are in the log group
   `/aws/containerinsights/{{cluster_name}}/prometheus` in the log stream `kubernetes-service-endpoints`

