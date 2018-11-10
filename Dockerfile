FROM golang:1.11-alpine as base
RUN apk add make
WORKDIR ${GOPATH}/src/github.com/bombergame/profiles-service
COPY . .
RUN make build && mv ./_build/service /tmp/service && mv ./repositories/postgres/scripts /tmp/scripts

FROM alpine:latest
WORKDIR /tmp
COPY --from=base /tmp/service .
COPY --from=base /tmp/scripts ./scripts
ENTRYPOINT ls -al -R
ENTRYPOINT ./service --http_port=80 --init_storage --storage_scripts_path=./scripts
EXPOSE 80
