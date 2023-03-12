gcloud functions deploy recharge-coins \
  --runtime=go119 \
  --region=europe-west2 \
  --entry-point=RechargeCoins \
  --trigger-http
