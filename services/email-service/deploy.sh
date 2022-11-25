REGION=europe-west2

gcloud beta run deploy email-service \
  --source . \
  --region=$REGION \
  --update-secrets="EMAIL_PASSWORD=EMAIL_PASSWORD:latest"