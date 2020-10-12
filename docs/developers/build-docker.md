### Build Docker Image

aws-otel-collector is an open source project, we’re looking for the contributions from the open source community to make it better. You can build your own executable binaries for your own customization. We provide the following command for you the build your own executables.

#### Steps,
1. Update DockerHub namespace you want to upload your docker image. (you can skip this step if you only want to build and run docker image locally)
2. Build your own docker image.  
```make docker-build ```
3. Push your own aws-otel-collector docker image to your own docker hub repository.  
```make docker-push```
4. Run AOC docker image with the default configuration file.  
```docker-compose up -d```
