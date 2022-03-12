FROM golang:latest AS STAGEONE

WORKDIR /usr/scratch

COPY . .

RUN CGO_ENABLED=0 go build . 

FROM scratch

COPY --from=STAGEONE /usr/scratch/rpi-auto-display /bin/autodash
COPY --from=STAGEONE /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=STAGEONE /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=STAGEONE /etc/passwd /etc/passwd
COPY --from=STAGEONE /etc/group /etc/group
COPY --from=STAGEONE /bin/sh /bin/sh
COPY --from=STAGEONE /usr/scratch/data /data
RUN mkdir /tmp && chmod 777 -R /tmp && chmod o+t -R /tmp

ENTRYPOINT [ "/bin/autodash" ]