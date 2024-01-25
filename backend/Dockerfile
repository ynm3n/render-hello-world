FROM golang:latest AS builder

WORKDIR /app
COPY . .
RUN go build ./cmd/app

FROM debian:stable-slim
WORKDIR /app
COPY --from=builder app .
EXPOSE 8080
CMD [ "./app" ]
