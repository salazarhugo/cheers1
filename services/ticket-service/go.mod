module github.com/salazarhugo/cheers1/services/ticket-service

go 1.19

require (
	cloud.google.com/go/firestore v1.8.0
	github.com/salazarhugo/cheers1/gen/go v0.0.0-20221209142646-301101629442
	github.com/salazarhugo/cheers1/libs/auth v0.0.0-20221113151338-1ae307745902
	github.com/salazarhugo/cheers1/libs/profiler v0.0.0-20221109222439-50f4ec2ba1bb
	github.com/salazarhugo/cheers1/libs/utils v0.0.0-20221208094612-0cd8a5fa5c43
	github.com/sirupsen/logrus v1.9.0
	github.com/stripe/stripe-go/v72 v72.122.0
	golang.org/x/net v0.2.0
	google.golang.org/api v0.103.0
	google.golang.org/grpc v1.50.1
)

require (
	cloud.google.com/go v0.106.0 // indirect
	cloud.google.com/go/compute v1.12.1 // indirect
	cloud.google.com/go/compute/metadata v0.2.1 // indirect
	cloud.google.com/go/iam v0.7.0 // indirect
	cloud.google.com/go/longrunning v0.3.0 // indirect
	cloud.google.com/go/profiler v0.3.0 // indirect
	cloud.google.com/go/pubsub v1.26.0 // indirect
	cloud.google.com/go/storage v1.28.0 // indirect
	firebase.google.com/go/v4 v4.9.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/pprof v0.0.0-20220412212628-83db2b799d1f // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.0 // indirect
	github.com/googleapis/gax-go/v2 v2.7.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.12.0 // indirect
	github.com/labstack/echo/v4 v4.9.1 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/neo4j/neo4j-go-driver/v4 v4.4.4 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/crypto v0.2.0 // indirect
	golang.org/x/oauth2 v0.2.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.2.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	golang.org/x/time v0.2.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/appengine/v2 v2.0.2 // indirect
	google.golang.org/genproto v0.0.0-20221109142239-94d6d90a7d66 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

//replace github.com/salazarhugo/cheers1/gen/go => ../../gen/go

//replace github.com/salazarhugo/cheers1/libs/utils => ../../libs/utils

//replace github.com/salazarhugo/cheers1/libs/auth => ../../libs/auth
