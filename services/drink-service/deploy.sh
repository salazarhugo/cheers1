REGION=europe-west2
SPANNER_DSN="projects/cheers-a275e/instances/free-cheers/databases/party"

gcloud run deploy drink-service \
    --source .  \
    --region=$REGION \
    --set-env-vars SPANNER_DSN=$SPANNER_DSN \
    --use-http2