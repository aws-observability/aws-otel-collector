### Run AWSOTelCollector Beta Examples with Docker

This example will introduce how to run AWSOTelCollector Beta in the Docker container. This example uses a AWS data emitter container image that will generate Open Telemetry Protocol (OTLP) format based metrics and traces data to AWS CloudWatch and X-Ray consoles.  

Please follow the steps below to try AWS OTel Collector Beta.

#### Prerequisite

If you haven't setup your AWS Credential profile yet, please follow the [instruction](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html) for setting up your AWS credentials.

#### Run a single Aws-OTel-Collector instance in Docker
You can bring up a single `aws-otel-collector` instance in Docker.
```
git clone https://github.com/aws-observability/aws-otel-collector.git ; \
    cd aws-otel-collector; \
    docker run --rm -p 4317:4317 -p 55680:55680 -p 8889:8888 \
      -e AWS_REGION=us-west-2 \
      -v "${PWD}/examples/config-test.yaml":/otel-local-config.yaml \
      --name awscollector public.ecr.aws/aws-observability/aws-otel-collector:latest \
      --config otel-local-config.yaml;
```

#### Run Aws-OTel-Collector with Sample App in Docker Compose

1. Checkout `aws-otel-collector` source code, and open the ```docker-compose.yaml``` under ```examples``` folder.
Please make sure you have the right aws credential path (eg, `~/.aws:/root/.aws`) and the collector config file (eg, `../config.yaml:/etc/otel-agent-config.yaml`) set.
 You can also directly use your AWS credential key by setting up these environment variables ```AWS_ACCESS_KEY_ID```, ```AWS_SECRET_ACCESS_KEY``` and ```AWS_REGION``` in the config.
  The region is where the data will be sent to.
```# Agent aws-otel-collector:
    image: public.ecr.aws/aws-observability/aws-otel-collector:latest
    command: ["--config=/etc/otel-agent-config.yaml", "--log-level=DEBUG"]
    environment:
      - AWS_ACCESS_KEY_ID=<to_be_added>
      - AWS_SECRET_ACCESS_KEY=<to_be_added>
      - AWS_REGION=<to_be_added>
    volumes:
      - ../config.yaml:/etc/otel-agent-config.yaml // use default config
      - ~/.aws:/root/.aws
```
2. Once you have the ```docker-compose.yaml``` file setup and saved, run the following make command.
```
cd examples; docker-compose up 
```
3. Now you can view you data in AWS console

    * X-Ray - aws console
    * CloudWatch - aws console  
    
**AWS Metrics Sample Data**   
* ![aws metrics](../images/metrics_sample.png)  
**AWS Traces Sample Data**
* ![aws traces](../images/traces_sample.png)  

4. Stop the running AWSOTelCollector in Docker container
```
make docker-stop
```
