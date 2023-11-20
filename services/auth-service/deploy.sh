REGION=europe-west2
SPANNER_DSN="projects/cheers-a275e/instances/free-cheers/databases/party"


gcloud run deploy auth-service \
    --source .  \
    --region=$REGION \
    --use-http2 \
    --set-secrets="DB_ENDPOINT=FRIENDSHIP_DB_ENDPOINT:latest" \
    --set-secrets="DB_PASSWORD=FRIENDSHIP_DB_PASSWORD:latest" \
    --set-env-vars SPANNER_DSN=$SPANNER_DSN
