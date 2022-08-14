# go-grpc

### Pre-requisites
- [Golang migrate](https://github.com/golang-migrate/migrate)
- [SQL Boiler](https://github.com/volatiletech/sqlboiler)


## Installation

A little intro about the installation.
```
$ git clone https://github.com/mauryasaurav/go-grpc

Client
$ cd go-grpc/client
$ go mod tidy
$ go run main.go

Server
$ cd go-grpc/server
$ go mod tidy
$ go run main.go
```

### Execute the migrations 
`make migrate-prepare` install dependencies
`make db-migrateup` executes the sql in `db/migrations` migrations to postgres db

### Create the repository layer
`make repository-gen` will generate the api/repository package using sql boiler and provides 
necessary interface for accessing data from db.