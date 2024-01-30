REGION=europe-west2

gcloud run deploy web-gateway \
  --source . \
  --region=$REGION