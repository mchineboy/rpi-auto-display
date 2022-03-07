FROM golang:latest

RUN go build . -o autodash

COPY autodash /bin/autodash

ENTRYPOINT [ "/bin/autodash" ]