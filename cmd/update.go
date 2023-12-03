package cmd

import (
	"github.com/nabadeep25/gswitch/pkg/user"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update a user from  list",
	Long:  `update a user from  list`,
	Run: func(cmd *cobra.Command, args []string) {
		user.UpdateUser()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

}
