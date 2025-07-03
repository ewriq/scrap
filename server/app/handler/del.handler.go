package app

import (	
	"scrap-server/cache"
	code "scrap-server/error"
)

func Del(key string) string {
	if key != ""  {
		cache.Delete(key)
		return string(code.OK)
	}
	return string(code.NotFound)
}
