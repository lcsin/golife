package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/golife/gin-layout/database/biz"
	"github.com/golife/gin-layout/log"
	"github.com/golife/gin-layout/pkg/response"
)

func Get(c *gin.Context) {
	// ...
	query, err := biz.Query(c)
	if err != nil {
		log.Errorf("error: %v", err)
		response.Error(c, 500, err.Error())
		return
	}
	response.OK(c, query, "ok")
}
