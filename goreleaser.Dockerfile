FROM debian:buster-slim

ARG TARGETPLATFORM

RUN set -ex \
  && if [ "${TARGETPLATFORM}" = "linux/amd64" ]; then export TARGETPLATFORM=amd64; fi \
  && if [ "${TARGETPLATFORM}" = "linux/arm64" ]; then export TARGETPLATFORM=arm64; fi 

WORKDIR /app

ADD tdex /usr/local/bin/tdex
ADD "tdexd-linux-$TARGETPLATFORM" /usr/local/bin/tdexd
ADD tdexdconnect /usr/local/bin/tdexdconnect


# $USER name, and data $DIR to be used in the `final` image
ARG USER=tdex
ARG DIR=/home/tdex

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates

# NOTE: Default GID == UID == 1000
RUN adduser --disabled-password \
            --home "$DIR/" \
            --gecos "" \
            "$USER"
USER $USER

# Prevents `VOLUME $DIR/.tdex-daemon/` being created as owned by `root`
RUN mkdir -p "$DIR/.tdex-daemon/"

# Expose volume containing all `tdexd` data
VOLUME $DIR/.tdex-daemon/

# expose trader and operator interface ports
EXPOSE 9945
EXPOSE 9000

ENTRYPOINT ["tdexd"]
