FROM golang:alpine as builder

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bot ./cmd/bot

FROM alpine

WORKDIR /app

COPY --from=builder /src/bot .
COPY --from=builder /src/bot.yml .



ENTRYPOINT ["/app/bot"]