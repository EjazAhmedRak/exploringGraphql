# ExploringGraphql
This is just a practise project for exploring GraphQl with Go

## Pre-requisites
* Create a schema in mysql with the schema name explore_graphql
* Run the commands :
` 
    `go build -o bin/dbInitialize cmd/dbInitialize.go`
    `./bin/dbInitialize`
`
## Starting the Application

* Run the commands :
` 
    `go build -o bin/goGraphQl cmd/goGraphQl.go`
    `./bin/goGraphQl`
`
* The application is served at localhost:8080
* For testing via POSTMAN: `http://localhost:8080/graphql?query={user(id:"b999b167-e91b-48f9-9182-53e6dc65f2de"){name, uuid, email}}`

## Running Tests

* Run the commands : 
    `cd test ; go test run Query ; cd ..`

## Libraries Used:

* graphql-go [https://github.com/graphql-go/graphql]
* gorm [https://github.com/go-gorm/gorm]
* uuid [https://github.com/satori/go.uuid]

