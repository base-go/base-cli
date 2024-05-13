package generators

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed templates/*
var templatesFS embed.FS

// Define the command line application
var ModuleCmd = &cobra.Command{
	Use:   "module [name] [fields]",
	Short: "Generates a new module with specified fields",
	Args:  cobra.MinimumNArgs(2),
	Run:   runGenerateModule,
}

func runGenerateModule(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: module [name] [field1:type1] [field2:type2] ...")
		os.Exit(1)
	}
	moduleName := args[0]
	fields := parseFields(args[1:])
	namespace, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current directory: %v\n", err)
		os.Exit(1)
	}
	currDir := filepath.Base(namespace)

	data := map[string]interface{}{
		// Namespace is the current directory name
		"Namespace":           currDir,
		"ModuleName":          moduleName,
		"ModuleNameLowerCase": strings.ToLower(moduleName),
		"ModuleNameCapital":   cases.Title(language.English).String(moduleName),
		"Fields":              fields,
	}

	templates := []struct {
		filename string
	}{
		{"templates/main.go.tpl"},
		{"templates/types/type.go.tpl"},
		{"templates/queries/get.go.tpl"},
		{"templates/mutations/create.go.tpl"},
		{"templates/mutations/update.go.tpl"},
		{"templates/mutations/delete.go.tpl"},
	}

	for _, t := range templates {
		if err := processTemplate(t.filename, moduleName, data); err != nil {
			fmt.Fprintf(os.Stderr, "Error processing template %s: %v\n", t.filename, err)
			os.Exit(1)
		}
	}
	fmt.Printf("Module %s generated successfully.\n", moduleName)
}
func parseFields(fieldSpecs []string) []map[string]string {
	var fields []map[string]string
	caser := cases.Title(language.English) // Create a Title Casing caser for English

	for _, f := range fieldSpecs {
		parts := strings.Split(f, ":")
		if len(parts) != 2 {
			fmt.Fprintf(os.Stderr, "Invalid field format: %s. Expecting 'name:type'.\n", f)
			continue
		}
		name, fieldType := parts[0], parts[1]

		// Here, add logic to determine GqlType based on fieldType if needed.
		gqlType := "graphql.String" // Default, change logic as needed based on actual type

		if fieldType == "int" {
			gqlType = "graphql.Int"
		} else if fieldType == "float" {
			gqlType = "graphql.Float"
		} else if fieldType == "bool" {
			gqlType = "graphql.Boolean"
		} else if fieldType == "time.Time" {
			gqlType = "graphql.DateTime"
		} else if fieldType == "[]string" {
			gqlType = "graphql.NewList(graphql.String)"
		} else if strings.HasPrefix(fieldType, "[]") {
			innerType := strings.TrimPrefix(fieldType, "[]")
			gqlType = fmt.Sprintf("graphql.NewList(%s)", innerType)
		}

		fields = append(fields, map[string]string{
			"Name":       name,
			"TitledName": caser.String(name), // Title case the name (e.g., "first_name" -> "First_Name"
			"Type":       fieldType,
			"GqlType":    gqlType,
		})
	}
	return fields
}
func processTemplate(templatePath, moduleName string, data map[string]interface{}) error {
	parts := strings.Split(templatePath, "/")
	if len(parts) < 2 {
		return fmt.Errorf("invalid template path: %s", templatePath)
	}
	mainDir := parts[1]             // This is 'mutations', 'queries', or 'types'
	filename := parts[len(parts)-1] // Last part of the path
	if mainDir == filename {
		filename = "main.go"
		mainDir = "."
	}
	// Lowercase the module name for directory naming
	moduleNameLower := strings.ToLower(moduleName)

	// Create the output directory based on the module name and the main directory
	outputDir := filepath.Join("app", moduleNameLower, mainDir)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %v", outputDir, err)
	}

	// Determine the base filename by trimming the '.tpl' extension and adjusting for specific cases
	baseFilename := strings.TrimSuffix(filename, ".tpl")
	var newFilename string
	if baseFilename == "type" {
		newFilename = moduleNameLower // Special case for type.go to be named as user.go
	} else {
		newFilename = baseFilename // Normal case, directly use the base filename
	}

	// Read the template file
	tmplData, err := templatesFS.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template file: %v", err)
	}
	tmpl, err := template.New(filename).Parse(string(tmplData))
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}
	// Replace template data with the module name
	data["ModuleNameLower"] = moduleNameLower

	// Construct the output path for the new file
	outputPath := filepath.Join(outputDir, newFilename)
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", outputPath, err)
	}
	defer outputFile.Close()

	// Execute the template and write the output
	if err := tmpl.Execute(outputFile, data); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	return nil
}
