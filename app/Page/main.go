package page

import (
	"base-cli/app/page/mutations"
	"base-cli/app/page/queries"
	"base-cli/app/page/types"
	"base-cli/core/database"
	"fmt"

	"github.com/graphql-go/graphql"
)

func InitPageModule(schema *graphql.Schema) {
	// Initialize the page module
	fmt.Println("Initializing page module")
	Migrate()
	CreateSchema()
	CreateQuery()
	CreateMutation()
	fmt.Println("Page module initialized")
}
func Migrate() {
	// Migrate the page module
	database.DB.AutoMigrate(&types.Page{})
}
func CreateSchema() graphql.Schema {
	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    CreateQuery(),
			Mutation: CreateMutation(),
		},
	)
	return schema
}

func CreateQuery() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"getAllPages": queries.GetAllPagesField(),
				"getPageById": queries.GetPageByIDField(),
			},
		},
	)
}

func CreateMutation() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"createPage": mutations.CreatePageField(),
				"updatePage": mutations.UpdatePageField(),
				"deletePage": mutations.DeletePageField(),
			},
		},
	)
}
