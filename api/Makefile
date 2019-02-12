# Makefile

all: build

build:
	protoc -I api/ -I/home/syntax/go/src --go_out=plugins=grpc:api api/api.proto