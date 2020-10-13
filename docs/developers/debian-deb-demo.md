### Run AWSOTelCollector Beta on AWS EC2 Debian(ubuntu)

To run AWSOTelCollector on AWS EC2 debian host, you can choose to install AWSOTelCollector Debian on your host by the following steps.

**Steps,**

1. Login on AWS Debian host and download aws-otel-collector DEB with the following command.
```
wget https://aws-observability-collector-test.s3.amazonaws.com/debian/amd64/latest/aws-otel-collector.deb
```
2. Install aws-otel-collector DEB by the following command on the host
```
sudo dpkg -i -E ./aws-otel-collector.deb
```
3. Once DEB is installed, it will create AWSOTelCollector in directory /opt/aws/aws-otel-collector/

4. We provided a control script to manage AWSOTelCollector. Customer can use it to Start, Stop and Check Status of AWSOTelCollector.

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
5. Test the data with the running AWSOTelCollector on EC2. you can run the following command on EC2 host. (Docker app has to be pre-installed)
```
docker run --rm -it -e "otlp_endpoint=172.17.0.1:55680" -e "otlp_instance_id=test_insance" mxiamxia/aws-otel-metric-generator:latest
```
