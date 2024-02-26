REGION=europe-west2
GATEWAY_URL=https://android-gateway-clzdlli7.nw.gateway.dev
SPANNER_DSN="projects/cheers-a275e/instances/free-cheers/databases/party"

gcloud run deploy note-service \
    --source .  \
    --region=$REGION \
    --use-http2 \
    --set-env-vars GATEWAY_URL=$GATEWAY_URL \
    --set-env-vars SPANNER_DSN=$SPANNER_DSN
