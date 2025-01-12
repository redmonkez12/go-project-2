FROM golang:1.23-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:3.21.2
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY db/migrations ./db/migrations

EXPOSE 8080
CMD [ "/app/main" ]
