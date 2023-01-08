package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseApi struct {
}

func (b *BaseApi) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}
