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

# copy artifacts
WORKDIR /workspace
COPY cmd  /workspace/cmd
COPY config.yaml /workspace/config.yaml
COPY pkg /workspace/pkg
COPY tools /workspace/tools
COPY go.mod /workspace/go.mod
COPY go.sum /workspace/go.sum
COPY VERSION /workspace/VERSION
COPY .git /workspace/.git
COPY Makefile /workspace/Makefile

# build binary
RUN make build

################################
#	Final Stage            #	
#			       #	
################################
FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /workspace/build/linux/aoc_linux_x86_64 /awscollector
COPY --from=build /workspace/config.yaml /etc/otel-config.yaml

ENV RUN_IN_CONTAINER="True"
ENTRYPOINT ["/awscollector"]
CMD ["--config=/etc/otel-config.yaml"]
EXPOSE 55680 55681
