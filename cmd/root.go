package cmd

import (
	"github.com/OldBigBuddha/socialnet/service"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "account",
	Short: "account service",
	RunE: func(cmd *cobra.Command, args []string) error {
		return service.Run()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
