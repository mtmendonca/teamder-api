module github.com/mtmendonca/teamder-api/gateway

go 1.13

require (
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/mtmendonca/teamder-api/common v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.4.2
	github.com/urfave/negroni v1.0.0
)

replace github.com/mtmendonca/teamder-api/common => ../common
