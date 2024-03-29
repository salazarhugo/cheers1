module github.com/salazarhugo/cheers1/services/story-service

go 1.19

require (
	github.com/google/uuid v1.3.0
	github.com/improbable-eng/grpc-web v0.15.0
	github.com/neo4j/neo4j-go-driver/v4 v4.4.4
	github.com/salazarhugo/cheers1/gen/go v0.0.0-20230109095939-243e3b46cf0a
	github.com/salazarhugo/cheers1/libs/auth v0.0.0-20221018182022-91ddac90c15e
	github.com/salazarhugo/cheers1/libs/profiler v0.0.0-20221010151320-33187e7a23f8
	github.com/salazarhugo/cheers1/libs/utils v0.0.0-20221013185914-f468203a7c71
	github.com/sirupsen/logrus v1.9.0
	google.golang.org/grpc v1.50.1
	google.golang.org/protobuf v1.28.1
)

require (
	cloud.google.com/go v0.104.0 // indirect
	cloud.google.com/go/compute v1.10.0 // indirect
	cloud.google.com/go/firestore v1.6.1 // indirect
	cloud.google.com/go/iam v0.5.0 // indirect
	cloud.google.com/go/profiler v0.3.0 // indirect
	cloud.google.com/go/storage v1.27.0 // indirect
	firebase.google.com/go/v4 v4.9.0 // indirect
	github.com/cenkalti/backoff/v4 v4.1.1 // indirect
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/pprof v0.0.0-20220412212628-83db2b799d1f // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.0 // indirect
	github.com/googleapis/gax-go/v2 v2.5.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.12.0 // indirect
	github.com/klauspost/compress v1.11.7 // indirect
	github.com/labstack/echo/v4 v4.9.1 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/rs/cors v1.7.0 // indirect
	github.com/stretchr/testify v1.8.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/crypto v0.0.0-20220926161630-eccd6366d1be // indirect
	golang.org/x/net v0.0.0-20221002022538-bcab6841153b // indirect
	golang.org/x/oauth2 v0.0.0-20221014153046-6fdb5e3db783 // indirect
	golang.org/x/sys v0.0.0-20220928140112-f11e5e49a4ec // indirect
	golang.org/x/text v0.3.8 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/api v0.98.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/appengine/v2 v2.0.2 // indirect
	google.golang.org/genproto v0.0.0-20221014213838-99cd37c6964a // indirect
	nhooyr.io/websocket v1.8.6 // indirect
)

//replace github.com/salazarhugo/cheers1/genproto => ../../genproto

//replace github.com/salazarhugo/cheers1/libs/hello => ../../libs/hello

//replace github.com/salazarhugo/cheers1/libs/utils => ../../libs/utils

//replace github.com/salazarhugo/cheers1/libs/auth => ../../libs/auth
