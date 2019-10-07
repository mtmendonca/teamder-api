module github.com/mtmendonca/teamder-api/account

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/mtmendonca/teamder-api/common v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.4.2
	go.mongodb.org/mongo-driver v1.1.2
	google.golang.org/api v0.11.0
	google.golang.org/grpc v1.24.0
)

replace github.com/mtmendonca/teamder-api/common => ../common
