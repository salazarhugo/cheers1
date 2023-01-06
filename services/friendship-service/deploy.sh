REGION=europe-west2

gcloud run deploy friendship-service \
    --source .  \
    --region=$REGION \
    --use-http2 \
    --set-secrets="DB_ENDPOINT=FRIENDSHIP_DB_ENDPOINT:latest" \
    --set-secrets="DB_PASSWORD=FRIENDSHIP_DB_PASSWORD:latest"
