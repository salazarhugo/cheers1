GATEWAY_URL=https://android-gateway-clzdlli7.nw.gateway.dev
TOKEN=eyJhbGciOiJSUzI1NiIsImtpZCI6IjVkMzQwZGRiYzNjNWJhY2M0Y2VlMWZiOWQxNmU5ODM3ZWM2MTYzZWIiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoiSHVnbyBTYWxhemFyIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hLS9BRmRadWNvMWlZck1pRTVMMmRvS1hVaGZKbk1KZzRNU0FIckJnemRQNFRnZjVBPXM5Ni1jIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL2NoZWVycy1hMjc1ZSIsImF1ZCI6ImNoZWVycy1hMjc1ZSIsImF1dGhfdGltZSI6MTY2NDczMDQ0MywidXNlcl9pZCI6IjU1RkV2SGF3aW5RQ2E5amdIN1pkV0VTUjNyaTIiLCJzdWIiOiI1NUZFdkhhd2luUUNhOWpnSDdaZFdFU1IzcmkyIiwiaWF0IjoxNjY1NjkwMjM3LCJleHAiOjE2NjU2OTM4MzcsImVtYWlsIjoiaHVnb2Jyb2NrNzRAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZ29vZ2xlLmNvbSI6WyIxMDY4MTU3Mjk5MDc1OTc0MTQ5NDIiXSwiZW1haWwiOlsiaHVnb2Jyb2NrNzRAZ21haWwuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoiZ29vZ2xlLmNvbSJ9fQ.tKR_pMaoAqAfTW6WphtMbE85R-k9p7o-3l2xYELJLidYf4LVimH4I_2XD-OFY1dvcJuR16-Rwz3_TUVkEy5XSRm2UGEqi3xpYrIePtasVfKuVD81NXZY6A2LRsLAwYHdgTItLe_HROtdNTbRZXrQjuENpAHfjoIW6d_D6doDDM2joSh8sU4hxfpKVJP-gal2M_RPgL-bxyjNbXjkjBghjNaTGPPpfH_AgpRMrIRCv4BsDbIU320eTLYNKj1G9jnARVJSaqC1f1QOywdFg-20iytzrFgoyRRFOUlQUIswY16XF-ZBvO7NLRAVN6Bj7t9wimSjPEiNDs6Qsp0S5vmCJg

#curl -d client_id=$CLIENT_ID -d client_secret=$CLIENT_SECRET -d grant_type=authorization_code -d redirect_uri=urn:ietf:wg:oauth:2.0:oob -d code=4/AABvK4EPc__nckJBK9UGFIhhls_69SBAyidj8J_o3Zz5-VJN6nz54ew https://accounts.google.com/o/oauth2/token

curl --request POST \
  --header "Authorization: Bearer $TOKEN" \
  --http2 \
  -d '
  {
    "party": {
      "id": "android2",
      "caption": "This is my caption",
      "address": "",
      "privacy": "FRIENDS",
      "photos": [""],
      "is_comment_enabled": true,
      "location_name": "",
      "drink": "VODKA",
      "drunkenness": 20,
      "type": "IMAGE",
      "create_time": { "seconds": 0, "nanos": 0 },
      "comments_enabled": true,
      "share_enabled": true
    }
  }
  ' \
  "$GATEWAY_URL/v1/parties" | json_pp

curl --request GET \
  --header "Authorization: Bearer $TOKEN" \
  "$GATEWAY_URL/v1/parties/feed?page_size=2"

#curl --request DELETE \
#  --header "Authorization: Bearer $TOKEN" \
#  -d '
#  {
#    "id": "wfwf"
#  } ' \
#  "$GATEWAY_URL/v1/posts"
