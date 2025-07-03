package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.0.28:9000") // sunucu adresi
	if err != nil {
		fmt.Println("Bağlantı hatası:", err)
		return
	}
	defer conn.Close()

	fmt.Println("🔌 Sunucuya bağlandı. Komut gir (SET, GET, DEL, EXIT):")

	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewScanner(conn)

	for {
		fmt.Print(">>> ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		if cmd == "" {
			continue
		}

		// komutu gönder
		_, err := conn.Write([]byte(cmd + "\n"))
		if err != nil {
			fmt.Println("Yazma hatası:", err)
			break
		}

		// sunucudan cevabı oku
		if serverReader.Scan() {
			response := strings.TrimSpace(serverReader.Text())
			if response != "" {
				fmt.Println("🧠 Sunucu:", response)
			} else {
				fmt.Println("⚠️ Boş cevap geldi.")
			}
		} else {
			fmt.Println("📴 Sunucudan cevap alınamadı.")
			break
		}

		// çıkış komutu
		if strings.ToUpper(cmd) == "EXIT" {
			break
		}
	}
}
