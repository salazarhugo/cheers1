#!/bin/bash

PATH=$PATH:$GOPATH/bin
#export PATH="$PATH:$(go env GOPATH)/bin"

cp cheers ../genproto -r

protoc --proto_path=. \
  --go_out=../genproto --go_opt=paths=source_relative  \
  --go-grpc_out=../genproto --go-grpc_opt=paths=source_relative \
  --experimental_allow_proto3_optional \
  cheers/type/party/party.proto \
  cheers/type/privacy/privacy.proto \
  cheers/type/user/user.proto \
  cheers/type/post/post.proto \
  cheers/party/v1/party_service.proto \
  cheers/post/v1/post_service.proto \
  cheers/user/v1/user_service.proto

