FROM golang:1.23.1-alpine3.20

WORKDIR /app
COPY ./ ./

WORKDIR /app
RUN go mod download

RUN go install go.uber.org/mock/mockgen@v0.4.0
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.27.0
RUN go install github.com/pressly/goose/v3/cmd/goose@v3.22.1