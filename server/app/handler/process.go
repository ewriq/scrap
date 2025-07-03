package app

import (
	code "scrap-server/error"
	"strings"
)

func process(input string) string {
	parts := strings.SplitN(strings.TrimSpace(input), " ", 3)
	if len(parts) == 0 {
		return "❌ Geçersiz komut"
	}

	cmd := strings.ToUpper(parts[0])
	switch cmd {
	case "SET":

	if len(parts) != 3 {
		return "❌ SET <key> <value>"
	}


	err := Set(parts[1], parts[2])
	if err != string(code.OK) {

		return string(code.NotFound)
	}

	return string(code.OK)
	case "GET":
		if len(parts) != 2 {
			return "❌ GET <key>"
		}


		err := Get(parts[1])
		if err != string(code.NotFound) {
			return err
		}
		return string(code.NotFound)

	case "DEL":
		if len(parts) != 2 {
			return "❌ DEL <key>"
		}
		Del(parts[1])
	return "🗑️ Deleted"
	default:
		return "❓ Bilinmeyen komut"
	}
}
