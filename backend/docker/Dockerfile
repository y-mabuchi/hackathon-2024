FROM golang:1.21.2-alpine

RUN apk update && apk add git

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/golang/mock/mockgen@latest

WORKDIR /app

CMD ["air", "-c", ".air.toml"]
