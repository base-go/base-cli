package mutations

import (
	"base-cli/app/page/types"
	"base-cli/core/database"
	"log"

	"github.com/graphql-go/graphql"
)

func CreatePage(input types.PageInput) (*types.Page, error) {
	page := &types.Page{
		
		title: input.title,
		content: input.content,
	}
	if err := database.DB.Create(page).Error; err != nil {
		log.Printf("Error creating page: %v", err)
		return nil, err
	}
	return page, nil
}
// CreatePageField returns a GraphQL field configuration for creating a page.
func CreatePageField() *graphql.Field {
    return &graphql.Field{
        Type:        types.PageType,
        Description: "Create a new page",
        Args: graphql.FieldConfigArgument{
            "input": &graphql.ArgumentConfig{
                Type: graphql.NewNonNull(graphql.NewInputObject(graphql.InputObjectConfig{
                    Name: "PageInput",
                    Fields: graphql.InputObjectConfigFieldMap{
                        "title": &graphql.InputObjectFieldConfig{
                            Type: graphql.NewNonNull(String), // Adjusted to use dynamic type
                        },
                        "content": &graphql.InputObjectFieldConfig{
                            Type: graphql.NewNonNull(String), // Adjusted to use dynamic type
                        },
                    },
                })),
            },
        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            input, _ := p.Args["input"].(map[string]interface{})
            pageInput := types.PageInput{
                title: input["title"].(string), // Adjusted for dynamic casting
                content: input["content"].(string), // Adjusted for dynamic casting
            }
            return CreatePage(pageInput)
        },
    }
}
