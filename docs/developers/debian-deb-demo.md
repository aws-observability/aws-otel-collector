### Run AWSOTelCollector Beta on AWS EC2 Debian(ubuntu)

To run AWSOTelCollector on AWS EC2 debian host, you can choose to install AWSOTelCollector Debian on your host by the following steps.

**Steps,**

1. Login on AWS Debian EC2 host and download aws-otel-collector source code and build Deb file with the following command.
```
git clone https://github.com/aws-observability/aws-otel-collector.git  
.tools/packaging/debian/create_deb.sh
```
2. (Optional) Verify the package integrity.
```
wget https://aws-otel-collector.s3.amazonaws.com/aws-otel-collector.gpg
gpg --import aws-otel-collector.gpg
sudo apt install dpkg-sig
dpkg-sig --verify aws-otel-collector.deb
```
If the package is verified, you should see output similar to this:
```
GOODSIG _gpgorigin XXXXX...
```
However, if you see `BADSIG` or `UNKNOWNSIG`, you should stop and try download the deb package from the offical source again.

3. Install aws-otel-collector DEB by the following command on the host
```
sudo dpkg -i -E ./aws-otel-collector.deb
```
4. Once DEB is installed, it will create AWSOTelCollector in directory /opt/aws/aws-otel-collector/

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
docker run --rm -it -e "otlp_endpoint=172.17.0.1:55680" -e "otlp_instance_id=test_insance" mxiamxia/aws-otel-metric-generator:latest
```



#### enable debugging log

add a key value pair into `/opt/aws/aws-otel-collector/etc/.env` and restart collector

```
sudo echo "loggingLevel=DEBUG >> /opt/aws/aws-otel-collector/etc/.env"
sudo /opt/aws/aws-otel-collector/aws-otel-collector-ctl -a stop
sudo /opt/aws/aws-otel-collector/aws-otel-collector-ctl -a start
```
