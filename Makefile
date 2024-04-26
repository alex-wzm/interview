CLIENT_CONTAINER_NAME := client
SERVICE_CONTAINER_NAME := service
AUTHSERVICE_CONTAINER_NAME := authservice
COMMIT_SHA := $(shell git rev-parse HEAD)

go-build-client:
	go build \
		-ldflags="-s" \
		-trimpath -o client ./client-go/cmd/client/client.go

go-build-service:
	go build \
		-ldflags="-s" \
		-trimpath -o service ./service-go/cmd/service/service.go

go-build-authservice:
	go build \
		-ldflags="-s" \
		-trimpath -o authservice ./authservice-go/cmd/authservice/authservice.go

go-build: go-build-client go-build-service go-build-authservice

go-test-client:
	go test -v -race ./client-go/...

go-test-service:
	go test -v -race ./service-go/...

go-test-authservice:
	go test -v -race ./authservice-go/...

go-test: go-test-client go-test-service go-test-authservice

docker-build-client:
	docker build -t $(CLIENT_CONTAINER_NAME) ./client-go

docker-build-service:
	docker build -t $(SERVICE_CONTAINER_NAME) ./service-go

docker-build-authservice:
	docker build -t $(AUTHSERVICE_CONTAINER_NAME) ./authservice-go

docker-build: docker-build-client docker-build-service docker-build-authservice

docker-run: docker-build-service docker-build-authservice
	docker compose up

setup:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	./setup.sh

clean:
	./clean.sh
	rm -f client service
