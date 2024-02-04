SPANNER_DSN="projects/cheers-a275e/instances/free-cheers/databases/party"

gcloud functions deploy generate-resized-image \
  --gen2 \
  --runtime=go121 \
  --region=europe-west2 \
  --entry-point=GenerateResizedImage \
  --set-env-vars SPANNER_DSN=$SPANNER_DSN \
  --trigger-bucket=cheers-posts

#  --source=. \
