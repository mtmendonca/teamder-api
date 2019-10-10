module github.com/mtmendonca/teamder-api/gateway

go 1.13

require (
	github.com/auth0/go-jwt-middleware v0.0.0-20190805220309-36081240882b
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/mux v1.7.3
	github.com/graph-gophers/graphql-go v0.0.0-20190917030536-38a077bc812d
	github.com/mtmendonca/teamder-api/common/grpc v0.0.0-00010101000000-000000000000
	github.com/mtmendonca/teamder-api/common/grpc/account v0.0.0-00010101000000-000000000000
	github.com/mtmendonca/teamder-api/common/types v0.0.0-00010101000000-000000000000
	github.com/urfave/negroni v1.0.0
	google.golang.org/grpc v1.24.0
)

replace github.com/mtmendonca/teamder-api/common/types => ../common/types

replace github.com/mtmendonca/teamder-api/common/grpc => ../common/grpc

replace github.com/mtmendonca/teamder-api/common/grpc/account => ../common/grpc/account

replace github.com/mtmendonca/teamder-api/common/mongodb => ../common/mongodb
