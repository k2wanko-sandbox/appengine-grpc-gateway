#!/bin/bash

protoc -I/usr/local/include -I. \
       -I$GOPATH/src \
       -I$GOPATH/src/github.com/gengo/grpc-gateway/third_party/googleapis \
       --swagger_out=logtostderr=true:. \
       --grpc-gateway_out=logtostderr=true:. \
       --go_out=Mgoogle/api/annotations.proto=github.com/gengo/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
       internal/echo/*.proto
