################################
#	Building Stage         #	
#			       #	
################################
FROM alpine:latest as build 

# update cert
RUN apk --update add ca-certificates

WORKDIR /workspace

ARG TARGETPLATFORM
RUN echo "Target platform is $TARGETPLATFORM"

# copy artifacts
# always assume binary is created
COPY build/linux/* /workspace/
COPY config.yaml /workspace/config/otel-config.yaml
COPY config/ /workspace/config/
COPY cmd/awscollector/copy_docker_binary.sh /workspace/copy_binary.sh
RUN ./copy_binary.sh

################################
#	Final Stage            #	
#			       #	
################################
FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /workspace/build/linux/aoc_linux /awscollector
COPY --from=build /workspace/config/ /etc/

ENV RUN_IN_CONTAINER="True"
# aws-sdk-go needs $HOME to look up shared credentials
ENV HOME=/root
ENTRYPOINT ["/awscollector"]
CMD ["--config=/etc/otel-config.yaml"]
EXPOSE 4317 55681 2000

