GATEWAY_URL=https://android-gateway-clzdlli7.nw.gateway.dev
TOKEN=eyJhbGciOiJSUzI1NiIsImtpZCI6Ijk5NjJmMDRmZWVkOTU0NWNlMjEzNGFiNTRjZWVmNTgxYWYyNGJhZmYiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoiSHVnbyBTYWxhemFyIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hLS9BRmRadWNvMWlZck1pRTVMMmRvS1hVaGZKbk1KZzRNU0FIckJnemRQNFRnZjVBPXM5Ni1jIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL2NoZWVycy1hMjc1ZSIsImF1ZCI6ImNoZWVycy1hMjc1ZSIsImF1dGhfdGltZSI6MTY2NDczMDQ0MywidXNlcl9pZCI6IjU1RkV2SGF3aW5RQ2E5amdIN1pkV0VTUjNyaTIiLCJzdWIiOiI1NUZFdkhhd2luUUNhOWpnSDdaZFdFU1IzcmkyIiwiaWF0IjoxNjY1NTEzNTc4LCJleHAiOjE2NjU1MTcxNzgsImVtYWlsIjoiaHVnb2Jyb2NrNzRAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZ29vZ2xlLmNvbSI6WyIxMDY4MTU3Mjk5MDc1OTc0MTQ5NDIiXSwiZW1haWwiOlsiaHVnb2Jyb2NrNzRAZ21haWwuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoiZ29vZ2xlLmNvbSJ9fQ.OdiArBAa1oJRalja2P0re3jw4y0ksccD6gdd7YK-Ye1c78w0ALzybB2tCo20DlkuNkyXU4HGwtK0Iw5RiwwKcxOV_HtxV4IeenAa7-LfreW2FIoSWNjoKPqbr_r8DP1vubBbN1sb3gXQT0OkgLj9kS0YkkvGnxxY7O2JklbJ3cuMU2Xbr9kGtiQ-EAFXz5vOawuirBktvxzdGspUEWh8UO7trKqpFdKYFq78Sd7Dl3XIck7_B5UwV--ZepuWTVt6siz48xyUShZvFPDjJVhMSIi2CaCYeiP3u_AvYFH1AAf88irtrPKxz7LOeUgSMr8zWIcq836EDPvcC_-Ksk_ZIw
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

#curl --request GET \
#  --header "Authorization: Bearer $TOKEN" \
#  --http2 \
#  -d '
#  {
#    "parent" = "",
#    "page_size" = 20,
#    "page_token" = ""
#  }
#  ' \
#  "$GATEWAY_URL/v1/parties"

curl --request DELETE \
  --header "Authorization: Bearer $TOKEN" \
  "$GATEWAY_URL/v1/posts"
