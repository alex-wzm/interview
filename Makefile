.PHONY: setup gen
setup: 
	./setup.sh
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

gen: setup
	cd ./gen/proto && ./generate.sh && cd ../..

.PHONY: run-service
run-service: gen
	$(MAKE) -C service-go run-service

.PHONY: run-client
run-client: gen
	$(MAKE) -C client-go run-client
