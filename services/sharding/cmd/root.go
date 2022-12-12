package cmd

import "github.com/spf13/cobra"

const dialect = "postgres"

var rootCmd = &cobra.Command{
	Use:   "service",
	Short: "run service as cli program",
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
