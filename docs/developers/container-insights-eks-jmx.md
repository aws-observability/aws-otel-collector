# Container Insights with Java/JMX

This example shows how to use aws-otel-collector to collect Prometheus metrics of Java Virtual Machine (JVM), Java, and
Tomcat (Catalina) on EKS, and export them to CloudWatch.

## Prerequisites

1. [Set Up EKS Cluster](setup-eks.md).
2. Docker.

## Publish Sample Application Images

1. Download [Sample JMX traffic with Tomcat](https://github.com/aws-observability/aws-otel-test-framework/tree/terraform/sample-apps/jmx#sample-jmx-traffic-with-tomcat).

2. Authenticate to Amazon ECR

```bash
export AWS_ACCOUNT_ID={aws_account_id}
export AWS_REGION={{region}}
aws ecr get-login-password --region $AWS_REGION | docker login --username AWS --password-stdin $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com
```

3. Create a repository

```bash
aws ecr create-repository --repository-name prometheus-sample-tomcat-jmx \
  --image-scanning-configuration scanOnPush=true \
  --region $AWS_REGION 
```

4. Build Docker image and push to ECR.

```bash
docker build -t $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/prometheus-sample-tomcat-jmx:latest .
docker push $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/prometheus-sample-tomcat-jmx:latest 
```

## Install Sample Applications

```bash
SAMPLE_TRAFFIC_NAMESPACE=javajmx-sample
curl https://raw.githubusercontent.com/aws-observability/aws-otel-test-framework/terraform/sample-apps/jmx/examples/prometheus-metrics-sample.yaml | 
sed "s/{{aws_account_id}}/$AWS_ACCOUNT_ID/g" |
sed "s/{{region}}/$AWS_REGION/g" |
sed "s/{{namespace}}/$SAMPLE_TRAFFIC_NAMESPACE/g" | 
kubectl apply -f -
```

## Install ADOT Collector

Please follow [Install ADOT Collector](container-insight-install-aoc.md) to complete the installation.

## View Your Data

1. View your Container Insights auto-dashboard

   Open the [CloudWatch console](https://console.aws.amazon.com/cloudwatch/). In the AWS Region where your cluster is
   running, choose `Container Insights - Performance monitoring` in the navigation pane. Then
   select `EKS Prometheus Java/JMX` and `{{cluster_name}}` from the drop down menu in the main pane. The dashboard shows
   the runtime status of Java applications in the cluster.

2. View Your Metrics

   To see the CloudWatch metrics, choose Metrics in the navigation pane. The metric are in the
   `ContainerInsights/Prometheus` namespace.

3. View your logs

   To see the CloudWatch Logs events, choose Log groups in the navigation pane. The events are in the log group
   `/aws/containerinsights/{{cluster_name}}/prometheus` in the log stream `kubernetes-service-endpoints`

