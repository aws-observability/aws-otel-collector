### Using AWS-OTel-Collector on Amazon EKS

This example will introduce how to use AWS-OTel-Collector to send application traces and metrics on AWS EKS. This instruction provided the data emitter image that will generate OTLP format of metrics and traces data to AWS CloudWatch and X-Ray consoles.  Please follow the steps below to try ADOT Collector Beta.

### Create EKS-AWSOTel IAM Policy 
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

#### Deploy AWSOTelCollector on Amazon EKS as sidecar
The easiest way to deploy AWSOTelCollector on Amazon EKS is to run it as a sidecar, defining it in the same pod definition as your application.

1. Create a Kubernetes namespace.
```bash
kubectl create namespace aws-otel-eks
```

2. Setup AWS_REGION and CLUSTER_NAME for exporting metrics to CloudWatch
```
export CLUSTER_NAME=<eks-cluster-name>
export AWS_REGION=<aws-region>
```

3. Deploy the application. An example config template can be found [here](../../deployment-template/eks/otel-cloudwatch-sidecar.yaml).
* Replace `{{region}}` with the name of the region where the logs are published (e.g. `us-west-2`).
* Replace `{{cluster_name}}` with the actual eks cluster name.
```bash
cat otel-cloudwatch-sidecar.yaml |
sed "s/{{region}}/$AWS_REGION/g" | 
sed "s/{{cluster_name}}/$CLUSTER_NAME/g" |
kubectl apply -f - 
```
4. View the resources in the `aws-otel-eks` namespace.
```bash
kubectl get all -n aws-otel-eks
```
5. View the details of the deployed deployment.
```bash
kubectl -n aws-otel-eks describe deployment aws-otel-eks-sidecar
```

The example template provided runs the AWS-OTel-Collector as sidecar to send application metrics and traces on Amazon EKS. We run two applications: the customer’s application (`aws-otel-emitter`) and the AWSOTelCollector `aws-otel-collector`. Running the AWSOTelCollector in the same application as the main application allows the AWSOTelCollector to collect the metric/trace data for the customer’s application. We also call running the AWSOTelCollector in this way a "Sidecar". 

**View Your Metrics**  
You should now be able to view your metrics in your [CloudWatch console](https://console.aws.amazon.com/cloudwatch/). In the navigation bar, click on **Metrics**. The collected AWSOTelCollector metrics can be found in the **AWSObservability/CloudWatchEKSService** namespace. Ensure that your region is set to the region set for your cluster.
