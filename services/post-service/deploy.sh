REGION=europe-west2
GATEWAY_URL=https://android-gateway-clzdlli7.nw.gateway.dev
SPANNER_DSN="projects/cheers-a275e/instances/free-cheers/databases/party"

gcloud run deploy post-service \
  --source . \
  --region=$REGION \
  --use-http2 \
  --set-env-vars GATEWAY_URL=$GATEWAY_URL \
  --set-env-vars SPANNER_DSN=$SPANNER_DSN \
  --set-secrets="DB_ENDPOINT=POST_DB_ENDPOINT:latest" \
  --set-secrets="DB_PASSWORD=POST_DB_PASSWORD:latest" \
  --set-secrets="NEO4J_URI=NEO4J_URI:latest" \
  --set-secrets="NEO4J_PASSWORD=NEO4J_PASSWORD:latest"
