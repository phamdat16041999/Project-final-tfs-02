package main

import (
	"hotel/migrate"
	"hotel/pkg/websocket"
	"hotel/router"
)

func main() {
	websocket.ChatAPI()
	migrate.CreateTable()
	router.Run()
}
