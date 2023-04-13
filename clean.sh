#!/bin/bash

echo "📁 Removing generated files:"

rm -v -f ./client-go/internal/api/interview/*pb.go

rm -v -f ./service-go/internal/api/interview/*pb.go

echo "🏁 Generated files removed"