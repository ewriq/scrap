package app

import (
	"scrap-server/cache"
	code "scrap-server/error"
)

func Set(key, value string) string  {
	if key != "" && value != "" {
		cache.Set(key,value)
		return string(code.OK)
	}
	return string(code.NotFound)
}
