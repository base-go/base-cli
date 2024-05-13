package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [Project]",
	Short: "Create a new BaseQL project",
	Long:  `Clone the base-project template and create a new BaseQL project with the specified name.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		project := args[0]

		// Clone base-project to a temporary directory
		tempDir := "/tmp/base-project-clone"
		_, err := exec.Command("git", "clone", "https://github.com/base-go/base-project.git", tempDir).Output()
		if err != nil {
			fmt.Println("Error cloning base-project:", err)
			return
		}

		// Rename directory to project name
		newDir := fmt.Sprintf("./%s", project)
		err = os.Rename(tempDir, newDir)
		if err != nil {
			fmt.Println("Error renaming directory:", err)
			return
		}

		fmt.Printf("Created new BaseQL project '%s'\n", project)
	},
}

func init() {
	RootCmd.AddCommand(newCmd)
}
