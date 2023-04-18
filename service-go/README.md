# Service (Go)

Serve the Interview API.

## Run
You can pass the log level to the service executable with a flag `logLevel`, default lofLevel=4 (Info)
You can turn ON grpc middleware logging with a boolean flag `logGrpc`, default OFF 

```sh
go run main.go  -logLevel=5 -logGrpc
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
5. Stop the service using container ID or name
` docker stop stupefied_villani`
`