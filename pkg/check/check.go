package check

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	command "github.com/aliamerj/cup-cli/internal/cmd"
	"github.com/aliamerj/cup-cli/internal/style"
	"github.com/fatih/color"
)


type BackendHealth struct {
	Host         string `json:"host"`
	Route        string `json:"route"`
	IsHealthy    bool   `json:"is_healthy"`
	ErrorMessage string `json:"error_message"`
	ErrorName    string `json:"error_name"`
}

func CheckAll() {
	executeHealthCheck(nil, nil)
}

func CheckRoutes(routes []string) {
	executeHealthCheck(routes, nil)
}

func CheckBackends(backends []string) {
	executeHealthCheck(nil, backends)
}

func executeHealthCheck(routes, backends []string) {
	printBanner()

	response, err := command.CmdCall("check-all")
	if err != nil {
		handleError("Failed to fetch backend data")
	}

	var backendHealths []BackendHealth
	if err := json.Unmarshal([]byte(response), &backendHealths); err != nil {
		handleError("Failed to parse backend data")
	}

	if routes != nil {
		displayByFilter(backendHealths, routes, func(b BackendHealth) string { return b.Route })
	} else if backends != nil {
		displayByFilter(backendHealths, backends, func(b BackendHealth) string { return b.Host })
	} else {
		groupByAndDisplay(backendHealths, func(b BackendHealth) string { return b.Route })
	}

	fmt.Println("\n✅ Health checks completed successfully.")
}

func handleError(message string) {
	fmt.Println(color.RedString("❌ " + message))
	os.Exit(1)
}

func displayByFilter(backendHealths []BackendHealth, filters []string, keyFunc func(BackendHealth) string) {
	grouped := groupBy(backendHealths, keyFunc)
	for _, filter := range filters {
		if backends, exists := grouped[filter]; exists {
			displayGroupedBackends(filter, backends)
		} else {
			fmt.Printf("\n%s %s %s\n", color.RedString("❌"), color.BlueString(filter), color.RedString("does not exist or has no backends"))
		}
	}
}

func groupByAndDisplay(backendHealths []BackendHealth, keyFunc func(BackendHealth) string) {
	grouped := groupBy(backendHealths, keyFunc)
	for key, backends := range grouped {
		displayGroupedBackends(key, backends)
	}
}

func groupBy(backendHealths []BackendHealth, keyFunc func(BackendHealth) string) map[string][]BackendHealth {
	group := make(map[string][]BackendHealth)
	for _, backend := range backendHealths {
		key := keyFunc(backend)
		group[key] = append(group[key], backend)
	}
	return group
}

func displayGroupedBackends(key string, backends []BackendHealth) {
	bold := color.New(color.Bold).SprintFunc()
	blue := color.New(color.FgHiBlue).SprintFunc()

	fmt.Printf("\n%s %s\n", bold("📍 Route:"), blue(key))
	for _, backend := range backends {
		fmt.Printf("   🔗 Host: %s\n", bold(backend.Host))
		displayHealthStatus(backend)
	}
}

func displayHealthStatus(backend BackendHealth) {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	if backend.IsHealthy {
		fmt.Printf("   Status: %s\n", green("HEALTHY ✅"))
	} else {
		fmt.Printf("   Status: %s\n", red("UNHEALTHY ❌"))
		fmt.Printf("   Error: %s - %s\n", red(backend.ErrorMessage), red(backend.ErrorName))
	}
}

func printBanner() {
  style.PrintBanner()
	bold := color.New(color.Bold).SprintFunc()
	fmt.Println(bold("🚀 Cloud-Cup Backend Health Check"))
	fmt.Println(bold("-------------------------------"))
	time.Sleep(1 * time.Second)
	fmt.Println(color.GreenString(" Done! ✅\n"))
	fmt.Println("⏳ Running backend health checks...")
}

