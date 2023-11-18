package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gswitch",
	Short: "switch global git user",
	Long: `The Git User Manager gswitch is a  convenient utility designed to 
simplify the process of managing global Git user configurations on your 
system. Whether you're working on multiple projects with different Git 
accounts or collaborating with various teams, gswitch provides an intuitive
and efficient way to add, list, and select Git users globally.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
