#!/usr/bin/env bash

### protoc ###
curl -OL https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
unzip -o protoc-3.6.1-linux-x86_64.zip bin/protoc -d ${GOPATH}
rm -f protoc-3.6.1-linux-x86_64.zip
