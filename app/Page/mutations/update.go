package mutations

import (
    "base-cli/app/page/types"
    "base-cli/core/database"
    "log"

    "github.com/graphql-go/graphql"
)

func UpdatePage(input types.UpdatePageInput) (*types.Page, error) {
    var page types.Page
    // Fetch the existing record to update
    if err := database.DB.First(&page, input.ID).Error; err != nil {
        log.Printf("Error finding page: %v", err)
        return nil, err
    }

    // Update only the fields that were actually provided
    if input.title != nil {
        page.title = *input.title
    }
    if input.content != nil {
        page.content = *input.content
    }

    // Save the updated record
    if err := database.DB.Save(&page).Error; err != nil {
        log.Printf("Error updating page: %v", err)
        return nil, err
    }

    return &page, nil
}

func UpdatePageField() *graphql.Field {
    return &graphql.Field{
        Type:        types.PageType,
        Description: "Update an existing page",
        Args: graphql.FieldConfigArgument{
            "input": &graphql.ArgumentConfig{
                Type: graphql.NewNonNull(graphql.NewInputObject(graphql.InputObjectConfig{
                    Name: "UpdatePageInput",
                    Fields: graphql.InputObjectConfigFieldMap{
                        "id": &graphql.InputObjectFieldConfig{
                            Type: graphql.NewNonNull(graphql.Int), // Ensure this matches the ID type in your schema
                        },
                        "title": &graphql.InputObjectFieldConfig{
                            Type: graphql.NewNonNull(String), // Use dynamic GraphQL types
                        },
                        "content": &graphql.InputObjectFieldConfig{
                            Type: graphql.NewNonNull(String), // Use dynamic GraphQL types
                        },
                    },
                })),
            },
        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            inputMap := p.Args["input"].(map[string]interface{})
            updateInput := types.UpdatePageInput{
                ID:      inputMap["id"].(int), // Ensure casting matches the ID type
                title: optionalstring(inputMap["title"]),
                content: optionalstring(inputMap["content"]),
            }
            return UpdatePage(updateInput)
        },
    }
}

// This needs to be adjusted if you support more types than strings.
func optional<no value>(val interface{}) *<no value> {
    if v, ok := val.(<no value>); ok {
        return &v
    }
    return nil
}
