package mutations

import (
    "{{.Namespace}}/app/{{.ModuleNameLowerCase}}/types"
    "{{.Namespace}}/core/database"
    "log"

    "github.com/graphql-go/graphql"
)

func Update{{.ModuleNameCapital}}(input types.Update{{.ModuleNameCapital}}Input) (*types.{{.ModuleNameCapital}}, error) {
    var {{.ModuleNameLowerCase}} types.{{.ModuleNameCapital}}
    // Fetch the existing record to update
    if err := database.DB.First(&{{.ModuleNameLowerCase}}, input.ID).Error; err != nil {
        log.Printf("Error finding {{.ModuleNameLowerCase}}: %v", err)
        return nil, err
    }

    // Update only the fields that were actually provided
    {{- range .Fields }}
    if input.{{.TitledName}} != nil {
        {{$.ModuleNameLowerCase}}.{{.TitledName}} = *input.{{.TitledName}}
    }
    {{- end }}

    // Save the updated record
    if err := database.DB.Save(&{{.ModuleNameLowerCase}}).Error; err != nil {
        log.Printf("Error updating {{.ModuleNameLowerCase}}: %v", err)
        return nil, err
    }

    return &{{.ModuleNameLowerCase}}, nil
}

func Update{{.ModuleNameCapital}}Field() *graphql.Field {
    return &graphql.Field{
        Type:        types.{{.ModuleNameCapital}}Type,
        Description: "Update an existing {{.ModuleNameLowerCase}}",
        Args: graphql.FieldConfigArgument{
            "input": &graphql.ArgumentConfig{
                Type: graphql.NewNonNull(graphql.NewInputObject(graphql.InputObjectConfig{
                    Name: "Update{{.ModuleNameCapital}}Input",
                    Fields: graphql.InputObjectConfigFieldMap{
                        "id": &graphql.InputObjectFieldConfig{
                            Type: graphql.NewNonNull(graphql.Int), // Ensure this matches the ID type in your schema
                        },
                        {{- range .Fields }}
                        "{{.Name}}": &graphql.InputObjectFieldConfig{
                            Type: graphql.{{.GqlType}}, // Use dynamic GraphQL types
                        },
                        {{- end }}
                    },
                })),
            },
        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            inputMap := p.Args["input"].(map[string]interface{})
            updateInput := types.Update{{.ModuleNameCapital}}Input{
                ID:      inputMap["id"].(int), // Ensure casting matches the ID type
                {{- range .Fields }}
                {{.TitledName}}: optionalString(inputMap["{{.Name}}"]),
                {{- end }}
            }
            return Update{{.ModuleNameCapital}}(updateInput)
        },
    }
}

// This needs to be adjusted if you support more types than strings.

func optionalString(val interface{}) *string {
	if str, ok := val.(string); ok {
		return &str
	}
	return nil
}
