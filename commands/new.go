package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Creates a new project or entity",
	Long: `Creates a new project or entity with the specified name.
Example usage:
	base new projectName`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("New project created with name:", args[0])
	},
}

func init() {
	RootCmd.AddCommand(newCmd)
}
