FROM golang:1.11-alpine as base
RUN apk add make
WORKDIR ${GOPATH}/src/github.com/bombergame/profiles-service
COPY . .
RUN make build && mv ./_build/service /tmp/service

FROM alpine:latest
WORKDIR /tmp
COPY --from=base /tmp/service .
ENTRYPOINT ./service --service_port=80
EXPOSE 80
