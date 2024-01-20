REGION=europe-west2
REDIS_ENDPOINT=redis-18086.c228.us-central1-1.gce.cloud.redislabs.com:18086
SPANNER_DSN="projects/cheers-a275e/instances/free-cheers/databases/party"

gcloud run deploy notification-service \
  --source . \
  --region=$REGION \
  --use-http2 \
  --set-env-vars DB_ENDPOINT=$REDIS_ENDPOINT \
  --set-secrets="DB_PASSWORD=NOTIFICATION_DB_PASSWORD:latest" \
  --set-env-vars SPANNER_DSN=$SPANNER_DSN
