package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"embed"

	"github.com/spf13/cobra"
)

// Embed the mutation templates
//
//go:embed templates/*
var mutationTemplates embed.FS

var MutationCmd = &cobra.Command{
	Use:   "mutation [module] [fields]",
	Short: "Generates GraphQL mutation handlers for a module",
	Long: `Generates GraphQL mutation handlers for the specified module based on provided fields.
Example usage:
    base generate mutation Post title:string content:string`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		moduleName := args[0]
		fieldsSpecs := args[1:]
		fields, err := parseFields(fieldsSpecs)
		if err != nil {
			fmt.Println("Error parsing fields:", err)
			return
		}
		if err := generateMutations(moduleName, fields); err != nil {
			fmt.Println("Error generating mutations for module:", err)
			return
		}
		fmt.Printf("Successfully generated mutations for module: %s\n", moduleName)
	},
}

func generateMutations(moduleName string, fields []Field) error {
	baseDir := filepath.Join("app", moduleName, "mutations")
	if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
		return err
	}

	// Process each template for mutations
	templateFiles := []string{"create.go.tpl", "update.go.tpl", "delete.go.tpl"}
	for _, t := range templateFiles {
		templatePath := fmt.Sprintf("templates/mutations/%s", t)
		outputPath := filepath.Join(baseDir, strings.TrimSuffix(t, ".tpl"))

		// Ensure context map has the necessary keys
		contextMap := map[string]interface{}{
			"Namespace":           getCurrentDir(), // Current dir is the namespace
			"ModuleName":          moduleName,
			"ModuleNameLowerCase": strings.ToLower(moduleName),
			"ModuleNameCapital":   strings.Title(moduleName),
			"Fields":              fields, // Pass fields directly
		}

		if err := processTemplate(templatePath, outputPath, contextMap, mutationTemplates); err != nil {
			return err
		}
	}

	return nil
}
