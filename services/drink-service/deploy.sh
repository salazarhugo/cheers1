REGION=europe-west2

gcloud run deploy drink-service \
    --source .  \
    --region=$REGION \
    --use-http2