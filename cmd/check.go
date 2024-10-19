/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (

	"github.com/aliamerj/cup-cli/pkg/check"
	"github.com/spf13/cobra"
)

var routeFlag bool 
var backendFlag bool

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the health of backend servers",
	Long: `Check the health status of backend servers in the Cloud-Cup load balancer.
The command will display which backends are healthy and list any issues with those that are not.
  `,

	Run: func(cmd *cobra.Command, args []string) {

		if routeFlag {
			check.CheckRoutes(args)
      return
		}

    if backendFlag {
      check.CheckBackends(args)
      return
    }

		check.CheckAll()
	},
}

func init() {
	checkCmd.Flags().BoolVarP(&routeFlag, "route", "r", false, "Check the health of all backend servers in specific route")
	checkCmd.Flags().BoolVarP(&backendFlag, "backend", "b", false, "Check the health of specific backend")

  rootCmd.AddCommand(checkCmd)
}
