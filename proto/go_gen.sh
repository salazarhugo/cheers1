export PATH="$PATH:$(go env GOPATH)/bin"

protoc --proto_path=. \
  --go_out=../genproto --go_opt=paths=source_relative  \
  --go-grpc_out=../genproto --go-grpc_opt=paths=source_relative \
  --experimental_allow_proto3_optional \
    cheers/api/v1/main.proto cheers/type/party.proto