module github.com/salazarhugo/cheers1/services/party-service

go 1.19

require (
	cloud.google.com/go/datastore v1.12.0
	github.com/google/uuid v1.3.0
	github.com/neo4j/neo4j-go-driver/v4 v4.4.4
	github.com/salazarhugo/cheers1/gen/go v0.0.0-20230814193002-ccda85dd57dc
	github.com/salazarhugo/cheers1/libs/auth v0.0.0-20221113151338-1ae307745902
	github.com/salazarhugo/cheers1/libs/profiler v0.0.0-20221010151320-33187e7a23f8
	github.com/salazarhugo/cheers1/libs/utils v0.0.0-20230814193002-ccda85dd57dc
	github.com/sirupsen/logrus v1.9.0
	golang.org/x/net v0.12.0
	google.golang.org/grpc v1.56.2
)

require (
	cloud.google.com/go v0.110.4 // indirect
	cloud.google.com/go/compute v1.20.1 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/firestore v1.11.0 // indirect
	cloud.google.com/go/iam v1.1.0 // indirect
	cloud.google.com/go/longrunning v0.5.1 // indirect
	cloud.google.com/go/profiler v0.3.1 // indirect
	cloud.google.com/go/pubsub v1.32.0 // indirect
	cloud.google.com/go/storage v1.32.0 // indirect
	firebase.google.com/go/v4 v4.9.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/pprof v0.0.0-20221103000818-d260c55eee4c // indirect
	github.com/google/s2a-go v0.1.4 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.5 // indirect
	github.com/googleapis/gax-go/v2 v2.12.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.12.0 // indirect
	github.com/labstack/echo/v4 v4.9.1 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/oauth2 v0.10.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/api v0.132.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/appengine/v2 v2.0.2 // indirect
	google.golang.org/genproto v0.0.0-20230706204954-ccb25ca9f130 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230706204954-ccb25ca9f130 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230711160842-782d3b101e98 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace github.com/salazarhugo/cheers1/gen/go => ../../gen/go

//replace github.com/salazarhugo/cheers1/libs/utils => ../../libs/utils

//replace github.com/salazarhugo/cheers1/libs/auth => ../../libs/auth
