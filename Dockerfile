FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git make

COPY . .

RUN go mod download 

CMD ["/app/inventoryApp"]
