package cache

import (
	"scrap-server/model"
)

var globalCache *model.Cache

func init() {
	globalCache = &model.Cache{
		Data: make(map[string]string),
	}
}

