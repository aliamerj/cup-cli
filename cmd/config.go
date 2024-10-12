/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/aliamerj/cup-cli/pkg/show/config"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Display the current Cloud-Cup configuration",
	Long: `Display detailed information about the current Cloud-Cup server configuration, 
including root server settings, routing rules, and backend strategies. 
`,
	Run: func(cmd *cobra.Command, args []string) {
		config.ShowConfig()
	},
}

func init() {
	showCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
