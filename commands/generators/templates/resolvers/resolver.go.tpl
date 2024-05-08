package resolvers

import (
    "{{.Namespace}}/app/{{.ModuleNameLowerCase}}/types"
    "{{.Namespace}}/core/database"

    "github.com/graphql-go/graphql"
)

// {{.ModuleNameCapital}}Resolver resolves queries and mutations related to {{.ModuleNameLowerCase}}s
type {{.ModuleNameCapital}}Resolver struct{}

// GetAll{{.ModuleNameCapital}}s resolves the query for fetching all {{.ModuleNameLowerCase}}s
func (r *{{.ModuleNameCapital}}Resolver) GetAll{{.ModuleNameCapital}}s(params graphql.ResolveParams) (interface{}, error) {
    var {{.ModuleNameLowerCase}}s []types.{{.ModuleNameCapital}}
    if err := database.DB.Find(&{{.ModuleNameLowerCase}}s).Error; err != nil {
        return nil, err // Consider wrapping with a more user-friendly message
    }
    return {{.ModuleNameLowerCase}}s, nil
}
