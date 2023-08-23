REGION=europe-west2

gcloud run deploy web-gateway \
  --source . \
  --region=$REGION \
  --service-account api-gateway-management-service@cheers-a275e.iam.gserviceaccount.com