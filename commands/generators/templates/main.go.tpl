package {{ .ModuleNameLowerCase}}

import (
	"{{.Namespace}}/app/{{.ModuleNameLowerCase}}/mutations"
	"{{.Namespace}}/app/{{.ModuleNameLowerCase}}/queries"
	"{{.Namespace}}/app/{{.ModuleNameLowerCase}}/types"
	"{{.Namespace}}/core/database"
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
				"create": mutations.Create{{.ModuleNameCapital}}Field(),
				"update": mutations.Update{{.ModuleNameCapital}}Field(),
				"delete": mutations.Delete{{.ModuleNameCapital}}Field(),
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
	fmt.Println("Migrating {{.ModuleNameLowerCase}} model...")
	database.DB.AutoMigrate(&types.{{.ModuleNameCapital}}{})
	if err := database.DB.AutoMigrate(&types.{{.ModuleNameCapital}}{}); err != nil {
		fmt.Println("{{.ModuleNameCapital}} model migration failed:", err)
		return err
	}
	fmt.Println("{{.ModuleNameCapital}} model migration completed.")
	return nil
}
