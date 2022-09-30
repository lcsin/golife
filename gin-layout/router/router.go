package router

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	engine := gin.Default()
	engine.GET("/ping", ping)

	// 注册业务路由
	r := engine.Group("/api/v1")
	registryBizRouter(r)

	return engine
}

func ping(c *gin.Context) {
	_, _ = c.Writer.WriteString("ok")
}
