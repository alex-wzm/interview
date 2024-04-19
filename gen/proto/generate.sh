#!/bin/bash -e

SERVICE_RELATIVE_PATH="../../service-go/internal/api"
AUTH_SERVICE_RELATIVE_PATH="../../auth-go/internal/api"

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

protoc --go_out=$CLIENT_RELVATIVE_PATH --go_opt=paths=source_relative \
    --go-grpc_out=$CLIENT_RELVATIVE_PATH --go-grpc_opt=paths=source_relative \
    interview/auth/auth.proto

protoc --go_out=$AUTH_SERVICE_RELATIVE_PATH --go_opt=paths=source_relative \
    --go-grpc_out=$AUTH_SERVICE_RELATIVE_PATH --go-grpc_opt=paths=source_relative \
    interview/auth/auth.proto

echo "✅ Compiled proto stubs for auth-go"
