/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display information about the Cloud-Cup server and its components.",
	Long: `The "show" command provides detailed information about the Cloud-Cup reverse proxy server. 
Currently, it supports checking whether the Cloud-Cup server is running or not.

Planned commands include:
- **show status**: Check if the Cloud-Cup server is running or stopped. This command is already implemented and provides a simple output like "Cloud-Cup is running" or "Cloud-Cup is NOT running."
  
- **show backends**: List all registered backend servers and their statuses (e.g., active, down, retrying). This would allow users to monitor their backend health and performance.

- **show -b [backend_id]**: Show detailed information about a specific backend, including:
  - Backend host and port
  - Current health status (healthy, failed)
  - Maximum failures before marking as unhealthy
  - Traffic stats (number of requests forwarded, errors, etc.)
  
- **show routes**: Display all configured routes, including their mapping to backend servers. This command will help users verify their route configurations and diagnose any misconfigurations.

- **show config**: Output the current server configuration (root address, load-balancing strategy, SSL/TLS settings, etc.). This helps in verifying that the right configuration is loaded.

- **show connections**: Show all active connections to the server, including:
  - Client IPs
  - Protocols (HTTP/2, HTTP/3, gRPC, WebSocket)
  - Total connections count
  - Any dropped or retried connections
  
- **show traffic**: Show traffic statistics for the server and backends, like:
  - Requests per second
  - Total requests handled
  - Total data transferred
  - Request distribution across backends (useful for load balancing analysis)
  
- **show ssl**: Display SSL/TLS certificate information:
  - Active certificates (expiry, issuer, etc.)
  - Supported protocols (HTTP/2, HTTP/3)
  - Status of SSL handshake for incoming connections
  
- **show errors**: Display recent server or backend errors, giving insight into what might be causing issues in the system (e.g., connection timeouts, SSL errors, backend failures).

These commands will make the "show" functionality comprehensive and essential for monitoring and managing the Cloud-Cup reverse proxy in production environments. It allows users to diagnose problems quickly and ensures smoother operations.`}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
