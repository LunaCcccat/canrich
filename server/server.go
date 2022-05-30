package server

import (
	"CanRich/casbin"
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
	//初始化数据库
	db.InitDB()
	//初始化路由
	CanRichServer.InitRouter()
	//初始化casbin
	err := casbin.InitCasbin()
	if err != nil {
		panic(err)
	}
	//初始化redis
	//err := cache.InitRedisClient()
	//if err != nil {
	//	panic(err)
	//}
	//初始化邮件工具
	//err = utils.InitEmailSender(10)
	//if err != nil {
	//	panic(err)
	//}
}

func (server *Server) Run() {
	_ = server.Router.Run(":" + config.GlobalConfig.GetString("server.port"))
}
