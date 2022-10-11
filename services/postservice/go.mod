module github.com/salazarhugo/cheers1/services/postservice

go 1.19

require (
	github.com/golang/protobuf v1.5.2
	github.com/improbable-eng/grpc-web v0.15.0
	github.com/neo4j/neo4j-go-driver/v4 v4.4.4
	github.com/salazarhugo/cheers1/genproto v0.0.0-20221011193314-8b137ecca37a
	github.com/salazarhugo/cheers1/libs/auth v0.0.0-20221004084225-130de449e1a1
	github.com/salazarhugo/cheers1/libs/profiler v0.0.0-20221010151320-33187e7a23f8
	github.com/sirupsen/logrus v1.7.0
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
)

require (
	cloud.google.com/go v0.104.0 // indirect
	cloud.google.com/go/compute v1.10.0 // indirect
	cloud.google.com/go/firestore v1.6.1 // indirect
	cloud.google.com/go/iam v0.5.0 // indirect
	cloud.google.com/go/profiler v0.3.0 // indirect
	cloud.google.com/go/storage v1.27.0 // indirect
	firebase.google.com/go/v4 v4.8.0 // indirect
	github.com/cenkalti/backoff/v4 v4.1.1 // indirect
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/pprof v0.0.0-20220412212628-83db2b799d1f // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.0 // indirect
	github.com/googleapis/gax-go/v2 v2.5.1 // indirect
	github.com/klauspost/compress v1.11.7 // indirect
	github.com/labstack/echo/v4 v4.9.0 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/rs/cors v1.7.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/crypto v0.0.0-20220926161630-eccd6366d1be // indirect
	golang.org/x/net v0.0.0-20221002022538-bcab6841153b // indirect
	golang.org/x/oauth2 v0.0.0-20220909003341-f21342109be1 // indirect
	golang.org/x/sys v0.0.0-20220928140112-f11e5e49a4ec // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/api v0.98.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/appengine/v2 v2.0.2 // indirect
	google.golang.org/genproto v0.0.0-20220930163606-c98284e70a91 // indirect
	nhooyr.io/websocket v1.8.6 // indirect
)

//replace github.com/salazarhugo/cheers1/libs/hello => ../../libs/hello
//
//replace github.com/salazarhugo/cheers1/libs/utils => ../../libs/utils
//
//replace github.com/salazarhugo/cheers1/libs/auth => ../../libs/auth
