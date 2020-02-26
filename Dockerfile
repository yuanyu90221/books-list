FROM golang:1.12.14

ARG ELEPHANT_SQL_URL

ENV ELEPHANT_SQL_URL=${ELEPHANT_SQL_URL}

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main