FROM golang:1.20

WORKDIR /usr/src/app

RUN apt-get update  \
    && apt-get install -y make
RUN apt-get install -y postgresql-client

COPY go.mod go.sum ./
RUN go mod download \
    && go mod verify \
    && go install github.com/pressly/goose/v3/cmd/goose@latest

COPY ./ ./
RUN chmod u+x ./entrypoint.sh \
    && make build

ENTRYPOINT ["/bin/sh", "./entrypoint.sh"]