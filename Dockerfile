FROM golang:alpine AS STAGEONE

WORKDIR /usr/scratch
RUN apk -U add sqlite libspatialite libspatialite-dev gcc musl-utils tzdata make binutils dev86 musl-dev
COPY . .
RUN chmod 777 -R /tmp && chmod o+t -R /tmp
RUN go build -ldflags "-s -w" . 

COPY rpi-auto-display /bin/autodash
# #COPY /usr/share/zoneinfo /usr/share/zoneinfo
# COPY /lib/ld-musl-* /lib/
# COPY /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY /etc/passwd /etc/passwd
# COPY /etc/group /etc/group
# COPY /bin/sh /bin/sh
COPY data /data

ENTRYPOINT [ "/bin/autodash" ]