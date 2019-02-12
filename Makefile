# Makefile

all: build

build:
	make -C cmd/csv-reader build
	make -C cmd/csv-storer build 
	make -C api build 

run:
	make -C cmd/csv-reader run
	make -C cmd/csv-storer run 

demo: build
	cmd/csv-storer/bin/server
	cmd/csv-reader/bin/reader
