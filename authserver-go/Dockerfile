FROM golang:1.19.13-alpine AS builder
WORKDIR /build
COPY . ./
RUN go mod download
RUN mkdir bin
RUN cd cmd/authserver/ && go build -o ../../bin/

FROM golang:1.19.13-alpine
WORKDIR authserver-go/
EXPOSE 8081
COPY --from=builder /build/bin/authserver ./
ENTRYPOINT ["./authserver"]



