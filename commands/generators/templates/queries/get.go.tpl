package queries

import (
    "{{.Namespace}}/app/{{.ModuleNameLowerCase}}/types"
    "{{.Namespace}}/core/database"
    "errors"

    "github.com/graphql-go/graphql"
)

func GetAll{{.ModuleNameCapital}}sField() *graphql.Field {
    return &graphql.Field{
        Type:        graphql.NewList(types.{{.ModuleNameCapital}}Type),
        Description: "Get all {{.ModuleNameLowerCase}}s",
        Resolve: func(params graphql.ResolveParams) (interface{}, error) {
            var {{.ModuleNameLowerCase}}s []types.{{.ModuleNameCapital}}
            if err := database.DB.Find(&{{.ModuleNameLowerCase}}s).Error; err != nil {
                return nil, err // Consider wrapping with a more user-friendly message
            }
            return {{.ModuleNameLowerCase}}s, nil
        },
    }
}

func Get{{.ModuleNameCapital}}ByIDField() *graphql.Field {
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
                return nil, err // Consider wrapping with a more user-friendly message
            }
            return &{{.ModuleNameLowerCase}}, nil
        },
    }
}
