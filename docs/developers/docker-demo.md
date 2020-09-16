### Run AOC Beta Examples with Docker

This example will introduce how to run AOC Beta in the Docker container. This demo also includes AWS data emitter container image that will generate OTLP format of metrics and traces data to AWS CloudWatch and X-Ray consoles.  Please follow the steps below to try AWS Observability Collector Beta.

#### Steps,

**Prerequisite**  
AOC docker image is not yet available for Beta release. You may have to follow [build docker](build-docker.md) instruction to build your own AOC docker image to run the example.

1. Download the source code of this repo and edit the following section in ```docker-composite.yaml``` under ```examples``` folder. You will need to add your own ```aoc-image```, ```AWS_ACCESS_KEY_ID```, ```AWS_SECRET_ACCESS_KEY``` and ```AWS_REGION``` in the config. The region is where the data will be sent to.
```# Agent aws-observability-collector:
    image: <aoc>/<aoc-image>
    command: ["--config=/etc/otel-agent-config.yaml", "--log-level=DEBUG"]
    environment:
      - AWS_ACCESS_KEY_ID=<set your aws key> // TO EDIT
      - AWS_SECRET_ACCESS_KEY=<set your aws credential> // TO EDIT
      - AWS_REGION=<aws region> // TO EDIT
    volumes:
      - ../config.yaml:/etc/otel-agent-config.yaml // use default config
      - ~/.aws:/root/.aws
    ports:
      - "1777:1777"   # pprof extension
      - "55679:55679" # zpages extension
      - "55680:55680" # OTLP receiver
      - "13133"       # health_check 
```
2. Once your own docker image and AWS credential has been updated in the ```docker-composite.yaml``` file, run the following make command.
```
make docker-composite
```
3. Now you can view you data in AWS console

    * X-Ray - aws console
    * CloudWatch - aws console  
    
**AWS Metrics Sample Data**   
* ![aws metrics](../images/metrics_sample.png)  
**AWS Traces Sample Data**
* ![aws traces](../images/traces_sample.png)  

4. Stop the running AOC in Docker container
```
make docker-stop
```
