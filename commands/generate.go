// File: commands/generate.go
package commands

import (
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate is the command used for generating various parts of the application",
	Long:  `Generate is used to scaffold entire parts of the application such as models, controllers, etc.`,
}

func init() {
	RootCmd.AddCommand(generateCmd)
}
