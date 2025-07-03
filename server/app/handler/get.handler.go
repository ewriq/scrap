package app

import (
	"scrap-server/cache"
	code "scrap-server/error"
)

func Get(key string) string {
	val, ok:= cache.Get(key)
	if ok {
		return val
	} else {
		return string(code.NotFound)
	}
}
