REGION=europe-west2

gcloud run deploy comment-service \
    --source .  \
    --region=$REGION \
    --use-http2 \
    --set-secrets="DB_ENDPOINT=COMMENT_DB_ENDPOINT:latest" \
    --set-secrets="DB_PASSWORD=COMMENT_DB_PASSWORD:latest"
