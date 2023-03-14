gcloud functions deploy create-user \
  --runtime=go119 \
  --region=europe-west2 \
  --entry-point=CreateUser \
  --trigger-event providers/firebase.auth/eventTypes/user.create

#  --source=. \
