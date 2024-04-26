ARG GO_VERSION=1.22
FROM golang:${GO_VERSION} as builder

WORKDIR /app

COPY . .

RUN go build -ldflags="-s" -trimpath -o client cmd/client/client.go

FROM debian

LABEL component="client" \
      repository="github.com/alex-wzm/the-interview-project"

COPY --from=builder /app/client /app/client

ENTRYPOINT ["/app/client"]
