package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "rbac-service",
	Short: "run rbac-service as cli program",
}

func init() {

}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}