package resolvers

import (
    "{{.Namespace}}/app/{{.ModuleNameLowerCase}}/mutations"
    "{{.Namespace}}/app/{{.ModuleNameLowerCase}}/types"
)

// MutationResolver resolves mutations related to {{.ModuleNameCapital}}s
type MutationResolver struct{}

// Create{{.ModuleNameCapital}} resolves the mutation for creating a new {{.ModuleNameLowerCase}}
func (r *MutationResolver) Create{{.ModuleNameCapital}}(input types.{{.ModuleNameCapital}}Input) (*types.{{.ModuleNameCapital}}, error) {
    return mutations.Create{{.ModuleNameCapital}}(input)
}

// Update{{.ModuleNameCapital}} resolves the mutation for updating an existing {{.ModuleNameLowerCase}}
func (r *MutationResolver) Update{{.ModuleNameCapital}}(id int, input types.Update{{.ModuleNameCapital}}Input) (*types.{{.ModuleNameCapital}}, error) {
    return mutations.Update{{.ModuleNameCapital}}(id, input)
}

// Delete{{.ModuleNameCapital}} resolves the mutation for deleting a {{.ModuleNameLowerCase}}
func (r *MutationResolver) Delete{{.ModuleNameCapital}}(id int) (string, error) {
    return mutations.Delete{{.ModuleNameCapital}}(id)
}
