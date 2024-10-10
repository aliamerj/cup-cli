/*
Copyright © 2024 aliamerj
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cup",
	Short: "Cup CLI tool for managing and interacting with Cloud-Cup reverse proxy",
	Long: `Cup is a command-line interface (CLI) tool designed to manage and interact with Cloud-Cup, 
a reverse proxy for cloud-native microservices. With the Cup CLI, users can monitor server 
status, reload configurations, and perform essential administrative tasks without downtime.

This tool simplifies managing your Cloud-Cup proxy by offering commands to check server health, 
restart or reload services, and update configuration files seamlessly.

Key features include:
1. **Server Management**
   - **Check server status**: Easily check whether Cloud-Cup is running or stopped.
   - **Start, stop, and restart**: Commands to start or stop the Cloud-Cup service, as well as gracefully restart.
   - **Hot reload configurations**: Reload the server's configuration on-the-fly without requiring a restart, ensuring zero downtime during updates.
   - **View server logs**: Output and filter server logs to help diagnose issues and monitor real-time traffic.

2. **Configuration Management**
   - **Edit or apply configurations**: Apply new configuration files or edit existing settings (e.g., backends, routes, SSL) without service disruption.
   - **Support for multiple environments**: Seamlessly switch between production, staging, and development configurations.

3. **Backend Management**
   - **Monitor backends**: View the status of backend servers (e.g., active, failed, retrying).
   - **Add/remove backends**: Dynamically register new backend servers or remove unhealthy ones from the configuration.
   - **Health checks**: Monitor backend health status and traffic distribution for load balancing.

4. **Traffic Management**
   - **View traffic statistics**: Display real-time stats such as total requests, requests per second, and data transferred across the proxy and backends.
   - **Rate limiting**: Set rate limits to prevent traffic overloads on specific routes or backends.
   - **View active connections**: Monitor client connections, protocols (HTTP/2, HTTP/3), and current traffic load.

5. **SSL/TLS and Security**
   - **Set up SSL/TLS**: Configure and manage SSL certificates for secure HTTP/2 and HTTP/3 connections.
   - **View SSL/TLS status**: Display information about active certificates, supported protocols, and expiration dates.
   - **Basic authentication**: Add basic authentication for routes to secure access to certain endpoints.

6. **Kubernetes and Docker Integration**
   - **Automatic service discovery**: Seamlessly integrate with Kubernetes and Docker for dynamic service discovery and traffic routing.
   - **Deploy in Kubernetes/Docker**: Easily deploy Cloud-Cup in containerized environments with minimal configuration.

7. **Advanced Traffic Control** (Planned)
   - **Circuit breaking**: Implement circuit breaking to automatically stop sending requests to failing backends.
   - **Retries and fallback**: Set up retry policies and fallback mechanisms for failed requests.

8. **CLI Usability**
   - **Customizable output**: Customize CLI output formats (e.g., JSON, YAML) for easy integration with automation tools.
   - **Command completion**: Enable tab completion to make it easier to navigate through commands and options.
   - **Help and documentation**: Access detailed documentation and examples for each command via "cup help".

Cup empowers developers and operations teams to efficiently manage Cloud-Cup, ensuring smooth operation 
in production environments and making administration as simple as possible.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cup-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
