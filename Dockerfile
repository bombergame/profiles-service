FROM golang:1.11-alpine as base
RUN apk add make
WORKDIR ${GOPATH}/src/github.com/bombergame/profiles-service
COPY . .
RUN make build && mv ./service /tmp/service

FROM alpine:latest
WORKDIR /tmp
COPY --from=base /tmp/service .
ENTRYPOINT ./service --http_port=80 --grpc_port=3000
EXPOSE 80 3000
