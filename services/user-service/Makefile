build: test
	go build cmd/app/main.go

run:
	go run cmd/app/main.go

deploy:
	./deploy.sh

test:
	go test -v -cover ./...

clean:
	rm -rf pkg

update:
	go get github.com/salazarhugo/cheers1/gen/go@main