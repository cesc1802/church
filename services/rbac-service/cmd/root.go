package cmd

import "github.com/spf13/cobra"

const dialect = "postgres"

var rootCmd = &cobra.Command{
	Use:   "rbac-service",
	Short: "run rbac-service as cli program",
}

func init() {
	rootCmd.AddCommand(migrationCmd)
	rootCmd.AddCommand(serverCmd)
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
