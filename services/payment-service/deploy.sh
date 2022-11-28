REGION=europe-west2

gcloud run deploy payment-service \
    --source .  \
    --region=$REGION \
    --use-http2 \
    --set-secrets="STRIPE_SK=STRIPE_SK:latest" \
    --set-secrets="STRIPE_WH=STRIPE_WH:latest"
