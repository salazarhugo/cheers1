REGION=europe-west2

gcloud run deploy auth-service \
    --source .  \
    --region=$REGION \
    --use-http2