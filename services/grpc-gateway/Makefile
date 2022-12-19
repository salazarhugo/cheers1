run:
	cd cmd/app; GOOGLE_APPLICATION_CREDENTIALS=/mnt/c/Users/HugoSalazar/GolandProjects/cheers1/services/grpc-gateway/api-gateway-service-account.json go run main.go cors.go

deploy:
	./deploy.sh

clean:
	rm -rf pkg

update:
	go get github.com/salazarhugo/cheers1/gen/go@main
