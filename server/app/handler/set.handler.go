package app

import (
	"scrap-server/cache"
	code "scrap-server/error"
)

func Set(key, value string) int  {
	if key != "" && value != "" {
		cache.Set(key,value)
		return code.OK
	}
	return code.NotFound
}
