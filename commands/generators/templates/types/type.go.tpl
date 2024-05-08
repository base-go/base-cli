package types

import (
    "github.com/graphql-go/graphql"
    "gorm.io/gorm"
)

// {{.ModuleNameCapital}} struct for GORM model
type {{.ModuleNameCapital}} struct {
    gorm.Model
    {{- range .Fields }}
    {{.TitledName}} {{.Type}}
    {{- end }}
}

// {{.ModuleNameCapital}}Type for GraphQL schema
var {{.ModuleNameCapital}}Type = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "{{.ModuleNameCapital}}",
        Fields: graphql.Fields{
            {{- range .Fields }}
            "{{.LowerName}}": &graphql.Field{
                Type: graphql.{{.GqlType}},
                Description: "Field for {{.TitledName}}",
            },
            {{- end }}
        },
    },
)
