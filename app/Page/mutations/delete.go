package mutations

import (
    "base-cli/app/page/types"
    "base-cli/core/database"
    "log"

    "github.com/graphql-go/graphql"
)

func DeletePage(id int) (string, error) {
    var page types.Page
    // Delete the record with the provided ID
    if err := database.DB.Delete(&page, id).Error; err != nil {
        log.Printf("Error deleting page with ID %d: %v", id, err)
        return "", err
    }
    return "Page successfully deleted", nil
}

func DeletePageField() *graphql.Field {
    return &graphql.Field{
        Type:        graphql.String,
        Description: "Delete a page",
        Args: graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{
                Type: graphql.NewNonNull(graphql.Int),
            },
        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            id := p.Args["id"].(int)
            // Call the delete function with the provided ID
            msg, err := DeletePage(id)
            if err != nil {
                return nil, err
            }
            return msg, nil
        },
    }
}
