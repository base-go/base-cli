package types

import "gorm.io/gorm"


// {{.ModuleNameCapital}}Input represents the input type for creating/updating a {{.ModuleNameLowerCase}}

type {{.ModuleNameCapital}}Input struct {
    gorm.Model
    {{- range .Fields }}
    {{.TitledName}} {{.Type}} `json:"{{.LowerName}}"`
    {{- end }}
}

// Update{{.ModuleNameCapital}}Input represents the input type for updating a {{.ModuleNameLowerCase}}

type Update{{.ModuleNameCapital}}Input struct {
    ID      int     `json:"id"`
    {{- range .Fields }}
    {{.TitledName}} *{{.Type}} `json:"{{.LowerName}}"`
    {{- end }}
}
