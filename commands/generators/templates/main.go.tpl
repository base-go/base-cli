package {{ .ModuleNameLowerCase}}

import (
	"{{.Namespace}}/app/{{.ModuleNameLowerCase}}/mutations"
	"{{.Namespace}}/app/{{.ModuleNameLowerCase}}/queries"
	"{{.Namespace}}/app/{{.ModuleNameLowerCase}}/types"
	"{{.Namespace}}/core/database"
	"fmt"

	"github.com/graphql-go/graphql"
)

func Init{{.ModuleNameCapital}}Module(schema *graphql.Schema) {
	// Initialize the {{.ModuleNameLowerCase}} module
	fmt.Println("Initializing {{.ModuleNameLowerCase}} module")
	Migrate()
	CreateSchema()
	CreateQuery()
	CreateMutation()
	fmt.Println("{{.ModuleNameCapital}} module initialized")
}
func Migrate() {
	// Migrate the {{.ModuleNameLowerCase}} module
	database.DB.AutoMigrate(&types.{{.ModuleNameCapital}}{})
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
				"getAll{{.ModuleNameCapital}}s": queries.GetAll{{.ModuleNameCapital}}sField(),
				"get{{.ModuleNameCapital}}ById": queries.Get{{.ModuleNameCapital}}ByIDField(),
			},
		},
	)
}

func CreateMutation() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"create{{.ModuleNameCapital}}": mutations.Create{{.ModuleNameCapital}}Field(),
				"update{{.ModuleNameCapital}}": mutations.Update{{.ModuleNameCapital}}Field(),
				"delete{{.ModuleNameCapital}}": mutations.Delete{{.ModuleNameCapital}}Field(),
			},
		},
	)
}





package user

import (
	"base-project/app/user/mutations"
	"base-project/app/user/queries"
	"base-project/app/user/types"
	"base-project/core/database"
	"fmt"

	"github.com/graphql-go/graphql"
)

type {{.ModuleNameCapital}}Module struct{}

func (p *{{.ModuleNameCapital}}Module) CreateQuery() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "{{.ModuleNameCapital}}Queries",
			Fields: graphql.Fields{
				"list": queries.All(),
				"show": queries.ByID(),
			},
		},
	)
}

func (p *{{.ModuleNameCapital}}Module) CreateMutation() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "{{.ModuleNameCapital}}Mutations",
			Fields: graphql.Fields{
				"create{{.ModuleNameCapital}}": mutations.Create{{.ModuleNameCapital}}Field(),
				"update{{.ModuleNameCapital}}": mutations.Update{{.ModuleNameCapital}}Field(),
				"delete{{.ModuleNameCapital}}": mutations.Delete{{.ModuleNameCapital}}Field(),
			},
		},
	)
}

func (u *{{.ModuleNameCapital}}Module) Resolvable() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		// Default resolve function or specific logic
		return struct{}{}, nil
	}
}

func (p *{{.ModuleNameCapital}}Module) Migrate() error {
	// Migrate the user database model
	fmt.Println("Migrating user model...")
	database.DB.AutoMigrate(&types.{{.ModuleNameCapital}}{})
	if err := database.DB.AutoMigrate(&types.{{.ModuleNameCapital}}{}); err != nil {
		fmt.Println("{{.ModuleNameCapital}} model migration failed:", err)
		return err
	}
	fmt.Println("{{.ModuleNameCapital}} model migration completed.")
	return nil
}
