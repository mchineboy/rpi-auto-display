FROM golang:latest AS STAGEONE

WORKDIR /usr/scratch

COPY . .

RUN go build . 

ENTRYPOINT [ "/usr/scratch/rpi-auto-display" ]