# builder image
FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go build -o run_rpc_server ./cmd/rpc_server

# run image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/run_rpc_server .

EXPOSE 8080

#CMD ["./run_server"]
# never separate the command and output redirection
CMD [ "/bin/sh", "-c", "/app/run_rpc_server > /var/log/fx.log" ]


