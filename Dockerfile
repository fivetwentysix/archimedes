FROM ubuntu:14.04

RUN apt-get update \
 && apt-get install -y --no-install-recommends ca-certificates \
 && apt-get clean
 
COPY ./archimedes /opt/
RUN chmod +x /opt/archimedes

ENTRYPOINT ["/opt/archimedes"]
