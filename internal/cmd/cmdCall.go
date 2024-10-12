package command

import "github.com/aliamerj/cup-cli/internal/uds"

func CmdCall(cmd string) (string, error) {
	conn, err := uds.EstablishConnectionUDS()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	if err := uds.SendRequest(conn, cmd+"\n"); err != nil {
		return "", err
	}
	responseData := make([]byte, 1024*1024)
	n, err := conn.Read(responseData)
	if err != nil {
		return "", err
	}
	return string(responseData[:n]), nil
}
