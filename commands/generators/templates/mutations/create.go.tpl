package mutations

import (
	"{{.Namespace}}/app/{{.ModuleNameLowerCase}}/types"
	"{{.Namespace}}/core/database"
	"log"

	"github.com/graphql-go/graphql"
)

func Create{{.ModuleNameCapital}}(input types.{{.ModuleNameCapital}}Input) (*types.{{.ModuleNameCapital}}, error) {
	{{.ModuleNameLowerCase}} := &types.{{.ModuleNameCapital}}{
		{{ range .Fields }}
		{{.Name}}: input.{{.Name}},
		{{- end }}
	}
	if err := database.DB.Create({{.ModuleNameLowerCase}}).Error; err != nil {
		log.Printf("Error creating {{.ModuleNameLowerCase}}: %v", err)
		return nil, err
	}
	return {{.ModuleNameLowerCase}}, nil
}
// Create{{.ModuleNameCapital}}Field returns a GraphQL field configuration for creating a {{.ModuleNameLowerCase}}.
func Create{{.ModuleNameCapital}}Field() *graphql.Field {
    return &graphql.Field{
        Type:        types.{{.ModuleNameCapital}}Type,
        Description: "Create a new {{.ModuleNameLowerCase}}",
        Args: graphql.FieldConfigArgument{
            "input": &graphql.ArgumentConfig{
                Type: graphql.NewNonNull(graphql.NewInputObject(graphql.InputObjectConfig{
                    Name: "{{.ModuleNameCapital}}Input",
                    Fields: graphql.InputObjectConfigFieldMap{
                        {{- range .Fields }}
                        "{{.Name}}": &graphql.InputObjectFieldConfig{
                            Type: graphql.NewNonNull({{.GqlType}}), // Adjusted to use dynamic type
                        },
                        {{- end }}
                    },
                })),
            },
        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            input, _ := p.Args["input"].(map[string]interface{})
            {{.ModuleNameLowerCase}}Input := types.{{.ModuleNameCapital}}Input{
                {{- range .Fields }}
                {{.Name}}: input["{{.Name}}"].({{.GoType}}), // Adjusted for dynamic casting
                {{- end }}
            }
            return Create{{.ModuleNameCapital}}({{.ModuleNameLowerCase}}Input)
        },
    }
}
