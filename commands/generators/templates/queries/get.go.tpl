package queries

import (
	"{{.Namespace}}/app/{{.ModuleNameLowerCase}}/types"
	"{{.Namespace}}/core/database"
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
)

// All returns a GraphQL field configuration for getting all {{.ModuleNameLowerCase}}s.
func All() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(types.{{.ModuleNameCapital}}Type),
		Description: "Get all {{.ModuleNameLowerCase}}s",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			fmt.Println("Executing All resolver...")
			var {{.ModuleNameLowerCase}}s []types.{{.ModuleNameCapital}}
			if err := database.DB.Find(&{{.ModuleNameLowerCase}}s).Error; err != nil {
				fmt.Printf("Error fetching {{.ModuleNameLowerCase}}s: %v\n", err)
				return nil, err
			}
			fmt.Printf("{{.ModuleNameCapital}}s retrieved: %d\n", len({{.ModuleNameLowerCase}}s))
			for _, {{.ModuleNameLowerCase}} := range {{.ModuleNameLowerCase}}s {
				fmt.Printf("{{.ModuleNameCapital}}: %v\n", {{.ModuleNameLowerCase}})
			}
			return {{.ModuleNameLowerCase}}s, nil
		},
	}
}

// GetAll{{.ModuleNameCapital}}s retrieves all {{.ModuleNameLowerCase}}s
func ByID() *graphql.Field {
	return &graphql.Field{
		Type:        types.{{.ModuleNameCapital}}Type,
		Description: "Get {{.ModuleNameLowerCase}} by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, ok := params.Args["id"].(int)
			if !ok {
				return nil, errors.New("id must be an integer")
			}
			var {{.ModuleNameLowerCase}} types.{{.ModuleNameCapital}}
			if err := database.DB.First(&{{.ModuleNameLowerCase}}, id).Error; err != nil {
				return nil, err
			}
			return &{{.ModuleNameLowerCase}}, nil
		},
	}
}
