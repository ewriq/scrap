package main

import (
	Handler "scrap-server/app/handler"
	"scrap-server/app"
)

func main() {

	
	app.Start(":9000", Handler.Connection)
}