REGION=europe-west2

gcloud run deploy location-service \
    --source .  \
    --region=$REGION \
    --use-http2 \
    --set-secrets="DB_ENDPOINT=LOCATION_DB_ENDPOINT:latest" \
    --set-secrets="DB_PASSWORD=LOCATION_DB_PASSWORD:latest"
