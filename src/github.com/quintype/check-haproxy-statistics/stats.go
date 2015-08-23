package haproxy

import (
	"net"
	"io"
)

func StatsStream(filename string) (io.ReadCloser, error) {
	conn, err := net.Dial("unix", filename)
	if (err != nil) {
		return nil, err
	}

	_, err = conn.Write([]byte("show stat\n"))
	if (err != nil) {
		conn.Close()
		return nil, err
	}

	n, err := conn.Read(make([]byte, 2))
	if (err != nil || n < 2) {
		conn.Close()
		return nil, err
	}

	return conn, err
}
