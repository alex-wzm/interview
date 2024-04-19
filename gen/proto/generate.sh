#!/bin/bash

# Base directory is the current directory where the script is run
BASE_DIR=$(pwd)

# Directories for generated code relative to gen/proto
SERVICE_TARGET_DIR="$BASE_DIR/../../service-go/internal/api/interview"
CLIENT_TARGET_DIR="$BASE_DIR/../../client-go/internal/api/interview"

# Create target directories if they do not exist
mkdir -p "$SERVICE_TARGET_DIR"
mkdir -p "$CLIENT_TARGET_DIR"

# Directory where the .proto files are located
PROTO_SOURCE_DIR="$BASE_DIR/interview"

# Generate the protocol buffers and gRPC stubs for service-go
echo "📁 Generating protocol buffers and gRPC stubs for service-go..."
protoc --proto_path="$PROTO_SOURCE_DIR" \
       --go_out="$SERVICE_TARGET_DIR" --go_opt=paths=source_relative \
       --go-grpc_out="$SERVICE_TARGET_DIR" --go-grpc_opt=paths=source_relative \
       "$PROTO_SOURCE_DIR"/interview.proto

echo "✅ Compiled proto stubs for service-go"

# Generate the protocol buffers and gRPC stubs for client-go
echo "📁 Generating protocol buffers and gRPC stubs for client-go..."
protoc --proto_path="$PROTO_SOURCE_DIR" \
       --go_out="$CLIENT_TARGET_DIR" --go_opt=paths=source_relative \
       --go-grpc_out="$CLIENT_TARGET_DIR" --go-grpc_opt=paths=source_relative \
       "$PROTO_SOURCE_DIR"/interview.proto

echo "✅ Compiled proto stubs for client-go"
