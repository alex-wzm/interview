FROM golang:1.19.13-alpine AS builder
WORKDIR /build
COPY . ./
RUN go mod download
RUN mkdir bin
RUN cd cmd/client/ && go build -o ../../bin/

FROM golang:1.19.13-alpine
WORKDIR client-go/
COPY --from=builder /build/bin/client ./
CMD ["./client"]
