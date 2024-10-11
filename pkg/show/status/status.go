package status

import (
	"github.com/aliamerj/cup-cli/internal/style"
	"github.com/aliamerj/cup-cli/internal/uds"
)

// ShowStatus sends a request to check the status of Cloud-Cup
func ShowStatus() {
	// Attempt to establish a connection with the UDS (Unix Domain Socket)
	conn, err := uds.EstablishConnectionUDS()
	if err != nil {
		style.DisplayError("Cloud-Cup is currently offline ⚠️", nil)
		return
	}
	defer conn.Close()

	// Send the "show-status" request to the server
	if err := uds.SendRequest(conn, "show-status\n"); err != nil {
		style.DisplayError("Unable to reach Cloud-Cup 💔", err)
		return
	}

	// Show a modern loading animation while checking status
	style.ShowLoading("Syncing with Cloud-Cup servers...")

	// Prepare to read the response
	responseData := make([]byte, 1024)
	n, err := conn.Read(responseData)
	if err != nil {
		style.DisplayError("Connection lost while checking status 💀", err)
		return
	}

	// If no data is received, handle it gracefully
	if n == 0 {
		style.DisplayError("No response from Cloud-Cup 🛑", nil)
		return
	}

	// Display the server's response in a modern, stylish format
	style.DisplaySuccess("All systems go 🚀 Cloud-Cup is up and running!", string(responseData[:n]))
}
