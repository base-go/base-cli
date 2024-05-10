package {{.PackageName}}

import (
	"base-project/core/database"
	"github.com/graphql-go/graphql"
)

// All retrieves all {{.StructName}}s
func All() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList({{.TypeName}}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var items []{{.StructName}}
			if err := database.DB.Find(&items).Error; err != nil {
				return nil, err
			}
			return items, nil
		},
	}
}

// ByID retrieves a single {{.StructName}} by its ID
func ByID() *graphql.Field {
	return &graphql.Field{
		Type: {{.TypeName}},
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, _ := p.Args["id"].(int)
			item := {{.StructName}}{}
			if err := database.DB.First(&item, id).Error; err != nil {
				return nil, err
			}
			return item, nil
		},
	}
}
