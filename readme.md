# Disclaimer
It's essential to understand in which part of the program something goes wrong. Primarily it's important when we have a lot of services that interact with each other.

# About the project
The folder structure of the project one layer higher than the standard one. Because we have more than one services
IDL for example should be a dedicated repository, and every service is a separate repository.


## Idea
We will use a middleware that set logID header for every http request. The same logID can be transferred in the base part of the RPC requests.


## Tech info
We have 2 services, http gateway to accept a user request, and rpc FX service. Our main goal to create a logging system between microservices.

### HTTP gateway
It's a simple service that accepts a request from a user, checks auth, takes response from the FX service and return it to the user.

### RPC FX service
It's our intenal service that accepts a request from the gateway, and 


# Logging system
Loggers are define in logger package, it has a structured and sugared loggers.

There is a middleware that logs every request and time of processing.

Additionally logger has a withContext wrapper,
the withContext func takes the trackID from the ctx, to avoid passing the trackID in every log function.



## Emulate error behaviour
Let's assume our fx can't process amounts more than 1000. We will emulate this behaviour in the rpc service.

Let's assume our http gateway can't process SGD currency. We will emulate this behaviour in the http gateway.     

# Building and running the project

## Running in docker compose
```bash
 docker-compose up -d
 ```

check that all services are running
```bash
docker ps
```



## Running compiled binaries
### 1. Run FX service

```bash
cd fx
go run ./cmd/rpc_server/main.go
```

you should see
```bash
2022/12/28 00:19:00 server listening at [::]:50051
```

validate that the service is running with rpc client
```bash
go run ./cmd/client_get_fx/main.go
```
todo: 2. check client run


# Check the logs

## Run an fx request
Send requests from gateway/api/example.http or use curl:
```bash
curl -X POST \
  http://localhost:8080/api/v1/fx \
  -H 'Content-Type: application/json' \
  -d '{
          "SourceCurrency": "CHF",
          "TargetCurrency": "USD",
          "SourceAmount": 100.09
      }'
```
In the response you can find a trackID header to track the request in the logs between microservices.

After that you can open grafana in your browser http://localhost:3000 and check the logs
