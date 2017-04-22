FROM alpine:3.4

RUN apk add --no-cache ca-certificates

ADD archimedes archimedes
RUN chmod +x archimedes

CMD ["./archimedes"]
