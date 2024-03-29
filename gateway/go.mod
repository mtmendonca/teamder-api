module github.com/mtmendonca/teamder-api/gateway

go 1.13

require (
	github.com/auth0/go-jwt-middleware v0.0.0-20190805220309-36081240882b
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/graph-gophers/graphql-go v0.0.0-20190917030536-38a077bc812d
	github.com/mtmendonca/teamder-api/common v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.4.2
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337 // indirect
	github.com/urfave/negroni v1.0.0
)

replace github.com/mtmendonca/teamder-api/common => ../common
