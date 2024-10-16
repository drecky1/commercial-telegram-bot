FROM golang:1.22
WORKDIR /app
COPY . ./
RUN env CGO_ENABLED=0 go build -v -o bot cmd/bot/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=0 /app/bot ./
CMD ["./bot"]