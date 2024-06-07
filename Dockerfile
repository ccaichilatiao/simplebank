# Build stage
FROM golang:1.20-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine-alpine3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

EXPOSE [8080, 9090]
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh", "/app/main" ]