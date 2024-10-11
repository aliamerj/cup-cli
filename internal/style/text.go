package style

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

// Utility function to display error messages with a modern twist
func DisplayError(message string, err error) {
	red := color.New(color.FgRed).SprintFunc()
 fmt.Fprintf(os.Stderr, "%s\n", red(message))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Details: %v\n", err)
	}
}

// Utility function to display success messages
func DisplaySuccess(message string, status string) {
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("%s\n", green(message))
	fmt.Printf("Status: %s\n", status)
}

// Simulate a loading animation with a message
func ShowLoading(message string) {
	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Println(yellow(message))
	time.Sleep(1 * time.Second) // Simulate a brief delay for loading effect
}
