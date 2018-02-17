#!/bin/bash

go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go

# generate that shit
protoc ./pb/doggo.proto --go_out=plugins=grpc:.