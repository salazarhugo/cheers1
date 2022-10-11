run:
	cd cmd/app; go run main.go

deploy:
	REGION=europe-west2
	GATEWAY_URL=https://android-gateway-clzdlli7.nw.gateway.dev
	gcloud run deploy post-service \
		--source .  \
		--region=$REGION \
		--use-http2 \
		--set-env-vars GATEWAY_URL=$GATEWAY_URL

gen:
	./genproto

clean:
	rm -rf pkg

