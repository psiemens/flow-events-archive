# Flow Events Archive

:warning: _This is an experimental tool.
The core Flow team is exploring other more robust solutions for archival data._

The Flow events archive is a simple gRPC proxy service 
that provides access to historical data from previous 
Flow network versions ([sporks](https://docs.onflow.org/node-operation/spork/)).

## Running Locally

This service is not yet deployed. 
However, you can run the proxy on your local machine and still access the real archive nodes.

```shell
git clone https://github.com/psiemens/flow-events-archive

cd flow-events-archive

go run cmd/archive/main.go

# gRPC (Flow Go SDK): http://localhost:9000
# HTTP (Flow JS SDK): http://localhost:8080
```
## Supported Methods

This proxy service implements a subset of the [Flow Access API](https://docs.onflow.org/access-api).
As of now, only a single method is supported:

- [GetEventsForHeightRange](https://docs.onflow.org/access-api#geteventsforheightrange)

## Available Network Data

| Network | Start Block Height | End Block Height |
|---------|--------------------|------------------|
| Candidate 4 | 1065711 | 2033591 |
| Candidate 5 | 2033592 | 3187930 |
| Candidate 6 | 3187931 | 4132132 |
| Candidate 7 | 4132133 | 4972986 |
| Candidate 8 | 4972987 | 6483245 |
| Candidate 9 | 6483246 | 7601062 |
| Mainnet 1   | 7601063 | 8742958 |
| Mainnet 2   | 8742959 | 9737132 |
| Mainnet 3   | 9737133 | 9992019 |
| Mainnet 4   | 9992020 | 12020336 |
