package main

import (
	"scrap-server/app"
	Handler "scrap-server/app/handler"
)

func main() {
	app.Start("192.168.0.28:9000", Handler.Connection)
}