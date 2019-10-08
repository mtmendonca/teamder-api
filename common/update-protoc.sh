#!/bin/bash

protoc -I grpc grpc/account/account.proto --go_out=plugins=grpc:grpc 