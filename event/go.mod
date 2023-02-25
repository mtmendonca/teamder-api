module github.com/mtmendonca/teamder-api/event

go 1.13

require (
	github.com/google/go-cmp v0.3.0 // indirect
	github.com/mtmendonca/teamder-api/common v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.4.2
	go.mongodb.org/mongo-driver v1.1.2
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c // indirect
	golang.org/x/sys v0.1.0 // indirect
	google.golang.org/genproto v0.0.0-20190502173448-54afdca5d873 // indirect
	google.golang.org/grpc v1.24.0
)

replace github.com/mtmendonca/teamder-api/common => ../common
