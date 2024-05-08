package types

// {{.ModuleNameCapital}}Input represents the input type for creating/updating a {{.ModuleNameLowerCase}}

type {{.ModuleNameCapital}}Input struct {
    gorm.Model
    {{- range .Fields }}
    {{.TitledName}} {{.GoType}} `gorm:"column:{{.LowerName}};json:\"{{.LowerName}}\""`
    {{- end }}
}

// Update{{.ModuleNameCapital}}Input represents the input type for updating a {{.ModuleNameLowerCase}}

type Update{{.ModuleNameCapital}}Input struct {
    ID      int     `json:"id"`
    {{- range .Fields }}
    {{.TitledName}} {{.GoType}} `json:"{{.LowerName}}"`
    {{- end }}
}
