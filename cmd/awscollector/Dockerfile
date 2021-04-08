################################
#	Building Stage         #	
#			       #	
################################
FROM alpine:latest as build 

# update cert
RUN apk --update add ca-certificates

# install libs
RUN apk add --update bash git make build-base

# config go
COPY --from=golang:1.15.3-alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

# download go modules ahead to speed up the building
WORKDIR /workspace
COPY . /workspace
RUN go mod download

# build binary
RUN make amd64-build

################################
#	Final Stage            #	
#			       #	
################################
FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /workspace/build/linux/aoc_linux_x86_64 /awscollector
COPY --from=build /workspace/config.yaml /etc/otel-config.yaml
COPY --from=build /workspace/config/eks/prometheus/config-appmesh.yaml /etc/eks/prometheus/config-appmesh.yaml
COPY --from=build /workspace/config/eks/prometheus/config-haproxy.yaml /etc/eks/prometheus/config-haproxy.yaml
COPY --from=build /workspace/config/eks/prometheus/config-jmx.yaml /etc/eks/prometheus/config-jmx.yaml
COPY --from=build /workspace/config/eks/prometheus/config-memcached.yaml /etc/eks/prometheus/config-memcached.yaml
COPY --from=build /workspace/config/eks/prometheus/config-nginx.yaml /etc/eks/prometheus/config-nginx.yaml
COPY --from=build /workspace/config/eks/prometheus/config-redis.yaml /etc/eks/prometheus/config-redis.yaml
COPY --from=build /workspace/config/eks/prometheus/config-all.yaml /etc/eks/prometheus/config-all.yaml
COPY --from=build /workspace/config/ecs/container-insights/otel-task-metrics-config.yaml /etc/ecs/container-insights/otel-task-metrics-config.yaml
COPY --from=build /workspace/config/ecs/ecs-default-config.yaml /etc/ecs/ecs-default-config.yaml

ENV RUN_IN_CONTAINER="True"
# aws-sdk-go needs $HOME to look up shared credentials
ENV HOME=/root
ENTRYPOINT ["/awscollector"]
CMD ["--config=/etc/otel-config.yaml"]
EXPOSE 4317 55681 2000

