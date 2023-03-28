REGION=europe-west2

gcloud run deploy note-service \
    --source .  \
    --region=$REGION \
    --use-http2