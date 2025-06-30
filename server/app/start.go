package app

import (
	"fmt"
	"net"
	"os"
)

func Start(address string, handler func(conn net.Conn)) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Server exit:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Listening :", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection err:", err)
			continue
		}
		go handler(conn) 
	}
}
