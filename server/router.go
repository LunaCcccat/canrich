package server

import (
	"CanRich/api/testapi"
	v1 "CanRich/api/v1"
	"CanRich/middleware"
	"github.com/gin-gonic/gin"
)

func (server *Server) InitRouter() {
	r := gin.New()
	//全局中间件注册
	r.Use(gin.Logger(), gin.Recovery(), middleware.Cors())

	visitor := r.Group("/api/v1")
	{
		visitor.POST("/login", v1.Login)
		visitor.POST("/register/:code", v1.Register)
	}

	test := r.Group("/api/test")
	{
		test.POST("/token", testapi.TokenTest, middleware.JwtMiddleware())
		test.POST("/casbin", testapi.CasbinTest)
	}

	server.Router = r
}
