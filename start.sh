export GO111MODULE=on
export GOOS=linux
export GOARCH=amd64
# build event service
pushd ./event
go build -o build/main cmd/main.go
popd

# build gateway
pushd ./gateway
go build -o build/main cmd/main.go
popd

# start docker-compose stack
docker-compose up --build
