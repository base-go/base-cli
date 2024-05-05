// File: commands/root.go
package commands

import (
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "base",
	Short: "Base is the primary command for the Base-Graph CLI",
	Long: `Base is a CLI application designed for creating and managing
GraphQL-based projects and APIs. It provides tools to generate necessary boilerplate,
manage database interactions, and more.`,
	// Uncomment the following line if your bare application has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return RootCmd.Execute()
}
