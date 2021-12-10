package cmd

import "github.com/spf13/cobra"

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
