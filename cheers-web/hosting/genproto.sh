PROTOC_GEN_TS_PATH="./node_modules/.bin/protoc-gen-ts"
PROTOC_OUT_DIR="./proto/generated"
mkdir -p ${PROTOC_OUT_DIR}
protoc \
       --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
       --js_out="import_style=commonjs,binary:${PROTOC_OUT_DIR}" \
       --ts_out="service=grpc-web:${PROTOC_OUT_DIR}" \
       --proto_path=./proto \
       proto/cheers/activity/v1/activity_service.proto \
       proto/cheers/type/party/party.proto \
       proto/cheers/type/privacy/privacy.proto \
       proto/cheers/type/user/user.proto \
       proto/cheers/type/post/post.proto \
       proto/cheers/type/story/story.proto \
       proto/google/type/latlng.proto \
       proto/cheers/party/v1/party_service.proto

