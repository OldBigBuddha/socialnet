package cmd

import (
	"github.com/spf13/cobra"
	"log/slog"
)

var rootCmd = &cobra.Command{
	Use:   "account",
	Short: "account service",
	RunE: func(cmd *cobra.Command, args []string) error {
		slog.Info("Hello World")
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}
