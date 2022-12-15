FROM golang:1.18-alpine3.16 AS builder

LABEL maintainer="Oleg_Dzhur"

WORKDIR /app

COPY . .

COPY ./configs ./configs

COPY ./migrate ./migrate

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# RUN chmod +x wait-for-postgres.sh

RUN go build -o main ./main.go
RUN apk add postgresql

EXPOSE 8080

CMD [ "echo","STARTING GOOSE IN CMD", "&&", "make", "goose", "&&", "./app/main" ]
