package types

import (
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

// {{.ModuleNameCapital}} struct
type {{.ModuleNameCapital}} struct {
	gorm.Model
	ID int
	{{range .Fields}} 
	{{.TitledName}} {{.Type}} 
	{{end}}
}

// {{.ModuleNameCapital}}Type for GraphQL schema
var {{.ModuleNameCapital}}Type = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "{{.ModuleNameCapital}}",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			{{range .Fields}} 
			"{{.Name}}": &graphql.Field{
				Type: {{.GqlType}},
			},
			{{end}}
		},
	},
)

// {{.ModuleNameCapital}}Input for   GORM model
type {{.ModuleNameCapital}}Input struct {
	ID      int    `json:"id"`
	{{range .Fields}} 
	{{.TitledName}} {{.Type}} `json:"{{.Name}}"`
	{{end}}
}

type Update{{.ModuleNameCapital}}Input struct {
	ID      int     `json:"id"`
	{{range .Fields}} 
	{{.TitledName}} *{{.Type}} `json:"{{.Name}}"`
	{{end}}
}
