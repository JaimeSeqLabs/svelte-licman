# License Manager - Backend

License Manager implementation using Go as the language

## Project structure
```
lm-backend
|--cmd -> Source code for the main application binaries
|   |--http -> Http server
|--pkg -> Application components
    |--controller -> Request handlers
    |--domain -> Domain entities
    |--repositories -> Persistence and repositories
    |   |--sql -> Implementation with SQL persistence
    |   |--ent -> Implementation with ent framework persistence
    |   |--*_repo.go -> Repository interface
    |--service -> Domain use cases
        |--*.go -> Service interface
        |--*_impl.go -> Service implementation
```


## Ent framework

Generate entity:
```console
go run -mod=mod entgo.io/ent/cmd/ent init User
```

Generate methods for entity after schema definition:
```console
go generate ./pkg/repositories/ent-fw/ent
```

Describe all the entity schemas:
```console
go run -mod=mod entgo.io/ent/cmd/ent describe ./pkg/repositories/ent-fw/ent/schema
```
