# Fx-web

Implement maintainable and scalable golang web applications based on fx, gin, ent, and zap components

# Instructions
## 1、init ent part
```go
cd internal
go run entgo.io/ent/cmd/ent init User

// 修改对应的ent schema文件
go run entgo.io/ent/cmd/ent generate ./ent/schema

// 修改表名
func (User) Annotations() []schema.Annotation {
    return []schema.Annotation{
    entsql.Annotation{Table: "user"},
    }
}
```

