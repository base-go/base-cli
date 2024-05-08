package generators

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

var rootDir string

//go:embed templates/*
var mainTemplate embed.FS

var ScaffoldCmd = &cobra.Command{
	Use:   "scaffold [module] [fields]",
	Short: "Generates a scaffold for a module",
	Long:  `Generates a scaffold for the specified module.`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		moduleName := args[0]
		fieldSpecs := args[1:]

		fields, err := parseFields(fieldSpecs)
		if err != nil {
			fmt.Println("Error parsing fields:", err)
			return
		}
		if err := generateScaffold(moduleName, fields); err != nil {
			fmt.Println("Error generating scaffold:", err)
			return
		}
		fmt.Printf("Successfully generated scaffold for module: %s\n", moduleName)
	},
}

func init() {
	ScaffoldCmd.Flags().StringVarP(&rootDir, "dir", "d", ".", "Root directory for generating files")
}
func generateMain(rootDir, moduleName string) error {
	// Define data for template
	data := struct {
		ModuleName          string
		ModuleNameLowerCase string
		ModuleNameCapital   string
		Namespace           string
	}{
		ModuleName:          moduleName,
		ModuleNameLowerCase: strings.ToLower(moduleName),
		ModuleNameCapital:   strings.Title(moduleName),
		Namespace:           getCurrentDir(), // Current dir is the namespace
	}

	// Read the template file content
	templateContent, err := mainTemplate.ReadFile("templates/main.go.tpl")
	if err != nil {
		return err
	}

	// Parse and execute template
	tmpl := template.Must(template.New("main").Parse(string(templateContent)))
	outputPath := filepath.Join(rootDir, "app", moduleName, "main.go")
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = tmpl.Execute(outputFile, data)
	if err != nil {
		return err
	}

	return nil
}

func generateScaffold(moduleName string, fields []Field) error {
	// Create module directory
	moduleDir := filepath.Join(rootDir, "app", moduleName)
	if err := os.MkdirAll(moduleDir, os.ModePerm); err != nil {
		return err
	}

	// Parse field specifications into Field structs

	fmt.Println("Generating scaffold for module:", moduleName)
	fmt.Println("Fields:", fields)

	// Generate mutations, queries, types, and main.go
	if err := generateMutations(moduleName, fields); err != nil {
		return err
	}
	if err := generateQueries(moduleName); err != nil {
		return err
	}
	if err := generateTypes(moduleName, fields); err != nil {
		return err
	}
	if err := generateMain(rootDir, moduleName); err != nil {
		return err
	}

	return nil
}
