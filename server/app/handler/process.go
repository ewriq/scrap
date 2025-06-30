package app

import (
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
		set(parts[1], parts[2])
		return "✅ OK"

	case "GET":
		if len(parts) != 2 {
			return "❌ GET <key>"
		}
		get(parts[1])
		return "⛔ Not Found"

	case "DEL":
		if len(parts) != 2 {
			return "❌ DEL <key>"
		}
		del(parts[1])
		return "🗑️ Deleted"

	case "EXIT":
		return "EXIT"

	default:
		return "❓ Bilinmeyen komut"
	}
}
