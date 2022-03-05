FROM golang:latest

RUN go build src/. -o autodash

COPY autodash /bin/autodash

ENTRYPOINT [ "/bin/autodash" ]