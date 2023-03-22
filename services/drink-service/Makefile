build:
	go build cmd/app/main.go

run:
	go run cmd/app/main.go

test: build
	go test

deploy:
	./deploy.sh

clean:
	rm -rf pkg

update:
	go get github.com/salazarhugo/cheers1/gen/go@main