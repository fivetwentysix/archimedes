FROM alpine:3.4

RUN apk add --no-cache ca-certificates

COPY ./archimedes /opt/
RUN chmod +x /opt/archimedes

ENTRYPOINT ["/opt/archimedes"]
