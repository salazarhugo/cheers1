REGION=europe-west2

gcloud run deploy order-service \
    --source .  \
    --region=$REGION \
    --use-http2 \
    --set-secrets="STRIPE_SK=STRIPE_SK:latest"
