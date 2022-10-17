#GATEWAY_URL=https://android-gateway-clzdlli7.nw.gateway.dev
GATEWAY_URL=https://party-service-r3a2dr4u4a-nw.a.run.app
TOKEN=eyJhbGciOiJSUzI1NiIsImtpZCI6IjVkMzQwZGRiYzNjNWJhY2M0Y2VlMWZiOWQxNmU5ODM3ZWM2MTYzZWIiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vY2hlZXJzLWEyNzVlIiwiYXVkIjoiY2hlZXJzLWEyNzVlIiwiYXV0aF90aW1lIjoxNjYxNzc0ODUzLCJ1c2VyX2lkIjoid29GdHBYVVQ5aVRVSU1Zc3NHNWtJSDVhbmg0MyIsInN1YiI6IndvRnRwWFVUOWlUVUlNWXNzRzVrSUg1YW5oNDMiLCJpYXQiOjE2NjU3NDg3MjEsImV4cCI6MTY2NTc1MjMyMSwiZW1haWwiOiJodWdvYnJvY2s3NCtjaGVlcnNAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsiaHVnb2Jyb2NrNzQrY2hlZXJzQGdtYWlsLmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6InBhc3N3b3JkIn19.2WEO60x6OYJQ61H9w1-ySTduhYxUZfY5H631ATONN7elC0SLkYCOm93urvXVaBclj8ZjsCxXuDHFyivyqpbwa5hAkud4UIZh_gyXGOtxcoXCGBNrcy2LUp56P9fi5kvj91683XCo0iFRdw0OQiK_MlHastat6V4u7vKc2vKLvwnHhfzYmjDf22cx2AXHXtn_TORIWC1lk0iuzB5ysmQ3UN4j8vFv4lKWqq9GcgrH9Pdr-x0UiyRhVaLK3MO7fP4FID77ItTPOGXwuF6O019ff9p4BN8cpqtTKfJE6gkCfyfeWWJkvfh1slt4Q4Xs7zdvb0wrx6oQLVSeVSvZo6L-rg
#curl -d client_id=$CLIENT_ID -d client_secret=$CLIENT_SECRET -d grant_type=authorization_code -d redirect_uri=urn:ietf:wg:oauth:2.0:oob -d code=4/AABvK4EPc__nckJBK9UGFIhhls_69SBAyidj8J_o3Zz5-VJN6nz54ew https://accounts.google.com/o/oauth2/token

curl --request POST \
  --header "Authorization: Bearer $TOKEN" \
  --http2 \
  -d '
  {
    "party": {
      "id": "android24",
      "caption": "This is my caption",
      "address": "",
      "privacy": "FRIENDS",
      "photos": [],
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
  "$GATEWAY_URL/v1/parties"

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