GATEWAY_URL=https://android-gateway-clzdlli7.nw.gateway.dev
TOKEN=eyJhbGciOiJSUzI1NiIsImtpZCI6Ijk5NjJmMDRmZWVkOTU0NWNlMjEzNGFiNTRjZWVmNTgxYWYyNGJhZmYiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vY2hlZXJzLWEyNzVlIiwiYXVkIjoiY2hlZXJzLWEyNzVlIiwiYXV0aF90aW1lIjoxNjYxNzc0ODUzLCJ1c2VyX2lkIjoid29GdHBYVVQ5aVRVSU1Zc3NHNWtJSDVhbmg0MyIsInN1YiI6IndvRnRwWFVUOWlUVUlNWXNzRzVrSUg1YW5oNDMiLCJpYXQiOjE2NjU2NjQ0MDcsImV4cCI6MTY2NTY2ODAwNywiZW1haWwiOiJodWdvYnJvY2s3NCtjaGVlcnNAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsiaHVnb2Jyb2NrNzQrY2hlZXJzQGdtYWlsLmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6InBhc3N3b3JkIn19.qBNXKPSpzvltltOWKo-Ck4YpTeX5Lt5SOjYPMW9HIt6K1xZaojDAzX6lfjXtgak1i5jWTbbMqLmCv8-gJ1tkH1ImvYe_EZpP-W4BMbkQUGGQymLYL4pc0B8DhYSTfl8QGAuo2J2q3SuCHpTtWS6sBapk271o4LN2le2cUdGG-4XdqUCGFRILvE-TEC3uiCjk2Xc43jreKkIEXKwb-FfRrs-52bCpmhjr-PaFB8XefQV6LBP23P5Q_r90Vwt7c-c15SmgvLAcTj34WKS3l-gOLFmOu9p_2ArerAzlCvqjBT83dkd06jXhcn_ceuLy9pVK9I4csqG9uyNFaYEQZQEMgw

#curl -d client_id=$CLIENT_ID -d client_secret=$CLIENT_SECRET -d grant_type=authorization_code -d redirect_uri=urn:ietf:wg:oauth:2.0:oob -d code=4/AABvK4EPc__nckJBK9UGFIhhls_69SBAyidj8J_o3Zz5-VJN6nz54ew https://accounts.google.com/o/oauth2/token

curl --request POST \
  --header "Authorization: Bearer $TOKEN" \
  --http2 \
  -d '
  {
    "post": {
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
  "$GATEWAY_URL/v1/posts" | json_pp

curl --request GET \
  --header "Authorization: Bearer $TOKEN" \
  "$GATEWAY_URL/v1/posts?page_size=2" | json_pp

#curl --request DELETE \
#  --header "Authorization: Bearer $TOKEN" \
#  -d '
#  {
#    "id": "wfwf"
#  } ' \
#  "$GATEWAY_URL/v1/posts"
