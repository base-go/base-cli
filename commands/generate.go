package commands

import (
	"base-go/base-graph/commands/generators"

	"github.com/spf13/cobra"
)

// generateCmd represents the base command when called without any subcommands
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate parts of the GraphQL application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command. For example: Generate is used to scaffold parts of a GraphQL application such as types, queries, mutations, and more.`,
	// Uncomment the following line if your bare application has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// This function will be called to add the above command to the root command.
func init() {
	RootCmd.AddCommand(generateCmd)
	// Here you would add subcommands
	generateCmd.AddCommand(generators.ScaffoldCmd)
	generateCmd.AddCommand(generators.QueryCmd)
	generateCmd.AddCommand(generators.MutationCmd)
	generateCmd.AddCommand(generators.TypeCmd)
}
