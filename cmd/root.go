package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gswitch",
	Short: "It is a tool to switch git user easily",
	Long:  ` It is a tool to switch git user easily`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
