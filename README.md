# Fx-web

Implement maintainable golang web applications based on fx, gin, ent, and zap components

# Instructions
## 1. init ent part
```go
cd internal
go run entgo.io/ent/cmd/ent init User

// modify ent schema
go run entgo.io/ent/cmd/ent generate ./ent/schema

// modify table name
func (User) Annotations() []schema.Annotation {
    return []schema.Annotation{
    entsql.Annotation{Table: "user"},
    }
}
```

## 2. init swaggo
```golang
go install github.com/swaggo/swag/cmd/swag@latest

swag init -g ./cmd/main.go -o docs

go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files


```