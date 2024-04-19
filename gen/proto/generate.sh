#!/bin/bash

SERVICE_RELATIVE_PATH="../../service-go/internal/api"

mkdir -p $SERVICE_RELATIVE_PATH

protoc --go_out=$SERVICE_RELATIVE_PATH --go_opt=paths=source_relative \
    --go-grpc_out=$SERVICE_RELATIVE_PATH --go-grpc_opt=paths=source_relative \
    interview/interview.proto

echo "✅ Compiled proto stubs for service-go"

CLIENT_RELVATIVE_PATH="../../client-go/internal/api"

mkdir -p $CLIENT_RELVATIVE_PATH

protoc --go_out=$CLIENT_RELVATIVE_PATH --go_opt=paths=source_relative \
    --go-grpc_out=$CLIENT_RELVATIVE_PATH --go-grpc_opt=paths=source_relative \
    interview/interview.proto

echo "✅ Compiled proto stubs for client-go"

SERVICE_RELATIVE_PATH="../../auth-service/internal/api"

mkdir -p $SERVICE_RELATIVE_PATH

protoc --go_out=$SERVICE_RELATIVE_PATH --go_opt=paths=source_relative \
    --go-grpc_out=$SERVICE_RELATIVE_PATH --go-grpc_opt=paths=source_relative \
    authservice/authservice.proto

echo "✅ Compiled proto stubs for auth service in auth-service"

CLIENT_RELVATIVE_PATH="../../client-go/internal/api"

mkdir -p $CLIENT_RELVATIVE_PATH

protoc --go_out=$CLIENT_RELVATIVE_PATH --go_opt=paths=source_relative \
    --go-grpc_out=$CLIENT_RELVATIVE_PATH --go-grpc_opt=paths=source_relative \
    authservice/authservice.proto

echo "✅ Compiled proto stubs for auth service client-go"

CLIENT_RELVATIVE_PATH="../../service-go/internal/api"

mkdir -p $CLIENT_RELVATIVE_PATH

protoc --go_out=$CLIENT_RELVATIVE_PATH --go_opt=paths=source_relative \
    --go-grpc_out=$CLIENT_RELVATIVE_PATH --go-grpc_opt=paths=source_relative \
    authservice/authservice.proto

echo "✅ Compiled proto stubs for auth service service-go"
