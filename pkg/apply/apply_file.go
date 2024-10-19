package apply

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	command "github.com/aliamerj/cup-cli/internal/cmd"
	"github.com/aliamerj/cup-cli/internal/style"
	"github.com/fatih/color"
)

func ApplyConfigFile(filePath string) {
	printBanner()

	// Convert to absolute path
	absPath, err := resolveAbsolutePath(filePath)
	if err != nil {
		handleError(fmt.Sprintf("Unable to resolve the absolute path for: %s", filePath))
	}

	// Check if file exists
	if !fileExists(absPath) {
		handleError(fmt.Sprintf("Configuration file not found: %s", absPath))
	}

	// Send command to Cloud-Cup
	commandStr := fmt.Sprintf("apply-config|%s", absPath)
	response, err := command.CmdCall(commandStr)
	if err != nil {
    handleError(fmt.Sprintf("Failed to apply the configuration file: %v", err))
	}

	// Success message
	fmt.Println(color.GreenString("✅ Configuration applied successfully"))
	fmt.Println(color.BlueString("Response from Cloud-Cup:"))
	fmt.Println(response)
}

func resolveAbsolutePath(filePath string) (string, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return "", err
	}
	return absPath, nil
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func handleError(message string) {
	fmt.Println(color.RedString("❌ " + message))
	os.Exit(1)
}

func printBanner() {
	style.PrintBanner()
	bold := color.New(color.Bold).SprintFunc()

	fmt.Println(bold("🚀 Cloud-Cup Configuration Apply"))
	fmt.Println(bold("-------------------------------"))
	time.Sleep(1 * time.Second)
	fmt.Println("⏳ Applying configuration file...")
}
