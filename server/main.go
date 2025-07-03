package main

import (
	"fmt"
	"scrap-server/app"
	Handler "scrap-server/app/handler"
	"scrap-server/cache"
)

func main() {
	app.Start("192.168.0.28:9000", Handler.Connection)
}