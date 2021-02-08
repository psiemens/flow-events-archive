# Flow Events Archive

The Flow events archive is a simple gRPC proxy service 
that provides access to historical data from previous 
Flow network versions (sporks).

## Supported Methods

This proxy service implements a subset of the Flow Access API.
As of now, only a single method is supported:

- `GetEventsByBlockHeight`

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
