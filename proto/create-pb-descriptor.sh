export PATH="$PATH:$(go env GOPATH)/bin"

cp cheers ../genproto -r

protoc \
  --include_imports \
  --include_source_info \
  --proto_path=. \
  --descriptor_set_out=../genproto/api_descriptor.pb \
  --experimental_allow_proto3_optional \
  cheers/type/party/party.proto \
  cheers/type/privacy/privacy.proto \
  cheers/type/user/user.proto \
  cheers/type/post/post.proto \
  cheers/party/v1/party_service.proto \
  cheers/post/v1/post_service.proto \
  cheers/user/v1/user_service.proto
