REGION=europe-west2
EVENT_BUS_TOPIC=chat-topic
SPANNER_DSN="projects/cheers-a275e/instances/free-cheers/databases/party"

gcloud run deploy chat-service \
  --source . \
  --region=$REGION \
  --set-env-vars SPANNER_DSN=$SPANNER_DSN \
  --set-env-vars EVENT_BUS_TOPIC=$EVENT_BUS_TOPIC