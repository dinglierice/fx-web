# Fx-web

一个用golang实现的web应用demo，fx、ent、zap

# Instructions
## 1、ent部分初始化
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

