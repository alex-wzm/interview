ARG GO_VERSION=1.22
FROM golang:${GO_VERSION} as builder

WORKDIR /app

COPY . .

RUN go build -ldflags="-s" -trimpath -o authservice cmd/authservice/authservice.go

FROM debian

LABEL component="authservice" \
      repository="github.com/alex-wzm/the-interview-project"

COPY --from=builder /app/authservice /app/authservice

ENTRYPOINT ["/app/authservice"]
