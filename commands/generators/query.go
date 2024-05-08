package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"embed"

	"github.com/spf13/cobra"
)

// Embed the query templates
//
//go:embed templates/queries/*
var queryTemplates embed.FS

var QueryCmd = &cobra.Command{
	Use:   "query [module]",
	Short: "Generates GraphQL query handlers for a module",
	Long: `Generates GraphQL query handlers for the specified module.
Example usage:
    base generate query Post`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		moduleName := args[0]
		if err := generateQueries(moduleName); err != nil {
			fmt.Println("Error generating queries for module:", err)
			return
		}
		fmt.Printf("Successfully generated queries for module: %s\n", moduleName)
	},
}

func generateQueries(moduleName string) error {
	baseDir := filepath.Join("app", strings.ToLower(moduleName), "queries")
	if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
		return err
	}

	// Process each template for queries
	templateFiles := []string{"get.go.tpl"}
	for _, t := range templateFiles {
		templatePath := fmt.Sprintf("templates/queries/%s", t)
		outputPath := filepath.Join(baseDir, strings.TrimSuffix(t, ".tpl"))
		if err := processTemplate(templatePath, outputPath, map[string]interface{}{
			"Namespace":           getCurrentDir(), // Current dir is the namespace
			"ModuleName":          moduleName,
			"ModuleNameLowerCase": strings.ToLower(moduleName),
			"ModuleNameCapital":   strings.Title(moduleName),
		}, queryTemplates); err != nil {
			return err
		}
	}

	return nil
}
