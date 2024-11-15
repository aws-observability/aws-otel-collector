### Run ADOT Collector on Amazon Linux 2
See [AWS Distro for OpenTelemetry documentation](https://aws-otel.github.io/docs/setup/build-collector-as-rpm) to run ADOT Collector on Amazon Linux 2 EC2 host.  

### Installing ADOT Collector via CloudFormation
See [AWS Distro for OpenTelemetry documentation](https://aws-otel.github.io/docs/setup/ec2) for detailed information on installing ADOT Collector via CloudFormation on EC2.

### Install ADOT Collector on ECS EC2
See [AWS Distro for OpenTelemetry documentation](https://aws-otel.github.io/docs/setup/ecs) for detailed information on installing ADOT Collector on ECS.

#### Enable debugging log

add a key value pair into `/opt/aws/aws-otel-collector/etc/extracfg.txt` and restart collector

```
echo "loggingLevel=DEBUG" | sudo tee -a /opt/aws/aws-otel-collector/etc/extracfg.txt
sudo /opt/aws/aws-otel-collector/bin/aws-otel-collector-ctl -a stop
sudo /opt/aws/aws-otel-collector/bin/aws-otel-collector-ctl -a start
```
