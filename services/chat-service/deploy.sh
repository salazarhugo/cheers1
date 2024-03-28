REGION=europe-west2
EVENT_BUS_TOPIC=chat-topic
SPANNER_DSN="projects/cheers-a275e/instances/free-cheers/databases/party"

gcloud run deploy chat-service \
  --source . \
  --region=$REGION \
  --set-env-vars SPANNER_DSN=$SPANNER_DSN \
  --set-secrets="DB_ENDPOINT=CHAT_DB_ENDPOINT:latest" \
  --set-secrets="DB_PASSWORD=CHAT_DB_PASSWORD:latest" \
  --set-env-vars EVENT_BUS_TOPIC=$EVENT_BUS_TOPIC \
  --project=cheers-a275e
