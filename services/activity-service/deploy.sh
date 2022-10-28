REGION=europe-west2
GATEWAY_URL=https://android-gateway-clzdlli7.nw.gateway.dev

gcloud run deploy activity-service \
  --source . \
  --region=$REGION \
  --set-env-vars GATEWAY_URL=$GATEWAY_URL