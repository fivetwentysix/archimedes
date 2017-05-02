FROM ubuntu:14.04

RUN apt-get update \
 && apt-get install -y --no-install-recommends ca-certificates \
 && apt-get clean

RUN mkdir /opt/data
COPY ./data/zip.csv /opt/data

COPY ./archimedes /opt/
RUN chmod +x /opt/archimedes

WORKDIR /opt
ENTRYPOINT ["/opt/archimedes"]
