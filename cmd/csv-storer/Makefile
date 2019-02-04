# Makefile

all: build run

run:
	# go run *.go
	bin/server

build:
	go get -d
	protoc -I api/ -I/home/syntax/go/src --go_out=plugins=grpc:api api/api.proto
	go build -i -v -o bin/server