REGION=europe-west2
EVENT_BUS_TOPIC=chat-topic

gcloud run deploy chat-service \
  --source . \
  --region=$REGION \
  --set-env-vars EVENT_BUS_TOPIC=$EVENT_BUS_TOPIC