# About IDL repository
The repository contains proto files, it's a main repository for the organisation. All developers who want to publish their microservice can add document here.

## Server integration
Define proto file for your service and create a PR

## Client integration
If you need to interact with other services by RPC, you can use the proto files from this repository. Generate an RPC client and make a request to the service.

## Base.proto
Base message should be integrated in all RPC requests and responses. It contains a logID field that will be used for logging.
(Potentially we can add more fields to the base message, auth, dates, callerID, etc.)

## For FX project
Gateway uses the proto files from this repository to generate a client for the FX service. The client is located in the `gateway/internal/fx` folder.

FX service uses the proto files from this repository to generate a server. The server is located in the `fx/internal/fx` folder.

# Generate code
This section describe how to generate the code for populating, serializing, and retrieving request and response message types.

## generate server
```bash
cd fx
protoc --proto_path=../idl \
    --go_out=./internal/rpc_gen --go_opt=paths=source_relative \
    --go-grpc_out=./internal/rpc_gen --go-grpc_opt=paths=source_relative \
    ../idl/fx.proto ../idl/base.proto
```

## generate client
```bash
cd gateway
protoc --proto_path=../idl \
    --go_out=./internal/rpc_gen --go_opt=paths=source_relative \
    --go-grpc_out=./internal/rpc_gen --go-grpc_opt=paths=source_relative \
    ../idl/fx.proto ../idl/base.proto
```