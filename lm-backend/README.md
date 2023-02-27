# License Manager - Backend

License Manager implementation using Go as the language

## Project structure
```
lm-backend
|--cmd -> Source code for the main application binaries
|   |--http -> Http server
|--pkg -> Application components
|   |--config -> App configuration providers
|   |--controller -> Request handlers
|   |   |--exchange -> Request/Response structures
|   |--domain -> Domain entities
|   |--pkgerror -> Custom errors
|   |--repositories -> Persistence and repositories
|   |   |--sql -> SQL implementation
|   |   |--ent -> Ent framework implementation
|   |   |--*_repo.go -> Repository generic interface
|   |--service -> Domain use cases
|       |--*.go -> Service interface
|       |--*_impl.go -> Service implementation
|--dev.env -> Configuration file
```

## Run server
To run the backend server you need an `app.env` configuration file in the same dir in which you are running the server, you can rename/copy `dev.env` to get started.

```console
go run cmd/http/main.go
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

Generate schema diagram:
```console
go run -mod=mod github.com/hedwigz/entviz/cmd/entviz ./pkg/repositories/ent-fw/ent/schema
```
