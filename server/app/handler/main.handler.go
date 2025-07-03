package app

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func Connection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewScanner(conn)

	for {
		if !reader.Scan() {
			break 
		}

		input := strings.TrimSpace(reader.Text())

		if strings.ToUpper(input) == "EXIT" {
			conn.Write([]byte("👋 Bağlantı kapatıldı.\n"))
			return
		}

		response := process(input)

		conn.Write([]byte(fmt.Sprintf("%s\n", response)))
	}
}
