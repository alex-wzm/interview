ARG GO_VERSION=1.22
FROM golang:${GO_VERSION} as builder

WORKDIR /app

COPY . .

RUN go build -ldflags="-s" -trimpath -o service cmd/service/service.go

FROM debian

LABEL component="service" \
      repository="github.com/alex-wzm/the-interview-project"

EXPOSE 8080

COPY --from=builder /app/service /app/service

ENTRYPOINT ["/app/service"]
