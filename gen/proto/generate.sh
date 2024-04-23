#!/bin/bash

SERVICE_RELATIVE_PATH="../../service-go/internal/api"

mkdir -p $SERVICE_RELATIVE_PATH

protoc --go_out=$SERVICE_RELATIVE_PATH --go_opt=paths=source_relative \
    --go-grpc_out=$SERVICE_RELATIVE_PATH --go-grpc_opt=paths=source_relative \
    interview/interview.proto

# protoc --go_out=$SERVICE_RELATIVE_PATH --go_opt=paths=source_relative \
#     --go-grpc_out=$SERVICE_RELATIVE_PATH --go-grpc_opt=paths=source_relative \
#     auth/auth.proto

echo "✅ Compiled proto stubs for service-go"


CLIENT_RELVATIVE_PATH="../../client-go/internal/api"

mkdir -p $CLIENT_RELVATIVE_PATH

protoc --go_out=$CLIENT_RELVATIVE_PATH --go_opt=paths=source_relative \
    --go-grpc_out=$CLIENT_RELVATIVE_PATH --go-grpc_opt=paths=source_relative \
    interview/interview.proto

protoc --go_out=$CLIENT_RELVATIVE_PATH --go_opt=paths=source_relative \
    --go-grpc_out=$CLIENT_RELVATIVE_PATH --go-grpc_opt=paths=source_relative \
    auth/auth.proto

echo "✅ Compiled proto stubs for client-go"


AUTHSERVER_RELATIVE_PATH="../../authserver-go/internal/api"

mkdir -p $AUTHSERVER_RELATIVE_PATH

protoc --go_out=$AUTHSERVER_RELATIVE_PATH --go_opt=paths=source_relative \
    --go-grpc_out=$AUTHSERVER_RELATIVE_PATH --go-grpc_opt=paths=source_relative \
    auth/auth.proto

echo "✅ Compiled proto stubs for authserver-go"
