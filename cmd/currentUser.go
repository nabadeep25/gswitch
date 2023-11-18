package cmd

import (
	"fmt"
	"os"

	"github.com/nabadeep25/gswitch/pkg/helper"
	"github.com/spf13/cobra"
)

var currentUser = &cobra.Command{
	Use:   "current",
	Short: "Display current global git user",
	Long:  `Display current global git user`,
	Run: func(cmd *cobra.Command, args []string) {
		activeUsername, activeEmail, err := helper.GetActiveGitUser()
		if err != nil {
			fmt.Println("Error getting current Git user:", err)
			os.Exit(1)
		}
		if activeUsername != "" {
			fmt.Printf("\nActive Git user: %s <%s>\n", activeUsername, activeEmail)
		} else {
			fmt.Println("\nNo active Git user set.")
		}

	},
}

func init() {
	rootCmd.AddCommand(currentUser)
}
