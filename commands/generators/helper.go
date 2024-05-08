package generators

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Field struct {
	Name       string
	GoType     string
	GqlType    string
	TitledName string
	LowerName  string
}

// Updated ValidFieldTypes map to handle more types like json and time.Time
var ValidFieldTypes = map[string]string{
	"int":       "Int",
	"string":    "String",
	"json":      "String", // Treat JSON as string for simplicity in GraphQL
	"time.Time": "String", // Treat time as string, possibly should be DateTime if available
}

func getCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return filepath.Base(dir)
}

func parseFields(fieldSpecs []string) ([]Field, error) {
	var fields []Field
	for _, fs := range fieldSpecs {
		parts := strings.Split(fs, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid field specification: %s", fs)
		}
		fieldName := parts[0]
		fieldType := parts[1]

		fields = append(fields, Field{
			Name:   fieldName,
			GoType: fieldType,
			// Define GQL types based on Go types
			// You can extend this mapping as needed
			GqlType:    getGqlType(fieldType),
			TitledName: strings.Title(fieldName),
			LowerName:  strings.ToLower(fieldName),
		})
	}
	return fields, nil
}

func getGqlType(goType string) string {
	// Map Go types to GraphQL types
	switch goType {
	case "int":
		return "Int"
	case "string":
		return "String"
	case "json":
		return "JSON"
	case "time.Time":
		return "DateTime"
	default:
		// Default to String for unknown types
		return "String"
	}
}

// processTemplate processes a template file with the provided data and writes the output to a specified path.
func processTemplate(templatePath, outputPath string, data interface{}, fs embed.FS) error {
	// Create a new template and parse it from the embedded file system
	t, err := template.New(filepath.Base(templatePath)).Funcs(createFuncMap()).ParseFS(fs, templatePath)
	if err != nil {
		return err
	}

	// Create the output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Execute the template with the provided data
	if err := t.Execute(outputFile, data); err != nil {
		return err
	}

	return nil
}

// createFuncMap returns a FuncMap for use in templates, adding any necessary custom functions
func createFuncMap() template.FuncMap {
	return template.FuncMap{
		"lower": strings.ToLower,
		"title": strings.Title,
	}
}
