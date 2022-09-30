package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golife/gin-layout/api/biz"
)

func registryBizRouter(r *gin.RouterGroup) {
	v1 := r.Group("/biz")
	{
		v1.GET("", biz.Get)
		// ...
	}
}
