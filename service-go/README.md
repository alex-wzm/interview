# Service (Go)

Serve the Interview API.

## Export enviornment

```sh
source env/local.env
```

## Run
You can pass the log level to the service executable with a flag `logLevel`, default lofLevel=4 (Info)
You can turn ON grpc middleware logging with a boolean flag `logGrpc`, default OFF 

```sh
go run cmd/service/service.go -logLevel=5 -logGrpc
```

## Test

```sh
go test ./...
```

# Containerize (Docker)
## Prerequisites
- Golang installation
- Docker installation
## Build service docker image
1. Run the build script.
`./build.sh`
2. Check the docker image, it's tagged interview-service
`docker images`
## Run docker container
3. Run the run script and pass the image ID found in step 2
`./run.sh 446bd644965b`
You can pass one or both flags to the executable: logLevel and logGrpc
`./run.sh  dcd4c4c0be76 -logLevel=5 -logGrpc`
4. Verify the docker container is running
`docker ps`
## Get service logs
Find container ID and Name in step 4
`docker logs <container ID or Name>`

## Stop the service
` docker stop <container ID or Name>`

## Send Client message to the Service
See [Client README](../client-go/README.md)

