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

## export enviornment

```sh
source env/local.env
```
