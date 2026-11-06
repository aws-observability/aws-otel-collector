### Run ADOTCollector Beta on Debian and Ubuntu

To run ADOTCollector on Debian (or Ubuntu) EC2 host, you can choose to install ADOTCollector on your host by the following steps.

**Steps,**

1. Login on Debian EC2 host and download aws-otel-collector installation file.
```
wget https://aws-otel-collector.s3.amazonaws.com/ubuntu/amd64/latest/aws-otel-collector.deb
```
2. (Optional) Verify the package integrity.

`dpkg-sig` is not available in Debian 12 or Ubuntu 24.04 repositories. The package contains a signed manifest with SHA-1 checksums. Verify the manifest with the published signing key, then check the contents against those checksums:

```sh
sudo apt install binutils gnupg gpgv
wget https://aws-otel-collector.s3.amazonaws.com/aws-otel-collector.gpg
package=$(realpath aws-otel-collector.deb)
key=$(realpath aws-otel-collector.gpg)
(
    set -e
    verify_dir=$(mktemp -d)
    trap 'rm -rf "$verify_dir"' EXIT
    gpg --dearmor < "$key" > "$verify_dir/keyring.gpg"
    cd "$verify_dir"
    test "$(ar t "$package" | wc -l)" -eq 4
    ar x "$package" debian-binary control.tar.gz data.tar.gz _gpgorigin
    gpgv --keyring ./keyring.gpg --output manifest _gpgorigin
    awk 'NF == 4 && $4 ~ /^(debian-binary|control.tar.gz|data.tar.gz)$/ {print $2 "  " $4}' manifest > checksums
    test "$(wc -l < checksums)" -eq 3
    sha1sum --check --strict checksums
)
```

Successful verification prints `Good signature` from `AWS OTel Collector` and all three checksums must pass:

```text
debian-binary: OK
control.tar.gz: OK
data.tar.gz: OK
```

If any command fails, do not install the package. Download it again from the official source and repeat verification. A good signature on the manifest alone does not verify the package contents.

3. Install aws-otel-collector DEB by the following command on the host
```
sudo dpkg -i -E ./aws-otel-collector.deb
```
4. Once DEB is installed, it will create ADOTCollector in directory /opt/aws/aws-otel-collector/

5. We provided a control script to manage ADOTCollector. Customer can use it to Start, Stop and Check Status of ADOTCollector.

    * Start ADOTCollector with CTL script. The config.yaml is optional, if it is not provided the default [config](../../config.yaml) will be applied.  
    ```
        sudo /opt/aws/aws-otel-collector/bin/aws-otel-collector-ctl -c </path/config.yaml> -a start
    ```
    * Stop the running ADOTCollector when finish the testing.
    ```
        sudo /opt/aws/aws-otel-collector/bin/aws-otel-collector-ctl  -a stop
    ```
    * Check the status of ADOTCollector
    ```
        sudo /opt/aws/aws-otel-collector/bin/aws-otel-collector-ctl  -a status
    ```
6. Test the data with the running ADOTCollector on EC2. you can run the following command on EC2 host. (Docker app has to be pre-installed)
```
docker run --rm -it --network host -e "OTEL_EXPORTER_OTLP_ENDPOINT=127.0.0.1:4317" -e "otlp_instance_id=test_insance_rpm" -e "OTEL_RESOURCE_ATTRIBUTES=service.namespace=ADOTCollectorRPMDemo,service.name=ADOTCollectorRPMDemoService" -e S3_REGION=us-west-2 aottestbed/aws-otel-collector-sample-app:java-0.1.0
```



#### enable debugging log

add a key value pair into `/opt/aws/aws-otel-collector/etc/extracfg.txt` and restart collector

```
echo "loggingLevel=DEBUG" | sudo tee -a /opt/aws/aws-otel-collector/etc/extracfg.txt
sudo /opt/aws/aws-otel-collector/bin/aws-otel-collector-ctl -a stop
sudo /opt/aws/aws-otel-collector/bin/aws-otel-collector-ctl -a start
```
