gcloud functions deploy delete-user \
  --runtime=go119 \
  --region=europe-west2 \
  --source=delete-user \
  --entry-point=DeleteAuth \
  --trigger-event providers/firebase.auth/eventTypes/user.delete
