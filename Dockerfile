FROM golang:latest AS STAGEONE

WORKDIR /usr/scratch

COPY . .

RUN go build . 

RUN ls -lh

FROM scratch

COPY --from=STAGEONE /usr/scratch/rpi-auto-dash /bin/autodash

ENTRYPOINT [ "/bin/autodash" ]