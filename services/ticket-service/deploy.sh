REGION=europe-west2

gcloud run deploy ticket-service \
    --source .  \
    --region=$REGION \
    --use-http2
