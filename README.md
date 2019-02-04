# csver

CSVer reads in csv files and stores the result in a database.
It uses grpc for communcating and has low-memory footprint as it reads line per line

## Project Structure

I've gone with the omni-repos approach with the idea of having a lib/ for common functionality. grpc elements especially would fit well with this structure, and have split the client server implementations under the /cmd folder.

I've create Makefiles to ensure common building terms amongst shared developers as well as taking care of abi generation and binary output.

I would ideally have a master Makefile at project root outside of /cmd to do the end to end demo by taking care of calling the child makes and spinning up the client/server

### Overview:

```text
/cmd/csv-reader
    built to read the csv and store in a structured manner
        /api
            client for grpc
        /bin
            binary output for the client (see TODO below)
        /data
            csv file
        dev.json
            development config - with view of creating qa/prod
            for different envs
        Makefile

/cmd/csv-storer
    built to accept csv records over grpc
        /api
            server for grpc with one service method for now
            to handle inbound storage requests
                /api.proto
                    protobuf service def
                /handler.go
                    server code
        /bin
            binary output for the server
        Makefile
```

## Libraries

All standard apart from:

* github.com/tkanos/gonfig
    A library to avoid the boilerplate in reading in config json files
    Also merges the idea of env vars with config files so easy config 
    on qa and prod servers where it should be all env

## Tools

* go version go1.11.2 linux/amd64

* VSCode with Golang plugins for linters and fmt/goimports on save

* Ubuntu 18.10

## Bugs

* Can't run the binaries from the Makefile so a permission issue I would need to look into

## TODO

* Test cases in particular unit tests for verifiying count of lines consumed by the reader for example. Integration tests would need to mock the grpc side of things which I'd need to look into best practice.

* Benchmarks/checking memory consumption. In particular as it's linear processing I would have wanted to test csv's ReuseRecord option for smaller footprint (though wouldn't eliminate concurrent safety).

* End to end working POC (!) But played it fair with the 2 hour limit and hopefully show that the POC is not far off

* Dockerising the grpc server

* I comment exported methods etc as I go so I feel those are ok, but haven't commented some of the packages

* Exporting the grpc port numbers. The storer doesn't have a notion of config yet

* The csv-reader should initialise the client and used
