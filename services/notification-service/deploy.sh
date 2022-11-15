REGION=europe-west2

gcloud run deploy notification-service \
  --source . \
  --region=$REGION \
  --use-http2
