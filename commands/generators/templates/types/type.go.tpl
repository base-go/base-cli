package {{.PackageName}}

import "github.com/graphql-go/graphql"

// {{.StructName}} defines the structure for the {{.StructName}} model
type {{.StructName}} struct {
	ID    uint   `json:"id"`
	{{range .Fields}} {{.Name}} {{.Type}} `json:"{{.JsonName}}"`{{end}}
}

// {{.TypeName}} defines the GraphQL type for {{.StructName}}
var {{.TypeName}} = graphql.NewObject(graphql.ObjectConfig{
	Name: "{{.StructName}}",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		{{range .Fields}} "{{.Name}}": &graphql.Field{
			Type: graphql.{{.GraphQLType}},
		},
		{{end}}
	},
})
