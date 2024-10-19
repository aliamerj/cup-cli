/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/aliamerj/cup-cli/pkg/apply"
	"github.com/spf13/cobra"
)

var filePath string

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
Use:   "apply",
Short: "Apply a configuration file to update Cloud-Cup settings dynamically.",
Long: `Apply a configuration from a file to the Cloud-Cup application.

This command allows you to dynamically load and apply configuration changes to your Cloud-Cup instance without restarting the service. The configuration will be validated, and if successful, the changes will be applied immediately.

Example usage:
  cup apply -f config.json

Options:
  -f, --file   Path to the configuration file to apply.

Cloud-Cup will automatically detect changes in the configuration and apply them through a hot reload mechanism, ensuring zero downtime.
`,	Run: func(cmd *cobra.Command, args []string) {
		if filePath != "" {
			apply.ApplyConfigFile(filePath)
			return
		}

		fmt.Printf("apply called without a file\n")
		cmd.Help()
	},
}

func init() {
	// Register the file flag with the apply command
	applyCmd.Flags().StringVarP(&filePath, "file", "f", "", "apply new config using the specified file path")
	rootCmd.AddCommand(applyCmd)
}
