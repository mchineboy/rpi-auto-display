FROM golang:latest AS STAGEONE

WORKDIR /usr/scratch

COPY . .

RUN go build . 

FROM scratch

COPY --from=STAGEONE /usr/scratch/rpi-auto-display /bin/autodash

ENTRYPOINT [ "/bin/autodash" ]