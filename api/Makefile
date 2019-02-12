# Makefile

GOPATH=$(shell go env GOPATH)

all: build

build:
	@GOPATH=$(GOPATH) protoc -I . -I/${GOPATH}/src --go_out=plugins=grpc:. api.proto
