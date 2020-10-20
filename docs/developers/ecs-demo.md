### Using AWS-OTel-Collector on Amazon ECS

This example will introduce how to use AWS-OTel-Collector to send application traces and metrics on AWS ECS. This instruction provided the data emitter image that will generate OTLP format of metrics and traces data to AWS CloudWatch and X-Ray consoles.  Please follow the steps below to try AWS OTel Collector Beta.

* Install AWS-OTel-Collector via Task Definition Template
* Install AWS-OTel-Collector via CloudFormation Template

### Install AWSOTelCollector via Task Definition Template Steps
#### Create ECS-AWSOTel IAM Policy 
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
5. Choose Review policy.
6. On the Review policy page, enter `ECS-AWSOTel` for the Name and choose Create policy.

#### Create ECS-AWSOTel IAM Role
1. Open the IAM console at https://console.aws.amazon.com/iam/.
2. In the navigation pane, choose **Roles, Create role**.
3. In the Select type of trusted entity section, choose **AWS service**.
4. For Choose a use case, choose Elastic Container Service.
5. For Select your use case, choose Elastic Container Service Task, then choose Next: Tags.
5. Choose Next: Review.
6. On the Review page, enter `ECS-AWSOTel` for the Name and choose Create role.

#### Install AWSOTelCollector on Amazon ECS as sidecar
The easiest way to deploy AWSOTelCollector on Amazon ECS is to run it as a sidecar, defining it in the same task definition as your application.

**ECS Fargate Template**  
You can find ECS Fargate AWSOTelCollector [Installing template](../../examples/ecs/ecs-fargate-sidecar.json). Instructions for setting up a service in your ECS cluster using Fargate can be found [here](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/getting-started-fargate.html).

**ECS EC2 Template**  
Another example is installing AWSOTelCollector on ECS EC2 instance as sidecar to send telemetry data.
You can find ECS EC2 AWSOTelCollector [Installing template](../../examples/ecs/ecs-ec2-sidecar.json). Instructions for setting up a service in your ECS cluster using EC2 can be found [here](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/getting-started-ecs-ec2.html).

**Configure The Task Definition**  
The two ECS task definition templates are provided to run AWS-OTel-Collector as sidecar to send application metrics and traces on Amazon ECS. Notice that in the task definition templates we provided, you have to replace `{{region}}` with the region your want to send the data to. Fill `{{ecsTaskRoleArn}}` and `{{ecsExecutionRoleArn}}` with the IAM role (`ECS-AWSOTel`) you created above.

In the task definition, we run two applications: the customer’s application (`aws-otel-emitter`) and the AWSOTelCollector `aws-otel-collector`. Running the AWSOTelCollector in the same application as the main application allows the AWSOTelCollector to collect the metric/trace data for the customer’s application. We also call running the AWSOTelCollector in this way a "Sidecar".


**View Your Metrics**  
You should now be able to view your metrics in your [CloudWatch console](https://console.aws.amazon.com/cloudwatch/). In the navigation bar, click on **Metrics**. The collected AWSOTelCollector metrics can be found in the **AWSObservability/CloudWatchOTService** namespace. Ensure that your region is set to the region set for your cluster.

### Install AWSOTelCollector via CloudFormation Template Steps

#### Prerequisite
To Install AWSOTelCollector Via CloudFormation, you have to manually create an ECS cluster on your AWS console first.

#### Install AWSOTelCollector on ECS EC2
Download CloudFormation template file for installing AWSOTelCollector on ECS EC2 mode
```
curl -O https://github.com/mxiamxia/aws-opentelemetry-collector/blob/master/deployment-template/ecs/aws-otel-ec2-sidecar-deployment-cfn.template
```
Run CloudFormation the following command once ```Cluster_Name```, ```AWS_Region``` and  ```CFN_File_Downloaded``` are filled.
```
ClusterName=<Cluster_Name>
Region=<AWS_Region>
aws cloudformation create-stack --stack-name AWSOTelCollectorECS-${ClusterName}-${Region} \
    --template-body file://<CFN_File_Downloaded> \
    --parameters ParameterKey=ClusterName,ParameterValue=${ClusterName} \
                 ParameterKey=CreateIAMRoles,ParameterValue=false \
    --capabilities CAPABILITY_NAMED_IAM \
    --region ${Region}
```

#### Install AWSOTelCollector on ECS Fargate
Download CloudFormation template file for installing AWSOTelCollector on ECS Fargate mode
```
curl -O https://github.com/mxiamxia/aws-opentelemetry-collector/blob/master/deployment-template/ecs/aws-otel-fargate-sidecar-deployment-cfn.template
```
Run CloudFormation the following command once ```Cluster_Name```, ```AWS_Region``` and  ```CFN_File_Downloaded``` are filled.
```
ClusterName=<Cluster_Name>
Region=<AWS_Region>
aws cloudformation create-stack --stack-name AWSOTelCollectorECS-${ClusterName}-${Region} \
    --template-body file://<CFN_File_Downloaded> \
    --parameters ParameterKey=ClusterName,ParameterValue=${ClusterName} \
                 ParameterKey=CreateIAMRoles,ParameterValue=false \
    --capabilities CAPABILITY_NAMED_IAM \
    --region ${Region}
```




