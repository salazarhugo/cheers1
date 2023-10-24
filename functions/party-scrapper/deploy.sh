REGION=europe-west2

TOTAL_PAGE=10
SCRAPE_URL=https://shotgun.live/paris
GATEWAY_URL=https://web-gateway-r3a2dr4u4a-nw.a.run.app
PROXY_USERNAME=exologo
PROXY_PASSWORD=AqwXsz123987
PROXY_HOST=geo.iproyal.com
PROXY_PORT=32325

gcloud functions deploy scrape-parties \
  --gen2 \
  --runtime=python311 \
  --region=$REGION \
  --source=. \
  --entry-point=scrape_parties \
  --trigger-http \
  --set-env-vars TOTAL_PAGE=$TOTAL_PAGE \
  --set-env-vars SCRAPE_URL=$SCRAPE_URL \
  --set-env-vars GATEWAY_URL=$GATEWAY_URL \
  --set-env-vars PROXY_USERNAME=$PROXY_USERNAME \
  --set-env-vars PROXY_PASSWORD=$PROXY_PASSWORD \
  --set-env-vars PROXY_HOST=$PROXY_HOST \
  --set-env-vars PROXY_PORT=$PROXY_PORT

