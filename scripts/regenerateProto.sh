#! /bin/bash

# https://github.com/gogo/protobuf/issues/325
cd protobuf

protoc -I. \
-I/usr/local/include \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--gogofast_out=\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
plugins=grpc:../internal/ \
--gotemplate_out=../ \
--swagger_out=logtostderr=true:../api/docs \
--grpc-gateway_out=logtostderr=true:../internal/ \
proto/*.proto