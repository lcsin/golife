package biz

import (
	"github.com/gin-gonic/gin"
)

func Query(c *gin.Context) (string, error) {
	// query for database ...
	return "biz example data", nil
}
