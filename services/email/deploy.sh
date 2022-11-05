REGION=europe-west2

gcloud run deploy email-service \
  --source . \
  --region=$REGION