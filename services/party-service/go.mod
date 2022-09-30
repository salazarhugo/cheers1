module github.com/salazarhugo/cheers1/microservices/party-service

go 1.19

require (
	github.com/golang/protobuf v1.5.2
	github.com/salazarhugo/cheers1/libs/auth v0.0.0-20220930103802-98305ab6231a
	github.com/salazarhugo/cheers1/proto/out v0.0.0-20220930103802-98305ab6231a
	google.golang.org/grpc v1.49.0
)

require (
	golang.org/x/net v0.0.0-20220909164309-bea034e7d591 // indirect
	golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220624142145-8cd45d7dbd1f // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

//replace github.com/salazarhugo/cheers1/libs/hello => ../../libs/hello
//
//replace github.com/salazarhugo/cheers1/libs/utils => ../../libs/utils
//
//replace github.com/salazarhugo/cheers1/libs/auth => ../../libs/auth
