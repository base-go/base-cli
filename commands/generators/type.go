package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"embed"

	"github.com/spf13/cobra"
)

// Embed the types templates
//
//go:embed templates/types/*
var typesTemplates embed.FS

var TypeCmd = &cobra.Command{
	Use:   "types [module] [fields]",
	Short: "Generates GraphQL type definitions for a module",
	Long: `Generates GraphQL type definitions for the specified module.
Example usage:
    base generate types Post title:string content:string`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		moduleName := args[0]

		fieldsSpecs := args[1:]
		fields, err := parseFields(fieldsSpecs)
		if err != nil {
			fmt.Println("Error parsing fields:", err)
			return
		}
		if err := generateTypes(moduleName, fields); err != nil {
			fmt.Println("Error generating types for module:", err)
			return
		}
		fmt.Printf("Successfully generated types for module: %s\n", moduleName)
	},
}

func generateTypes(moduleName string, fields []Field) error {
	baseDir := filepath.Join("app", strings.ToLower(moduleName), "types")
	if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
		return err
	}

	// Parse field specifications into Field structs

	// Process each template for types
	templateFiles := []string{"inputTypes.go.tpl", "type.go.tpl"}
	for _, t := range templateFiles {
		templatePath := fmt.Sprintf("templates/types/%s", t)
		outputPath := filepath.Join(baseDir, strings.TrimSuffix(t, ".tpl"))
		if err := processTemplate(templatePath, outputPath, map[string]interface{}{
			"Namespace":           getCurrentDir(), // Current dir is the namespace
			"ModuleName":          moduleName,
			"ModuleNameLowerCase": strings.ToLower(moduleName),
			"ModuleNameCapital":   strings.Title(moduleName),
			"Fields":              fields,
		}, typesTemplates); err != nil {
			return err
		}
	}

	return nil
}
