REGION=europe-west2

gcloud run deploy account-service \
    --source .  \
    --region=$REGION \
    --use-http2