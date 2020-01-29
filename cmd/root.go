package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "fetch data",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute executation of command
func Execute() {
	rootCmd.Execute()
}
