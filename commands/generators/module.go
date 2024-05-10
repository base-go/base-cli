package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

// Define the command line application
var ModuleCmd = &cobra.Command{
	Use:   "generate module [name] [fields]",
	Short: "Generates a new module with specified fields",
	Args:  cobra.MinimumNArgs(2),
	Run:   runGenerateModule,
}

func runGenerateModule(cmd *cobra.Command, args []string) {
	moduleName := args[0]
	fields := parseFields(args[1:])
	data := map[string]interface{}{
		"ModuleName": moduleName,
		"Fields":     fields,
	}

	templates := []struct {
		path     string
		filename string
	}{
		{"templates/types", "type.go.tpl"},
		{"templates/queries", "get.go.tpl"},
		{"templates/mutations", "create.go.tpl"},
		{"templates/mutations", "update.go.tpl"},
		{"templates/mutations", "delete.go.tpl"},
	}

	for _, t := range templates {
		if err := processTemplate(t.path, t.filename, moduleName, data); err != nil {
			fmt.Fprintf(os.Stderr, "Error processing template %s: %v\n", t.filename, err)
			os.Exit(1)
		}
	}
	fmt.Printf("Module %s generated successfully.\n", moduleName)
}

func parseFields(fieldSpecs []string) []map[string]string {
	var fields []map[string]string
	for _, f := range fieldSpecs {
		parts := strings.Split(f, ":")
		if len(parts) != 2 {
			fmt.Fprintf(os.Stderr, "Invalid field format: %s\n", f)
			continue
		}
		fields = append(fields, map[string]string{"name": parts[0], "type": parts[1]})
	}
	return fields
}

func processTemplate(dir, filename, moduleName string, data map[string]interface{}) error {
	tmpl, err := template.ParseFiles(filepath.Join(dir, filename))
	if err != nil {
		return err
	}

	outputDir := filepath.Join("output", moduleName, strings.TrimSuffix(filename, ".tpl"))
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	outputFile, err := os.Create(filepath.Join(outputDir, strings.ReplaceAll(filename, ".tpl", ".go")))
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return tmpl.Execute(outputFile, data)
}
