/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (

	"github.com/aliamerj/cup-cli/pkg/show/status"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show whether Cloud-Cup is running or stopped.",
	Long: `The 'status' command displays the current operational state of the Cloud-Cup server.
It checks if the Cloud-Cup reverse proxy is actively running or if it's stopped.

Use this command to quickly verify if the proxy is up and running, ensuring that your services are being routed correctly.

Example usage:
  cup status
  
Output:
  - "Cloud-Cup is running" when the server is active.
  - "Cloud-Cup is NOT running" if the server is stopped.`,

	Run: func(cmd *cobra.Command, args []string) {
		status.ShowStatus()
	},
}

func init() {
	showCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
