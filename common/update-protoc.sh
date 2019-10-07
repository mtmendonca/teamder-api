#!/bin/bash

protoc -I grpc grpc/user/user.proto --go_out=plugins=grpc:grpc 