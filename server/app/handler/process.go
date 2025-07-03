package app

import (
	code "scrap-server/error"
	"strings"
)

func process(input string) (int, string) {
	parts := strings.SplitN(strings.TrimSpace(input), " ", 3)
	if len(parts) == 0 {
		return code.InvalidCommand, ""
	}

	cmd := strings.ToUpper(parts[0])
	switch cmd {
	case "SET":

		if len(parts) != 3 {
			return code.Usage, ""
		}

		err := Set(parts[1], parts[2])
		if err != code.OK {

			return code.NotFound, ""
		}

		return code.OK, ""
	case "GET":
		if len(parts) != 2 {
			return code.Usage, ""
		}

		err := Get(parts[1])
		if err != string(code.NotFound) {
			return code.OK, err
		}
		return code.NotFound, ""
	case "DEL":
		if len(parts) != 2 {
			return code.Usage, ""
		}
		err := Del(parts[1])
		if err != code.OK {
			return code.NotFound, ""
		}
		return code.OK, ""
	default:
		return code.Undefiend, ""
	}
}
