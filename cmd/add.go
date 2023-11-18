package cmd

import (
	"github.com/nabadeep25/gswitch/pkg/user"
	"github.com/spf13/cobra"
)

var (
	username string
	email    string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add user to list",
	Long:  `add user to git list`,
	Run: func(cmd *cobra.Command, args []string) {
		user.AddUser(username, email)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&username, "username", "u", "", "Git username")
	addCmd.Flags().StringVarP(&email, "email", "e", "", "Git email address")
	addCmd.MarkFlagRequired("username")
	addCmd.MarkFlagRequired("email")
}
