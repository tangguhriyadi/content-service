FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o /content_service

FROM alpine:latest


COPY --from=builder /content_service /content_service
CMD ["/content_service"]