package cmd

import (
	"github.com/nabadeep25/gswitch/pkg/user"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set global git user from the list",
	Long:  `set global git user from the list`,
	Run: func(cmd *cobra.Command, args []string) {
		user.SelectUser()
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

}
