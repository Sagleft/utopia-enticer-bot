FROM golang:1.18.3-buster as builder
WORKDIR /app
COPY . .
RUN go build -o /app/bot

CMD ["./bot"]
