FROM golang:1.19.8-buster as builder

COPY ./ /go/src
WORKDIR /go/src
RUN go get -v -t -d ./...
RUN go build -o /app/bot

FROM gcr.io/distroless/base

COPY --from=builder /app /app
WORKDIR /app

CMD ["./bot"]
