package status

import (
	"fmt"
	"net"
)

// ShowStatus sends a request to check the status of Cloud-Cup
func ShowStatus() error {
	socketPath := "/tmp/cloud-cup.sock"

	// Establish connection to the Unix domain socket
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Send the "show-status" request to the server
	request := "show-status"
	if _, err := conn.Write([]byte(request)); err != nil {
		return err
	}

	// Prepare to read the response
	responseData := make([]byte, 1024)
	n, err := conn.Read(responseData)
	if err != nil {
		return err
	}

	// Display the server's response
	fmt.Println(string(responseData[:n]))
	return nil
}
