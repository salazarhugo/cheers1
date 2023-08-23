module github.com/salazarhugo/cheers1/services/friendship-service

go 1.19

require (
	github.com/go-redis/redis/v9 v9.0.0-rc.1
	github.com/salazarhugo/cheers1/gen/go v0.0.0-20230814193002-ccda85dd57dc
	github.com/salazarhugo/cheers1/libs/auth v0.0.0-20221116205822-91c0d9c12e5a
	github.com/salazarhugo/cheers1/libs/profiler v0.0.0-20221010151320-33187e7a23f8
	github.com/salazarhugo/cheers1/libs/utils v0.0.0-20230814193002-ccda85dd57dc
	github.com/sirupsen/logrus v1.9.0
	golang.org/x/net v0.1.0
	google.golang.org/grpc v1.50.1
)

require (
	cloud.google.com/go v0.105.0 // indirect
	cloud.google.com/go/compute v1.12.1 // indirect
	cloud.google.com/go/compute/metadata v0.2.1 // indirect
	cloud.google.com/go/firestore v1.6.1 // indirect
	cloud.google.com/go/iam v0.6.0 // indirect
	cloud.google.com/go/profiler v0.3.1 // indirect
	cloud.google.com/go/pubsub v1.3.1 // indirect
	cloud.google.com/go/storage v1.27.0 // indirect
	firebase.google.com/go/v4 v4.9.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/pprof v0.0.0-20221103000818-d260c55eee4c // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.0 // indirect
	github.com/googleapis/gax-go/v2 v2.7.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.12.0 // indirect
	github.com/labstack/echo/v4 v4.9.1 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/neo4j/neo4j-go-driver/v4 v4.4.4 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/oauth2 v0.0.0-20221014153046-6fdb5e3db783 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/api v0.102.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/appengine/v2 v2.0.2 // indirect
	google.golang.org/genproto v0.0.0-20221027153422-115e99e71e1c // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

//replace github.com/salazarhugo/cheers1/gen/go => ../../gen/go

//replace github.com/salazarhugo/cheers1/libs/utils => ../../libs/utils
