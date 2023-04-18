FROM golang:1.19-alpine

WORKDIR /go/work

COPY go.* ./
RUN go mod download

COPY cmd/service *.go ./cmd/service/
COPY internal/api/*.go ./internal/api/
COPY internal/api/interview/*.go ./internal/api/interview/
COPY internal/domain/greeter/greeter.go ./internal/domain/greeter/
COPY *.env ./

RUN go build  -o ./server ./cmd/service/*.go

EXPOSE 8080

ENTRYPOINT [ "./server" ]
