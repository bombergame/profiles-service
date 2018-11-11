FROM golang:1.11-alpine as base
RUN apk add make
WORKDIR ${GOPATH}/src/github.com/bombergame/profiles-service
COPY . .
RUN make build && mv ./_build/service /tmp/service && mv ./repositories/postgres/scripts /tmp/scripts

FROM alpine:latest
WORKDIR /tmp
COPY --from=base /tmp/service .
COPY --from=base /tmp/scripts ./scripts
ENTRYPOINT ./service --http_port=80 --grpc_port=3000
EXPOSE 80 3000
