REGION=europe-west2

gcloud beta run deploy web-gateway \
  --source . \
  --region=$REGION