## Deploy AWS OTel Collector on Amazon EKS with Sample Workloads

The tutorial shows you how to set up AWS OTel Collector on Amazon EKS
with [sample workloads]((https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/ContainerInsights-Prometheus-Sample-Workloads.html))
provided by Amazon CloudWatch

### Set Up IAM Policies

In order to allow AWS OTel Collector to send logs to CloudWatch, please ensure that the following policies exist on EKS node role.

#### Create EKS-AWSOTel IAM Policy
1. Open the IAM console at https://console.aws.amazon.com/iam/.
2. In the navigation pane, choose **Policies**.
3. Choose **Create policy, JSON**.
4. Enter the following policy:
```json
{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Effect": "Allow",
			"Action": [
				"logs:PutLogEvents",
				"logs:CreateLogGroup",
				"logs:CreateLogStream",
				"logs:DescribeLogStreams",
				"logs:DescribeLogGroups",
				"xray:PutTraceSegments",
				"xray:PutTelemetryRecords",
				"xray:GetSamplingRules",
				"xray:GetSamplingTargets",
				"xray:GetSamplingStatisticSummaries",
				"ssm:GetParameters"
			],
			"Resource": "*"
		}
	]
}
```
5. Choose **Review policy**.
6. On the Review policy page, enter `EKS-AWSOTel` for the Name and choose **Create policy**.

#### Attach EKS-AWSOTel IAM Role to worker nodes
1. Open the Amazon EC2 console at https://console.aws.amazon.com/ec2/.
2. Select one of the worker node instances and choose the IAM role in the description.
3. On the IAM role page, choose **Attach policies**.
4. In the list of policies, select the check box next to `EKS-AWSOTel`. If necessary, use the search box to find this policy.
5. Choose **Attach policies**.

### Install AWS OTel Collector
1. Set up variables to export metrics of your EKS cluster to the region where the logs should be published to.
```
export CLUSTER_NAME=<eks-cluster-name>
export AWS_REGION=<aws-region>
```
2. Deploy a standalone AWS OTel collector application. An example config template can be found [here](../../deployment-template/eks/standalone-otel-eks-deployment.yaml).
   * Replace `{{region}}` with the name of the region where the logs are published (e.g. `us-west-2`).
   * Replace `{{cluster_name}}` with the actual eks cluster name.
```bash
cat standalone-otel-eks-deployment.yaml |
sed "s/{{region}}/$AWS_REGION/g" | 
sed "s/{{cluster_name}}/$CLUSTER_NAME/g" |
kubectl apply -f - 
```
3. View the resources in the `aws-otel-eks` namespace.
```bash
kubectl get all -n aws-otel-eks
```
4. View Your Metrics  
You should now be able to view your metrics in your [CloudWatch console](https://console.aws.amazon.com/cloudwatch/). In
the navigation bar, click on **Metrics**. The collected AWSOTelCollector metrics can be found in the **
AWSObservability/CloudWatchOTService** namespace. Ensure that your region is set to the region set for your cluster.

