package generators

import (
	"fmt"

	"github.com/spf13/cobra"
)

// scaffoldCmd represents the scaffold command
var ScaffoldCmd = &cobra.Command{
	Use:   "scaffold [type] [fields]",
	Short: "Scaffolds both queries and mutations for a type",
	Long: `Scaffolds both GraphQL queries and mutations for the specified type.
Example usage:
	base generate scaffold User name:string email:string password:string`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Scaffolding GraphQL handlers for type:", args[0])
		// Assume we call both query and mutation generation functions here
	},
}

func init() {
	// You might want to add this command to a parent command if needed
}
