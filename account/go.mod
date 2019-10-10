module github.com/mtmendonca/teamder-api/account

go 1.13

require (
	github.com/mtmendonca/teamder-api/common/grpc/account v0.0.0-00010101000000-000000000000
	github.com/mtmendonca/teamder-api/common/mongodb v0.0.0-00010101000000-000000000000
	github.com/mtmendonca/teamder-api/common/types v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.1.2
	google.golang.org/grpc v1.24.0
)

replace github.com/mtmendonca/teamder-api/common/types => ../common/types

replace github.com/mtmendonca/teamder-api/common/grpc/account => ../common/grpc/account

replace github.com/mtmendonca/teamder-api/common/mongodb => ../common/mongodb
