# builder image
FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go build -o run_web_server ./cmd/web_server

# run image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/run_web_server .

EXPOSE 8080

#CMD ["./run_server"]
# never separate the command and output redirection
CMD [ "/bin/sh", "-c", "/app/run_web_server > /var/log/gateway.log" ]


