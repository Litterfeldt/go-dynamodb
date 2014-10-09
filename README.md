go-dynamodb
===========

Simple Interface lib for Amazon AWS DynamoDB, written in golang.
Currently works as a Key/Value store with a second index, at the moment only supports strings as keys and values.

## Example usage
Check `live_test.go`

## Setup
The following ENV variables needs to be set for this to run:
* `AWS_SECRET_ACCESS_KEY`
* `AWS_ACCESS_KEY_ID`
* `DYNAMODB_HOST` example: `dynamodb.eu-west-1.amazonaws.com`
* `DYNAMODB_ZONE` example: `eu-west-1`

then to install:
```bash
go get github.com/Litterfeldt/go-dynamodb`
```


