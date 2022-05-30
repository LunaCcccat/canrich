package main

import (
	"CanRich/config"
	"CanRich/logger"
	"CanRich/server"
)

func main() {
	logger.InitLogger()
	config.InitConfig()
	server.InitServer()
	server.CanRichServer.Run()
}
