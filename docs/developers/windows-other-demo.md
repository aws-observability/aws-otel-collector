### Run AWSOTelCollector on AWS Windows Ec2 Host

To run AWSOTelCollector on AWS windows ec2 host, you can choose to install AWSOTelCollector MSI on your host by the following steps.

**Steps,**
1. Login on AWS Windows EC2 host and download aws-otel-collector MSI with the following command.
```
wget https://aws-observability-collector-release.s3.amazonaws.com/windows/amd64/latest/aws-observability-collector.msi
or you can download by directly pasting the link https://aws-observability-collector-test.s3.amazonaws.com/windows/amd64/latest/aws-observability-collector.msi in windows browser 
```
2. Install aws-otel-collector MSI by running the following command on the host
```
msiexec /i aws-otel-collector.msi
or can be installed by double clicking the windows msi file.
```
`While Installing the AWSOTelCollector it will show a popup mentioning that the software is not from verified publisher, this is because we have not signed the MSI one it is signed this popup will be gone`

3. Once MSI is installed, it will create AWSOTelCollector in directory C:\Program Files\Amazon\AwsOpentelemetryCollector

4. We provided a control script to manage AWSOTelCollector. Customer can use it to Start, Stop and Check Status of AWSOTelCollector.
    * Start AWSOTelCollector with CTL script. The config.yaml is optional, if it is not provided the default [config](../../config.yaml) will be applied.
    ```
      & '.\Program Files\Amazon\AwsOpentelemetryCollector\aws-otel-collector-ctl.ps1' -a start 
    ```
    * Stop the running AWSOTelCollector when finish the testing.
    ```
      & '.\Program Files\Amazon\AwsOpentelemetryCollector\aws-otel-collector-ctl.ps1' -a stop 

    ```
    * Check the status of AWSOTelCollector
    ```
      & '.\Program Files\Amazon\AwsOpentelemetryCollector\aws-otel-collector-ctl.ps1' -a status 
    ```
