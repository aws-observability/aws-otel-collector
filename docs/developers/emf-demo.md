### Using AWS-Observability-Collector to send Embedded Metrics Format Logs

This example will introduce how to use AOC Beta to send EMF logs with metrics on AWS ECS/EKS. This instruction provided the data emitter image that will generate OTLP format of metrics and traces data to AWS CloudWatch and X-Ray consoles.  Please follow the steps below to try AWS Observability Collector Beta.

#### Steps
**Prerequisite**  
AOC docker image is not yet available for Beta release. You may have to follow [build docker](build-docker.md) instruction to build your own AOC docker image to run the example.

#### Create ECS-AWSObservability IAM Policy 
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
                "logs:CreateLogGroup",
                "logs:CreateLogStream",
                "logs:PutLogEvents",
                "logs:DescribeLogStreams",
                "xray:*",
                "s3:*"
            ],
            "Resource": [
                "arn:aws:logs:*:*:*"
            ]
        }
    ]
}
```
5. Choose Review policy.
6. On the Review policy page, enter `ECS-AWSObservability` for the Name and choose Create policy.

#### Create ECS-AWSObservability IAM Role
1. Open the IAM console at https://console.aws.amazon.com/iam/.
2. In the navigation pane, choose **Roles, Create role**.
3. In the Select type of trusted entity section, choose **AWS service**.
4. For Choose a use case, choose Elastic Container Service.
5. For Select your use case, choose Elastic Container Service Task, then choose Next: Tags.
5. Choose Next: Review.
6. On the Review page, enter `ECS-AWSObservability` for the Name and choose Create role.

#### Install AOC on Amazon ECS as sidecar
The easiest way to deploy AOC on Amazon ECS is to run it as a sidecar, defining it in the same task definition as your application.

**Configure The Task Definition**  
The following two examples are provided to run AWS-Observability-Collector as sidecar to send application metrics and traces on Amazon ECS. Notice that in the task definition templates we provided, you have to fill `{{aocImage}}` with the aoc docker image you created. Replace `{{region}}` with the region your want to send the data to. Fill `{{ecsTaskRoleArn}}` and `{{ecsExecutionRoleArn}}` with the IAM role (`ECS-AWSObservability`) you created above.

In the task definition, we run two applications: the customer’s application (`aoc-emitter`) and the AOC `aoc-collector`. Running the AOC in the same application as the main application allows the AOC to collect the metric/trace data for the customer’s application. We also call running the AOC in this way a "Sidecar".

**ECS Fargate**  
You can find ECS Fargate AOC [Installing template](https://github.com/mxiamxia/aws-opentelemetry-collector/blob/master/examples/ecs/ecs-fargate-sidecar.json). Instructions for setting up a service in your ECS cluster using Fargate can be found [here](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/getting-started-fargate.html).

**ECS EC2**  
Another example is installing AOC on ECS EC2 instance as sidecar to send telemetry data.
You can find ECS EC2 AOC [Installing template](https://github.com/mxiamxia/aws-opentelemetry-collector/blob/master/examples/ecs/ecs-ec2-sidecar.json). Instructions for setting up a service in your ECS cluster using EC2 can be found [here](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/getting-started-ecs-ec2.html).

**View Your Metrics**  
You should now be able to view your metrics in your [CloudWatch console](https://console.aws.amazon.com/cloudwatch/). In the navigation bar, click on **Metrics**. The collected AOC metrics can be found in the **AWSObservability/CloudWatchOTService** namespace. Ensure that your region is set to the region set for your cluster.

#### Install AOC on Amazon EKS as sidecar
[todo]
