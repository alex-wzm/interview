FROM golang:1.19.13-alpine AS builder
WORKDIR /build
COPY . ./
RUN go mod download
RUN mkdir bin
RUN cd cmd/service/ && go build -o ../../bin/

FROM golang:1.19.13-alpine
WORKDIR service-go/
EXPOSE 8080
COPY --from=builder /build/bin/service ./
ENTRYPOINT ["./service"]
