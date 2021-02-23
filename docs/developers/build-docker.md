### Build Docker Image

aws-otel-collector is an open source project, we’re looking for the contributions from the open source community to make it better. You can build your own executable binaries for your own customization. We provide the following command for you to build your own executables.

#### Steps,
1. Update DOCKER_NAMESPACE in Makefile if you want to upload your docker image to Docker Hub. (you can skip this step if you only want to build and run docker image locally)
2. Build your own docker image.  
```make docker-build ```
3. Run AWSOTelCollector docker image with the default configuration file. By default, it will pull the latest image from ECR. Modify docker-compose.yaml to use other images.  
```cd examples; docker-compose up -d```
