package app

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	code "scrap-server/error"
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

		response, data := process(input)
		if data != "" {
			if response != code.OK {
				conn.Write([]byte(fmt.Sprintf("%s\n", string(response))))
			} else {
				conn.Write([]byte(fmt.Sprintf("%s\n", data)))
			}
		}
	}
}
