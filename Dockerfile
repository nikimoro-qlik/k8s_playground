FROM alpine:3.18

RUN apk --update --no-cache upgrade
RUN apk add --no-cache ca-certificates


ARG VERSION
ARG REVISION
ARG CREATED

LABEL org.opencontainers.image.created=$CREATED
LABEL org.opencontainers.image.version=$VERSION
LABEL org.opencontainers.image.revision=$REVISION
LABEL org.opencontainers.image.url="https://ghcr.io/nikimoro-qlik/k8s_playground"
LABEL org.opencontainers.image.source="https://github.com/nikimoro-qlik/k8s_playground"

ADD k8s_playground k8s_playground

RUN addgroup -S 66900 && adduser -D -S -G 66900 66900
USER 66900:66900

ENTRYPOINT ["/k8s_playground"]
