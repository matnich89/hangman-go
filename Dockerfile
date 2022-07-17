FROM golang:1.18 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/web/*.go

FROM alpine:latest AS production
EXPOSE 8080
COPY --from=builder /app .
CMD ["./app"]
