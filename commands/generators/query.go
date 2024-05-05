package generators

import (
	"fmt"

	"github.com/spf13/cobra"
)

// queryCmd represents the query generation command
var QueryCmd = &cobra.Command{
	Use:   "query [type]",
	Short: "Generates a GraphQL query handler",
	Long: `Generates a GraphQL query handler for the specified type.
Example usage:
	base generate query UserType`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating query handler for type:", args[0])
	},
}
