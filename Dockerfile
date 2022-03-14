FROM golang:alpine AS STAGEONE

WORKDIR /usr/scratch
RUN apk -U --no-cache add sqlite libspatialite 
RUN apk -U --no-cache add --virtual .build libspatialite-dev gcc musl-utils tzdata make binutils dev86 musl-dev
COPY . .
RUN chmod 777 -R /tmp && chmod o+t -R /tmp
RUN GOOS=linux GOARCH=arm64 go build -v -ldflags "-s -w" --tags "sqlite_stat4 sqlite_vacuum_full" . 

COPY data /data
COPY fonts /fonts
RUN apk del .build
ENTRYPOINT [ "/usr/scratch/rpi-auto-display" ]