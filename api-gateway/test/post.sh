GATEWAY_URL=https://android-gateway-clzdlli7.nw.gateway.dev
TOKEN=eyJhbGciOiJSUzI1NiIsImtpZCI6IjVkMzQwZGRiYzNjNWJhY2M0Y2VlMWZiOWQxNmU5ODM3ZWM2MTYzZWIiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoiSHVnbyBTYWxhemFyIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hLS9BRmRadWNvMWlZck1pRTVMMmRvS1hVaGZKbk1KZzRNU0FIckJnemRQNFRnZjVBPXM5Ni1jIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL2NoZWVycy1hMjc1ZSIsImF1ZCI6ImNoZWVycy1hMjc1ZSIsImF1dGhfdGltZSI6MTY2NDczMDQ0MywidXNlcl9pZCI6IjU1RkV2SGF3aW5RQ2E5amdIN1pkV0VTUjNyaTIiLCJzdWIiOiI1NUZFdkhhd2luUUNhOWpnSDdaZFdFU1IzcmkyIiwiaWF0IjoxNjY1ODQ3NDA4LCJleHAiOjE2NjU4NTEwMDgsImVtYWlsIjoiaHVnb2Jyb2NrNzRAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZ29vZ2xlLmNvbSI6WyIxMDY4MTU3Mjk5MDc1OTc0MTQ5NDIiXSwiZW1haWwiOlsiaHVnb2Jyb2NrNzRAZ21haWwuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoiZ29vZ2xlLmNvbSJ9fQ.CfktCdTjYkZqunpWOse9Nf_iL69BF1nqEKTw6wP3OFlsp0gDQM-mmjH5xmDau-a9_2jXOUXboISxzYQfBB1tIo6ooMinpVSkJqB3zFqQGcCPWsCbDOQFN0co3cGKSdglFt8e5lWeAOCQBiNIAO-VHpLpwbrRCFowJyw9ExAIv5SJDZerBB4oFgmpYTfcyBfI8GvSQNw7mDqtgzCQU7ZcoLhzflNeR5WGMSe8chyZCEis-5yAPmXUijhXnROnvnItxVQvmiTHS8oQRMRcyXLm5iWbGS8fw8eYZJlluaC5BVp8OSfTSP7xoybmde-4MjqUE4jqHenrEX9bB8jLkcBFLA
#curl --request POST \
#  --header "Authorization: Bearer $TOKEN" \
#  --http2 \
#  -d '
#  {
#    "post": {
#      "id": "android2",
#      "caption": "This is my caption",
#      "address": "",
#      "privacy": "FRIENDS",
#      "photos": [""],
#      "is_comment_enabled": true,
#      "location_name": "",
#      "drink": "VODKA",
#      "drunkenness": 20,
#      "type": "IMAGE",
#      "create_time": { "seconds": 0, "nanos": 0 },
#      "comments_enabled": true,
#      "share_enabled": true
#    }
#  }
#  ' \
#  "$GATEWAY_URL/v1/posts" | json_pp

curl --request GET \
  --header "Authorization: Bearer $TOKEN" \
  "$GATEWAY_URL/v1/posts?page_size=20" | json_pp
