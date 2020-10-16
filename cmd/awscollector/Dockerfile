FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY awscollector /
COPY config.yaml /etc/otel-config.yaml
ENV RUN_IN_CONTAINER="True"
ENTRYPOINT ["/awscollector"]
CMD ["--config=/etc/otel-config.yaml"]
EXPOSE 55679 55680 55681 13133 1777
