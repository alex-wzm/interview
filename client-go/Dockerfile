FROM golang:1.22-alpine3.19
WORKDIR /interview-client
ENV CGO_ENABLED 0
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build ./cmd/client/
CMD ["./client"]
