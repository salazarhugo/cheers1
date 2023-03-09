gcloud functions deploy delete-user \
  --gen2 \
  --runtime=go119 \
  --region=europe-west2 \
  --source=delete-user \
  --entry-point=DeleteUser \
  --trigger-event providers/firebase.auth/eventTypes/user.delete
