package types

import (
    "github.com/graphql-go/graphql"
    "gorm.io/gorm"
)

// Page struct for GORM model
type Page struct {
    gorm.Model
    Title string `gorm:"column:title;json:\"title\""`
    Content string `gorm:"column:content;json:\"content\""`
}

// PageType for GraphQL schema
var PageType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Page",
        Fields: graphql.Fields{
            "title": &graphql.Field{
                Type: graphql.String,
                Description: "Field for Title",
            },
            "content": &graphql.Field{
                Type: graphql.String,
                Description: "Field for Content",
            },
        },
    },
)
