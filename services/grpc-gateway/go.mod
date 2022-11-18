module github.com/salazarhugo/cheers1/services/grpc-gateway

go 1.19

require (
	github.com/felixge/httpsnoop v1.0.3
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.12.0
	github.com/salazarhugo/cheers1/gen/go v0.0.0-20221114163206-2897b5f5263a
	google.golang.org/grpc v1.50.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220909164309-bea034e7d591 // indirect
	golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10 // indirect
	golang.org/x/text v0.3.8 // indirect
	google.golang.org/genproto v0.0.0-20221014213838-99cd37c6964a // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace github.com/salazarhugo/cheers1/gen/go => ../../gen/go