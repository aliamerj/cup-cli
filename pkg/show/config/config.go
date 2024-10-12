package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	command "github.com/aliamerj/cup-cli/internal/cmd"
	"github.com/aliamerj/cup-cli/internal/style"
	"github.com/fatih/color"
)

type Backend struct {
	Host       string `json:"host"`
	MaxFailure int    `json:"max_failure"`
}

type Route struct {
	Backends []Backend `json:"backends"`
	Strategy string    `json:"strategy,omitempty"`
}

type Config struct {
	Root   string           `json:"root"`
	Routes map[string]Route `json:"routes"`
}

func ShowConfig() {
	style.PrintBanner()

	blue := color.New(color.FgHiBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	fmt.Println(bold("🚀 Cloud-Cup Configuration"))
	// Simulate loading with a progress indicator
	fmt.Print("⏳ Loading configuration...")
	time.Sleep(1 * time.Second)
	fmt.Println(green(" Done! ✅\n"))

	data, err := command.CmdCall("show-config")
	if err != nil {
		fmt.Println(color.RedString("❌ Cloud-Cup is currently offline"))
		os.Exit(1)
	}

	var config Config
	errm := json.Unmarshal([]byte(data), &config)
	if errm != nil {
		fmt.Println(color.RedString("❌ Failed to parse config data: %v", err))
		os.Exit(1)
	}

	// Display the root server
	fmt.Printf("%s: %s\n", blue("🔗 Root Server"), green(config.Root))

	// Loop through routes and display in a more visually appealing way
	for route, routeData := range config.Routes {
		fmt.Printf("\n%s %s\n", bold("📍 Route:"), blue(route))
		for _, backend := range routeData.Backends {
			fmt.Printf("   🔗 %s (Max Failures: %s)\n", green(backend.Host), bold(fmt.Sprintf("%d", backend.MaxFailure)))
		}
		// Display strategy if present
		if routeData.Strategy != "" {
			fmt.Printf("   📊 Strategy: %s\n", blue(routeData.Strategy))
		}
	}

	// Final note
	fmt.Println("\n✅ Configuration loaded successfully.")

}
