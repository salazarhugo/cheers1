REGION=europe-west2
SPANNER_DSN="projects/cheers-a275e/instances/free-cheers/databases/party"


gcloud run deploy auth-service \
    --source .  \
    --region=$REGION \
    --use-http2 \
    --set-env-vars SPANNER_DSN=$SPANNER_DSN
