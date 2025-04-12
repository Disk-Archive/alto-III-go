package alto

import (
	"errors"
	"fmt"
	"net"
	"time"
)

type TcpConnection struct {
	IPAddress net.IP
	Port      int
}

func (t *TcpConnection) sendTcp(message string, timeout time.Duration) (string, error) {
	address := fmt.Sprintf("%s:%d", t.IPAddress, t.Port)

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error: %v", err))
	}
	defer conn.Close()

	_, err = conn.Write([]byte(message + "\n"))
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error: %v", err))
	}

	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error: %v", err))
	}

	return string(buffer[:n]), nil
}
