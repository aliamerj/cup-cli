package uds

import "net"

func EstablishConnectionUDS() (net.Conn, error) {
	socketPath := "/tmp/cloud-cup.sock"

	// Establish connection to the Unix domain socket
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func SendRequest(conn net.Conn, cmd string) error {
	if _, err := conn.Write([]byte(cmd)); err != nil {
		return err
	}
  return nil;
}
