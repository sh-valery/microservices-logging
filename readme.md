# Disclaimer
It's essential to understand in which part of the program something goes wrong. Primarily it's important when we have a lot of services that interact with each other.
The project show how to orginize logging in the organisation, that use different services and intercation ways between them.

# About the project
The folder structure of the project one layer higher than usualy, it can be considered as a mono-repository for the organisation.
IDL for example should be a dedicated repository, fx service and gateway as well.



## Logging base
Every request should be logged with the requestID. The requestID id is generated in the gateway by tracking middleware and injected to headers and context, to pass to other services. 
For demonstration requestID provides in the response header. The requestID is used to find the request in the logs.

### Logging in the http interaction
When service is interacted with other services by HTTP we use header to store  the requestID.

### Logging in the grpc interaction
When service is interacted with other services by GRPC we have 2 ways to pass the requestID.
* Use metadata to store the requestID.
* Use Base rpc struct and incapsulate it in the all request and response structs.
In the current project we use both ways for demonstration.

## Logging context
Logger has a context wrapper that takes the requestID from the context and adds it to the log message.
To avoid manual adding the requestID to the context.

## Tech info
We have 2 services, http gateway to accept a user request, and rpc FX service. Our main goal to create a logging system between microservices.

### HTTP gateway
It's a simple service that accepts a request from a user, checks auth, takes response from the FX service and return it to the user.

### RPC FX service
It's our intenal service that accepts a request from the gateway, and
returns fx rate for the currency pair. API is defined in IDL folder.


# Logging system
Loggers are define in logger package, it has a structured and sugared loggers.
There is a middleware that logs every request and time of processing.



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

## Import data sources and dashboards
Grafana API can be used to import data sources and dashboards to avoid manual import with grafana fe

```bash
### import loki data source to grafana
POST http://localhost:3000/api/datasources
Authorization: Basic admin admin
Content-Type: application/json

{
  "orgId": 1,
  "name": "Loki",
  "type": "loki",
  "typeName": "Loki",
  "typeLogoUrl": "public/app/plugins/datasource/loki/img/loki_icon.svg",
  "access": "proxy",
  "url": "http://loki:3100",
  "password": "",
  "user": "",
  "database": "",
  "basicAuth": false,
  "isDefault": false,
  "jsonData": {},
  "readOnly": false
}
```
```bash
### import prometheus data source to grafana
POST http://localhost:3000/api/datasources
Authorization: Basic admin admin
Content-Type: application/json

{
  "orgId": 1,
  "name": "Prometheus",
  "type": "prometheus",
  "typeName": "Prometheus",
  "typeLogoUrl": "public/app/plugins/datasource/prometheus/img/prometheus_logo.svg",
  "access": "proxy",
  "url": "http://prometheus:9090",
  "password": "",
  "user": "",
  "database": "",
  "basicAuth": false,
  "isDefault": true,
  "jsonData": {
    "httpMethod": "POST"
  },
  "readOnly": false
}
```


## Run an fx request
Send requests from gateway/api/example.http or use curl:
```bash
curl -X POST \
  http://localhost:8080/api/v1/fx \
  -H 'Content-Type: application/json' \
  -d '{
          "SourceCurrency": "USD",
          "TargetCurrency": "CHF",
          "SourceAmount": 100.09
      }'
```
In the response you can find a trackID header to track the request in the logs between microservices.

After that you can open grafana in your browser http://localhost:3000 and check the logs

### Check the logs in grafana


### Generate an error in the fx service

fx service supports only `USD` as source currencies and 		`USD, EUR, CHF, GBP, JPY, CNY` as target currencies.
If you send a request with other currencies you will get an error.
Lets generate a request with SGD currency from gateway and find where the
error happened.

```bash
curl -X POST \
  http://localhost:8080/api/v1/fx \
  -H 'Content-Type: application/json' \
  -d '{
          "SourceCurrency": "USD",
          "TargetCurrency": "SGD",
          "SourceAmount": 100.00
      }'
```
