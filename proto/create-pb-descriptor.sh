export PATH="$PATH:$(go env GOPATH)/bin"

protoc \
  --include_imports \
  --include_source_info \
  --proto_path=. \
  --descriptor_set_out=api_descriptor.pb \
  --experimental_allow_proto3_optional \
  cheers/api/v1/main.proto cheers/type/party/party.proto cheers/type/privacy/privacy.proto
