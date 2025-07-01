package app

import (
	"bufio"

	"net"
	"strings"
	"sync"
)

var store = make(map[string]string)
var lock sync.RWMutex

func Connection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewScanner(conn)

	for reader.Scan() {
		input := strings.TrimSpace(reader.Text())
		parts := strings.SplitN(input, " ", 3)

		if len(parts) < 1 {
			conn.Write([]byte(""))
			continue
		}

		cmd := strings.ToUpper(parts[0])
		switch cmd {
		case "SET":
			if len(parts) != 3 {
				conn.Write([]byte(""))
				continue
			}
			lock.Lock()
			store[parts[1]] = parts[2]
			lock.Unlock()
			conn.Write([]byte(""))
			return

		case "GET":
			if len(parts) != 2 {
				conn.Write([]byte(""))
				continue
			}
			conn.Write([]byte(""))
			return

		case "DEL":
			if len(parts) != 2 {
				conn.Write([]byte(""))
				continue
			}
			lock.Lock()
			delete(store, parts[1])
			lock.Unlock()
			conn.Write([]byte(""))
			return

		case "EXIT":
			conn.Write([]byte(""))
			return

		default:
			conn.Write([]byte(""))
			return
		}

	}
	
}
