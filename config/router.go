package config

import (
	"github.com/gin-gonic/gin"
	"go_blog/router"
	"net/http"
)

func Routers() *gin.Engine {
	r := gin.Default()
	systemRouter := new(router.RoutersGroupApp).System
	PublicGroup := r.Group("")
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup)
	}

	return r
}
