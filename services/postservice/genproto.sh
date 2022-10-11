#!/bin/bash

PATH=$PATH:$GOPATH/bin
protodir=../../proto

protoc --proto_path=$protodir \
  --go_out=./genproto --go_opt=paths=source_relative  \
  --go-grpc_out=./genproto --go-grpc_opt=paths=source_relative \
  --experimental_allow_proto3_optional \
   $protodir/cheers/post/v1/post_service.proto