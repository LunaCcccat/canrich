package server

import (
	"CanRich/config"
	"CanRich/db"
	"github.com/gin-gonic/gin"
)

var CanRichServer *Server

type Server struct {
	Router *gin.Engine
}

func InitServer() {
	CanRichServer = &Server{}
	db.InitDB()
	CanRichServer.InitRouter()
	//err := cache.InitRedisClient()
	//if err != nil {
	//	panic(err)
	//}
	//err = utils.InitEmailSender(10)
	//if err != nil {
	//	panic(err)
	//}
}

func (server *Server) Run() {
	_ = server.Router.Run(":" + config.GlobalConfig.GetString("server.port"))
}
