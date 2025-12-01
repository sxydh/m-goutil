package port_utils

import (
	"fmt"
	"net"
)

func GetFreePort() (int, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return -1, err
	}
	defer func(listener net.Listener) {
		_ = listener.Close()
	}(listener)
	return listener.Addr().(*net.TCPAddr).Port, nil
}

func GetFreePortInRange(start, end int) (int, error) {
	if start < 1 || end > 65535 || start > end {
		return -1, fmt.Errorf("invalid port range: %d-%d", start, end)
	}

	for port := start; port <= end; port++ {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
			_ = listener.Close()
			return port, nil
		}
	}

	return -1, fmt.Errorf("no available port in range [%d, %d]", start, end)
}
