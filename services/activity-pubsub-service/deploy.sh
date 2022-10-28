REGION=europe-west2
GATEWAY_URL=https://android-gateway-clzdlli7.nw.gateway.dev
POST_SERVICE_HOST=https://post-service-r3a2dr4u4a-nw.a.run.app

gcloud run deploy activity-pubsub-service \
  --source . \
  --region=$REGION \
  --set-env-vars GATEWAY_URL=$GATEWAY_URL, POST_SERVICE_HOST=$POST_SERVICE_HOST