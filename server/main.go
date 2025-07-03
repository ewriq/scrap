package main

import (
	"fmt"
	"scrap-server/app"
	Handler "scrap-server/app/handler"
	"scrap-server/cache"
)

func main() {

	cache.Set("ewriq:1","31")
	val, ok:= cache.Get("ewriq:1")
	if ok {
		fmt.Println(val)
	}
	app.Start("192.168.0.28:9000", Handler.Connection)
}