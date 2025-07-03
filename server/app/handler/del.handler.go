package app

import (	
	"scrap-server/cache"
	code "scrap-server/error"
)

func Del(key string) int {
	if key != ""  {
		cache.Delete(key)
		return code.OK
	}
	return code.NotFound
}
