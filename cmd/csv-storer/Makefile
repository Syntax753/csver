# Makefile

all: build run

run:
	go run *.go

build:
	go get -d
	go build -o csvreader
	protoc -I api/ -I/home/syntax/go/src --go_out=plugins=grpc:api api/api.proto