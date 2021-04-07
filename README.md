# price-calculator
price-calculator

## How to start 


```shell script
# To generate graphql schemas
# Note: see server/graph/schemas and also ./gqlgen.yml
$ make schema 
```

```shell script
# Use this command to start the service locally

$ make local
```

```shell script
# To run tests

$ make test 
```

```

```shell script
# build a binary for the api

$ make build 
```

```shell script
# generate mocks

$ make gen-mocks 
```

```shell script
# runs go vet

$ make vet 
```

```shell script
# runs go fmt to format files

$ make fmt 
```

### Errors
| Code | ErrorType | Detail |
| ----------- | ----------- | ----------- |
| 101 | InvalidRequestErr | invalid request parameters|
| 102 | InternalErr | Internal Error, you can try again at another time|