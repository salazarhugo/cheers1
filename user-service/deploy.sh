REGION=europe-west2
GATEWAY_URL=https://waf-clzdlli7.nw.gateway.dev

gcloud run deploy user-service \
    --source .  \
    --region=$REGION \
    --use-http2 \
    --set-env-vars GATEWAY_URL=$GATEWAY_URL