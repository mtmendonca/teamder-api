#!/bin/bash

protoc -I grpc grpc/account/account.proto --go_out=plugins=grpc:grpc
protoc -I grpc grpc/event/event.proto --go_out=plugins=grpc:grpc
