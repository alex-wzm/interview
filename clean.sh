#!/usr/bin/env bash

echo "📁 Removing generated files:"

rm -vf ./client-go/internal/api/**/*pb.go
rm -vf ./service-go/internal/api/**/*pb.go
rm -vf ./authservice-go/internal/api/**/*pb.go

echo "🏁 Generated files removed"
