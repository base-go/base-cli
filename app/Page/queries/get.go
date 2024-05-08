package queries

import (
    "base-cli/app/page/types"
    "base-cli/core/database"
    "errors"

    "github.com/graphql-go/graphql"
)

func GetAllPagesField() *graphql.Field {
    return &graphql.Field{
        Type:        graphql.NewList(types.PageType),
        Description: "Get all pages",
        Resolve: func(params graphql.ResolveParams) (interface{}, error) {
            var pages []types.Page
            if err := database.DB.Find(&pages).Error; err != nil {
                return nil, err // Consider wrapping with a more user-friendly message
            }
            return pages, nil
        },
    }
}

func GetPageByIDField() *graphql.Field {
    return &graphql.Field{
        Type:        types.PageType,
        Description: "Get page by ID",
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
            var page types.Page
            if err := database.DB.First(&page, id).Error; err != nil {
                return nil, err // Consider wrapping with a more user-friendly message
            }
            return &page, nil
        },
    }
}
