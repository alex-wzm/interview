# Service (Go)

Serve the Interview API.

## Run
You can pass the log level to the service executable with a flag `logLevel`, default lofLevel=4 (Info)
You can turn ON grpc middleware logging with a boolean flag `grpcLogs`, default OFF 

```sh
<<<<<<< HEAD
go run cmd/service/service.go
=======
go run main.go  -logLevel=5 -grpcLogs
>>>>>>> e3cc5eb (PR feedback changes)
```

## Test

```sh
go test ./...
```

## export enviornment

```sh
source env/local.env
```
