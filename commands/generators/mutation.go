package generators

import (
	"fmt"

	"github.com/spf13/cobra"
)

// mutationCmd represents the mutation generation command
var MutationCmd = &cobra.Command{
	Use:   "mutation [type]",
	Short: "Generates a GraphQL mutation handler",
	Long: `Generates a GraphQL mutation handler for the specified type.
Example usage:
	base generate mutation UserType`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating mutation handler for type:", args[0])
	},
}
