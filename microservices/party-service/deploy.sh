REGION=europe-west2
GATEWAY_DOMAIN=android-gateway-clzdlli7.nw.gateway.dev
TOKEN=eyJhbGciOiJSUzI1NiIsImtpZCI6IjJkMjNmMzc0MDI1ZWQzNTNmOTg0YjUxMWE3Y2NlNDlhMzFkMzFiZDIiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vY2hlZXJzLWEyNzVlIiwiYXVkIjoiY2hlZXJzLWEyNzVlIiwiYXV0aF90aW1lIjoxNjYxNzc0ODUzLCJ1c2VyX2lkIjoid29GdHBYVVQ5aVRVSU1Zc3NHNWtJSDVhbmg0MyIsInN1YiI6IndvRnRwWFVUOWlUVUlNWXNzRzVrSUg1YW5oNDMiLCJpYXQiOjE2NjMyNTIwOTUsImV4cCI6MTY2MzI1NTY5NSwiZW1haWwiOiJodWdvYnJvY2s3NCtjaGVlcnNAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsiaHVnb2Jyb2NrNzQrY2hlZXJzQGdtYWlsLmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6InBhc3N3b3JkIn19.UbP_iijY6b5_EFqLpn2mZTWTXBkV1LDK2xFXfVT3CHW9Ant0ceSRWlt31tYTbXcrkcyRCBbE_8-xu7we-Gaa26vbCmBC2JV-FRJTWzkLYma09SfTTBrhZ52iCrmi9Nw4hehyhdsO7lwGpaztm3IRKCn0cQKxFVsVE6EEMM5ayGpvo0qPSS5pKxh4QYbtfP3xEz6mVQGhwJYWexTyM6SireR1xccOk3Hjc2zWFDL6MLkgTRpdw_N2zcoxYm07zkLtwvDMClqArZcFAuzjmiiyaQKGJSoYCJh8gm0FH5zQ2rV0nqvQHvw0T_JhDV1_53twtkbNRYh5PjiSmzhDeOt1UQ

gcloud run deploy party-service \
    --source .  \
    --region=$REGION \
    --set-env-vars GATEWAY_DOMAIN=$GATEWAY_DOMAIN,TOKEN=$TOKEN