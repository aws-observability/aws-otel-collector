################################
#	Building Stage         #	
#			       #	
################################
FROM alpine:latest as build 

# update cert
RUN apk --update add ca-certificates

WORKDIR /workspace

ARG TARGETPLATFORM
RUN echo "Targate platform is $TARGETPLATFORM"

# copy artifacts
# always assume binary is created
COPY build/linux/* /workspace/
COPY cmd/awscollector/copy_docker_binary.sh  /workspace/cmd/awscollector/copy_docker_binary.sh
RUN cmd/awscollector/copy_docker_binary.sh

################################
#	Final Stage            #	
#			       #	
################################
FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /workspace/build/linux/aoc_linux /awscollector
COPY config.yaml /etc/otel-config.yaml
COPY config/eks/prometheus/config-appmesh.yaml /etc/eks/prometheus/config-appmesh.yaml
COPY config/eks/prometheus/config-haproxy.yaml /etc/eks/prometheus/config-haproxy.yaml
COPY config/eks/prometheus/config-jmx.yaml /etc/eks/prometheus/config-jmx.yaml
COPY config/eks/prometheus/config-memcached.yaml /etc/eks/prometheus/config-memcached.yaml
COPY config/eks/prometheus/config-nginx.yaml /etc/eks/prometheus/config-nginx.yaml
COPY config/eks/prometheus/config-redis.yaml /etc/eks/prometheus/config-redis.yaml
COPY config/eks/prometheus/config-all.yaml /etc/eks/prometheus/config-all.yaml
COPY config/ecs/container-insights/otel-task-metrics-config.yaml /etc/ecs/container-insights/otel-task-metrics-config.yaml
COPY config/ecs/ecs-*.yaml /etc/ecs/
COPY config/ecs/otel-instance-metrics-config.yaml /etc/ecs/otel-instance-metrics-config.yaml
COPY config/apprunner/apprunner-default-config.yaml /etc/apprunner/apprunner-default-config.yaml

ENV RUN_IN_CONTAINER="True"
# aws-sdk-go needs $HOME to look up shared credentials
ENV HOME=/root
ENTRYPOINT ["/awscollector"]
CMD ["--config=/etc/otel-config.yaml"]
EXPOSE 4317 55681 2000

