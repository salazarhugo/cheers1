REGION=europe-west2

gcloud run deploy order-service \
    --source .  \
    --region=$REGION \
    --use-http2
