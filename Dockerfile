FROM golang:latest AS STAGEONE

WORKDIR /usr/scratch

COPY . .

RUN go build . 

FROM scratch

COPY --from=STAGEONE /usr/scratch/rpi-auto-display /bin/autodash
COPY --from=STAGEONE /etc/hosts /etc/hosts
COPY --from=STAGEONE /etc/passwd /etc/passwd

ENTRYPOINT [ "/bin/autodash" ]