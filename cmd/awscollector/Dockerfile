# The build mode (options: build, copy) passed in as a --build-arg. If build is specified, then the copy
# stage will be skipped and vice versa. The build mode builds the binary from the source files, while
# the copy mode copies in a pre-built binary.
ARG BUILDMODE=build

################################
#	       Base Stage          #
#			                   #
################################
FROM alpine:latest AS base

ARG USERNAME=aoc
ARG USER_UID=4317

RUN addgroup \
    -g $USER_UID \
    $USERNAME && \
    adduser \
    -D \
    -g $USERNAME \
    -h "/home/${USERNAME}"\
    -G $USERNAME \
    -u $USER_UID \
    $USERNAME

# Use --no-scripts to avoid QEMU emulation issues with busybox triggers during ARM64 cross-platform builds
RUN apk --update --no-scripts add ca-certificates && \
    update-ca-certificates

################################
#	       Build Stage         #
#			                   #
################################
FROM golang:1.24 AS prep-build

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
COPY build/linux/$TARGETARCH/healthcheck /workspace/healthcheck

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

ARG USERNAME=aoc

COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group
COPY --from=base --chown=$USERNAME:$USERNAME /home/$USERNAME/ /home/$USERNAME
COPY --from=package /workspace/awscollector /awscollector
COPY --from=package /workspace/config/ /etc/
COPY --from=package /workspace/healthcheck /healthcheck

ENV RUN_IN_CONTAINER="True"

USER $USERNAME
# aws-sdk-go needs $HOME to look up shared credentials
ENV HOME=/home/$USERNAME
ENTRYPOINT ["/awscollector"]
CMD ["--config=/etc/otel-config.yaml"]
EXPOSE 4317 55681 2000
