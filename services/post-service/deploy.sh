REGION=europe-west2
GATEWAY_URL=https://android-gateway-clzdlli7.nw.gateway.dev

gcloud run deploy post-service \
  --source . \
  --region=$REGION \
  --use-http2 \
  --set-env-vars GATEWAY_URL=$GATEWAY_URL \
  --set-secrets="DB_ENDPOINT=POST_DB_ENDPOINT:latest" \
  --set-secrets="DB_PASSWORD=POST_DB_PASSWORD:latest"
