gcloud api-gateway api-configs create android-api-config-v19 \
    --api=android \
    --project=cheers-a275e \
    --grpc-files=../genproto/api_descriptor.pb,api_config_http.yaml,api_config.yaml