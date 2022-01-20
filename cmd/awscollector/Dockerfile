################################
#	Building Stage         #	
#			       #	
################################
FROM alpine:latest as build 

# update cert
RUN apk --update add ca-certificates

WORKDIR /workspace

ARG TARGETARCH

# copy artifacts
# always assume binary is created
COPY build/linux/$TARGETARCH/aoc /workspace/awscollector
COPY config.yaml /workspace/config/otel-config.yaml
COPY config/ /workspace/config/

################################
#	Final Stage            #	
#			       #	
################################
FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /workspace/awscollector /awscollector
COPY --from=build /workspace/config/ /etc/

ENV RUN_IN_CONTAINER="True"
# aws-sdk-go needs $HOME to look up shared credentials
ENV HOME=/root
ENTRYPOINT ["/awscollector"]
CMD ["--config=/etc/otel-config.yaml"]
EXPOSE 4317 55681 2000

