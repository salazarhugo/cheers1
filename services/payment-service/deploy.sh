REGION=europe-west2

gcloud run deploy payment-service \
    --source .  \
    --region=$REGION