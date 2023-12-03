package cmd

import (
	"github.com/nabadeep25/gswitch/pkg/user"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove a user from  list",
	Long:  `remove a user from  list`,
	Run: func(cmd *cobra.Command, args []string) {
		user.DeleteUser()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

}
