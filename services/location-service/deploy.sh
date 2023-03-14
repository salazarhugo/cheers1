REGION=europe-west2

gcloud run deploy location-service \
    --source .  \
    --region=$REGION \
    --use-http2