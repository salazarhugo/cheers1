REGION=europe-west2

gcloud beta run deploy email-service \
  --source . \
  --region=$REGION \
  --update-secrets="SENDGRID_WEB_API_KEY=SENDGRID_WEB_API_KEY:latest"
