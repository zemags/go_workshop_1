FROM golang:1.15-alpine as builder

WORKDIR /opt/code
ADD ./ /opt/code

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN go mod download

# bilding binary in folder bin/workshop_1 with compiling main.go
RUN go build -o bin/workshop_1 cmd/workshop/main.go
ENTRYPOINT [ "bin/workshop_1" ]
