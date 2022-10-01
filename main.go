package main

import (
	"api/middware"
	"api/router"
	"flag"
	"github.com/gin-gonic/gin"
)

func main() {
	var address string
	flag.StringVar(&address, "address", ":8080", "server address")
	flag.Parse()
	// 初始化gin
	r := gin.Default()
	// 注册中间件
	r.Use(middware.JsonWebToken())
	// 引入router下的路由
	router.Router.InitApiRouter(r)
	// 启动服务
	r.Run(address)
}
