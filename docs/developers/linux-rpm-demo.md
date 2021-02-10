### Run AWSOTelCollector Beta on AWS EC2 Linux

To run AWSOTelCollector on AWS EC2 Linux host, you can choose to install AWSOTelCollector RPM on your host by the following steps.

**Steps,**

1. Login on AWS Linux EC2 host and download aws-otel-collector installation file.
```
wget https://aws-otel-collector.s3.amazonaws.com/amazon_linux/amd64/latest/aws-otel-collector.rpm
```
2. (Optional) Verify the package integrity.
```
wget https://aws-otel-collector.s3.amazonaws.com/aws-otel-collector.gpg
sudo rpm --import aws-otel-collector.gpg
rpm --checksig aws-otel-collector.rpm
```
If the package is verified correctly, you'll see output similar to:
```
aws-otel-collector.rpm: digests signatures OK
```
However, if you see output like:
```
aws-otel-collector.rpm: digests SIGNATURES NOT OK
```
You should stop, and try download the rpm package from the offical source again.

3. Install aws-otel-collector RPM by the following command on the host
```
sudo rpm -Uvh  ./aws-otel-collector.rpm
```
4. Once RPM is installed, it will create AWSOTelCollector in directory ```/opt/aws/aws-otel-collector/```

5. We provided a control script to manage AWSOTelCollector. Customer can use it to Start, Stop and Check Status of AWSOTelCollector.

    * Start AWSOTelCollector with CTL script. The config.yaml is optional, if it is not provided the default [config](../../config.yaml) will be applied.  
    ```
        sudo /opt/aws/aws-otel-collector/bin/aws-otel-collector-ctl -c </path/config.yaml> -a start
    ```
    * Stop the running AWSOTelCollector when finish the testing.
    ```
        sudo /opt/aws/aws-otel-collector/bin/aws-otel-collector-ctl  -a stop
    ```
    * Check the status of AWSOTelCollector
    ```
        sudo /opt/aws/aws-otel-collector/bin/aws-otel-collector-ctl  -a status
    ```
      
6. Test the data with the running AWSOTelCollector on EC2. you can run the following command on EC2 host. (Docker app has to be pre-installed)
```
docker run --rm -it --network host -e "OTEL_EXPORTER_OTLP_ENDPOINT=127.0.0.1:4317" -e "otlp_instance_id=test_insance_rpm" -e "OTEL_RESOURCE_ATTRIBUTES=service.namespace=AWSOTelCollectorRPMDemo,service.name=AWSOTelCollectorRPMDemoService" -e S3_REGION=us-west-2 aottestbed/aws-otel-collector-sample-app:java-0.1.0
```

**View Your Metrics**  
You should now be able to view your metrics in your [CloudWatch console](https://console.aws.amazon.com/cloudwatch/). In the navigation bar, click on **Metrics**. The collected AWSOTelCollector metrics can be found in the **AWSOTelCollectorRPMDemo/AWSOTelCollectorRPMDemoService** namespace.

#### Installing AWSOTelCollector via CloudFormation
#### Install AWSOTelCollector on ECS EC2
Download CloudFormation template file for installing AWSOTelCollector on ECS EC2 mode
```
curl -O https://raw.githubusercontent.com/aws-observability/aws-otel-collector/main/deployment-template/ec2/aws-otel-ec2-deployment-cfn.yaml
```
Run CloudFormation the following command once ```IAMRole```, ```Region```, ```EC2Key``` and  ```CFN_File_Downloaded``` are filled.
```
Region=<aws-region>
IAMRole=<iam-role>
EC2Key=<ec2-ssh-key-name>
aws cloudformation create-stack --stack-name AWSOTelCollectorEC2-Test \
	--template-body file://<CFN_File_Downloaded> \
	--parameters ParameterKey=IAMRole,ParameterValue=${IAMRole} \
                 ParameterKey=KeyName,ParameterValue=${EC2Key} \
    --capabilities CAPABILITY_NAMED_IAM \
    --region ${Region}
```

#### enable debugging log

add a key value pair into `/opt/aws/aws-otel-collector/etc/extracfg.txt` and restart collector

```
sudo echo "loggingLevel=DEBUG >> /opt/aws/aws-otel-collector/etc/extracfg.txt"
sudo /opt/aws/aws-otel-collector/aws-otel-collector-ctl -a stop
sudo /opt/aws/aws-otel-collector/aws-otel-collector-ctl -a start
```
