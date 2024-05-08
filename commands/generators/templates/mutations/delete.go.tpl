package mutations

import (
    "{{.Namespace}}/app/{{.ModuleNameLowerCase}}/types"
    "{{.Namespace}}/core/database"
    "log"

    "github.com/graphql-go/graphql"
)

func Delete{{.ModuleNameCapital}}(id int) (string, error) {
    var {{.ModuleNameLowerCase}} types.{{.ModuleNameCapital}}
    // Delete the record with the provided ID
    if err := database.DB.Delete(&{{.ModuleNameLowerCase}}, id).Error; err != nil {
        log.Printf("Error deleting {{.ModuleNameLowerCase}} with ID %d: %v", id, err)
        return "", err
    }
    return "{{.ModuleNameCapital}} successfully deleted", nil
}

func Delete{{.ModuleNameCapital}}Field() *graphql.Field {
    return &graphql.Field{
        Type:        graphql.String,
        Description: "Delete a {{.ModuleNameLowerCase}}",
        Args: graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{
                Type: graphql.NewNonNull(graphql.Int),
            },
        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            id := p.Args["id"].(int)
            // Call the delete function with the provided ID
            msg, err := Delete{{.ModuleNameCapital}}(id)
            if err != nil {
                return nil, err
            }
            return msg, nil
        },
    }
}
