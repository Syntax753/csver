# Makefile

all: build run

run:
	# TODO: permission issue to run binary direct
	# go run *.go
	go run api/client.go

build:
	go get -d
	go build -o bin/client api/client.go