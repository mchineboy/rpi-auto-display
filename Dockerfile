FROM golang:alpine AS STAGEONE

WORKDIR /usr/scratch
RUN apk -U add sqlite libspatialite libspatialite-dev gcc musl-utils tzdata make binutils dev86 musl-dev
COPY . .
RUN chmod 777 -R /tmp && chmod o+t -R /tmp
RUN go build -ldflags "-s -w" . 
RUN ldd /usr/scratch/rpi-auto-display

FROM scratch

COPY --from=STAGEONE /usr/scratch/rpi-auto-display /bin/autodash
#COPY --from=STAGEONE /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=STAGEONE /lib/ld-musl-* /lib/
COPY --from=STAGEONE /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=STAGEONE /etc/passwd /etc/passwd
COPY --from=STAGEONE /etc/group /etc/group
COPY --from=STAGEONE /bin/sh /bin/sh
COPY --from=STAGEONE /usr/scratch/data /data
COPY --from=STAGEONE /tmp /tmp

ENTRYPOINT [ "/bin/autodash" ]