GATEWAY_URL=https://android-gateway-clzdlli7.nw.gateway.dev
TOKEN=eyJhbGciOiJSUzI1NiIsImtpZCI6Ijk5NjJmMDRmZWVkOTU0NWNlMjEzNGFiNTRjZWVmNTgxYWYyNGJhZmYiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoiSHVnbyBTYWxhemFyIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hLS9BRmRadWNvMWlZck1pRTVMMmRvS1hVaGZKbk1KZzRNU0FIckJnemRQNFRnZjVBPXM5Ni1jIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL2NoZWVycy1hMjc1ZSIsImF1ZCI6ImNoZWVycy1hMjc1ZSIsImF1dGhfdGltZSI6MTY2NDczMDQ0MywidXNlcl9pZCI6IjU1RkV2SGF3aW5RQ2E5amdIN1pkV0VTUjNyaTIiLCJzdWIiOiI1NUZFdkhhd2luUUNhOWpnSDdaZFdFU1IzcmkyIiwiaWF0IjoxNjY1NTk5NzI1LCJleHAiOjE2NjU2MDMzMjUsImVtYWlsIjoiaHVnb2Jyb2NrNzRAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZ29vZ2xlLmNvbSI6WyIxMDY4MTU3Mjk5MDc1OTc0MTQ5NDIiXSwiZW1haWwiOlsiaHVnb2Jyb2NrNzRAZ21haWwuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoiZ29vZ2xlLmNvbSJ9fQ.YMzr1q9WzxS-KngdJx3xlSQByJIcgL8PoLKfMw7dfET8VBccWZbQZJgvscS2qOo_8sbxczUzXYp1N42NYEdNMQ1-zHrsc88DJQyvb0_53-lmUwizMuurgAkb8-2_WBM5KnMpoi7FmyTKSBiDu8QM-v6psEcHvPB-8bJwQn0QC-LOLojOInXZ6iRNYAflKs3gMQAIQ3AZR6AuZds0FN7WIb3mRFlC-xKqVYfYkY7e-ZWFuhtGFPv7QEJz9tnThiVQP6Kxj72pCc7xopiZndP0UW2vOHt-IFhhDrfgZqMtL7BeaEQlogDPymjD5kzSD__0v3MMy3_9-OzC0BoohC7xow
#curl -d client_id=$CLIENT_ID -d client_secret=$CLIENT_SECRET -d grant_type=authorization_code -d redirect_uri=urn:ietf:wg:oauth:2.0:oob -d code=4/AABvK4EPc__nckJBK9UGFIhhls_69SBAyidj8J_o3Zz5-VJN6nz54ew https://accounts.google.com/o/oauth2/token
#curl --request POST \
#  --header "Authorization: Bearer $TOKEN" \
#  --http2 \
#  -d '
#  {
#    "party": {
#      "id": "android",
#      "name": "Partyy",
#      "description": "Fw",
#      "address": "",
#      "latlng": { "latitude": 0, "longitude": 0 },
#       "privacy": "FRIENDS",
#      "banner_url": "",
#       "start_date": { "seconds": 0, "nanos": 0 },
#      "end_date":  { "seconds": 0, "nanos": 0 },
#      "host_id": "",
#      "location_name": "",
#      "create_time": { "seconds": 0, "nanos": 0 }  }
#    }
#  }
#  ' \
#  "$GATEWAY_URL/v1/parties"

curl --request GET \
  --header "Authorization: Bearer $TOKEN" \
  "$GATEWAY_URL/v1/posts?page_size=2"

#curl --request DELETE \
#  --header "Authorization: Bearer $TOKEN" \
#  -d '
#  {
#    "id": "wfwf"
#  } ' \
#  "$GATEWAY_URL/v1/posts"
