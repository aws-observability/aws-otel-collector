# The build mode (options: build, copy) passed in as a --build-arg. If build is specified, then the copy
# stage will be skipped and vice versa. The build mode builds the binary from the source files, while
# the copy mode copies in a pre-built binary.
ARG BUILDMODE=build

################################
#	Certificate Stage      #
#			       #
################################
FROM alpine:latest AS certs

RUN apk --update add ca-certificates

################################
#	Build Stage            #
#			       #
################################
FROM golang:1.17 AS prep-build

ARG TARGETARCH

# pass in the GOPROXY as a --build-arg (e.g. --build-arg GOPROXY=direct)
ARG GOPROXY
ENV GOPROXY=${GOPROXY}

# download go modules ahead to speed up the building
WORKDIR /workspace
COPY go.mod .
COPY go.sum .
RUN go mod download -x

# copy source
COPY . .

# build
RUN make ${TARGETARCH}-build

# move
RUN mv /workspace/build/linux/$TARGETARCH/aoc /workspace/awscollector

################################
#	Copy Stage             #
#			       #	
################################
FROM scratch AS prep-copy

WORKDIR /workspace

ARG TARGETARCH

# copy artifacts
# always assume binary is created
COPY build/linux/$TARGETARCH/aoc /workspace/awscollector

################################
#	Packing Stage          #
#			       #
################################
FROM prep-${BUILDMODE} AS package

COPY config.yaml /workspace/config/otel-config.yaml
COPY config/ /workspace/config/

################################
#	Final Stage            #	
#			       #	
################################
FROM scratch

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=package /workspace/awscollector /awscollector
COPY --from=package /workspace/config/ /etc/

ENV RUN_IN_CONTAINER="True"
# aws-sdk-go needs $HOME to look up shared credentials
ENV HOME=/root
ENTRYPOINT ["/awscollector"]
CMD ["--config=/etc/otel-config.yaml"]
EXPOSE 4317 55681 2000
