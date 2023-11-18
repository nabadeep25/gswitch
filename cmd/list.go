/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/nabadeep25/gswitch/pkg/user"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display user list",
	Long:  `Display the list is added user`,
	Run: func(cmd *cobra.Command, args []string) {
		user.ListUsers()

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
