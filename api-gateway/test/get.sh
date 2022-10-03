GATEWAY_URL=https://android-gateway-clzdlli7.nw.gateway.dev
TOKEN=eyJhbGciOiJSUzI1NiIsImtpZCI6IjU4NWI5MGI1OWM2YjM2ZDNjOTBkZjBlOTEwNDQ1M2U2MmY4ODdmNzciLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vY2hlZXJzLWEyNzVlIiwiYXVkIjoiY2hlZXJzLWEyNzVlIiwiYXV0aF90aW1lIjoxNjYxNzc0ODUzLCJ1c2VyX2lkIjoid29GdHBYVVQ5aVRVSU1Zc3NHNWtJSDVhbmg0MyIsInN1YiI6IndvRnRwWFVUOWlUVUlNWXNzRzVrSUg1YW5oNDMiLCJpYXQiOjE2NjQ3OTg4MjEsImV4cCI6MTY2NDgwMjQyMSwiZW1haWwiOiJodWdvYnJvY2s3NCtjaGVlcnNAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsiaHVnb2Jyb2NrNzQrY2hlZXJzQGdtYWlsLmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6InBhc3N3b3JkIn19.ADySCJycxprEf-mJKzuv5xO088nFraaU4Z63KBa0kQfLXD504rZAKRnb-ydhStYaVmxof9S0DfAtdhQ2ru9ZId1l2_QWoL4yji2TGEgM2h0QFeWcPKssQiMMQrtQChUe-SNEmvtoLMRsLfTCQcLagDN9lnVwiUP0FkaPXuCZXBkDWDoaJQ8z8Og60rmhJsErj2qpqgQzzW6YoNEjeWrluQYB3dAcgqXya1iFJrSIwlR0mrIqOeGd9bQguKb3NuskcWEYadDc2A6yxW1ek9qZJIkiB6KjoVWcYRNIy4RXhPEQ27gbTRHr2WErHqAGud6nKbaFy-mnCxNjEiOU8IOagQ
#curl -d client_id=$CLIENT_ID -d client_secret=$CLIENT_SECRET -d grant_type=authorization_code -d redirect_uri=urn:ietf:wg:oauth:2.0:oob -d code=4/AABvK4EPc__nckJBK9UGFIhhls_69SBAyidj8J_o3Zz5-VJN6nz54ew https://accounts.google.com/o/oauth2/token
curl --request GET \
  --header "Authorization: Bearer $TOKEN" \
  "$GATEWAY_URL/users/hugosalazar"

curl --request POST \
  --header "Authorization: Bearer $TOKEN" \
  --http2 \
  -d '{"id": "android", "name": "Partyy", "description": "Fw", "address": "", "latlng": { "latitude": 0, "longitude": 0 }, "privacy": "FRIENDS", "banner_url": "", "start_date": { "seconds": 0, "nanos": 0 }, "end_date":  { "seconds": 0, "nanos": 0 }, "host_id": "", "location_name": "", "create_time": { "seconds": 0, "nanos": 0 }  }' \
  "$GATEWAY_URL/parties"


curl --request DELETE \
  --header "Authorization: Bearer $TOKEN" \
  "$GATEWAY_URL/parties"
