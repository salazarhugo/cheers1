GATEWAY_URL=https://android-gateway-clzdlli7.nw.gateway.dev
TOKEN=eyJhbGciOiJSUzI1NiIsImtpZCI6Ijk5NjJmMDRmZWVkOTU0NWNlMjEzNGFiNTRjZWVmNTgxYWYyNGJhZmYiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vY2hlZXJzLWEyNzVlIiwiYXVkIjoiY2hlZXJzLWEyNzVlIiwiYXV0aF90aW1lIjoxNjYxNzc0ODUzLCJ1c2VyX2lkIjoid29GdHBYVVQ5aVRVSU1Zc3NHNWtJSDVhbmg0MyIsInN1YiI6IndvRnRwWFVUOWlUVUlNWXNzRzVrSUg1YW5oNDMiLCJpYXQiOjE2NjU1ODM3MjgsImV4cCI6MTY2NTU4NzMyOCwiZW1haWwiOiJodWdvYnJvY2s3NCtjaGVlcnNAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsiaHVnb2Jyb2NrNzQrY2hlZXJzQGdtYWlsLmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6InBhc3N3b3JkIn19.IDrfoccCwnvcrxPI0j5hnfnPwd9pZ2nOCPoJBTkvt9KjhoXz9ZHWlxNujPe99duLgQmaFCtzdea_4WfEtScpt_-XsA1-KSEk_9yiUKNyHE4qY08RcfdwPzgg0UpCckn-IIiEHpFetvR4fRkm1EROkhBMLqr-13nS6sXBvgOrpAee0GT5vkn18N8h8BggN1OZXXQrfJmb2yhKRq-yHmy17XIRXOwJj_CMDQ4pMFkvA-srP0UrSNRu-0u74puUNgOf8LlSlCN3CZhupRdwutBCO4XpxR5eZBnfoLJStqlrNTkf-QGnh0cfeQl4Aha7__orWEDiS0Bj5iAIn4GVaYwPSA

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
  "$GATEWAY_URL/v1/posts"
#  -d '
#  {
#    "parent": "",
#    "page_size": 2,
#    "page_token": ""
#  } ' \

#curl --request DELETE \
#  --header "Authorization: Bearer $TOKEN" \
#  -d '
#  {
#    "id": "wfwf"
#  } ' \
#  "$GATEWAY_URL/v1/posts"
