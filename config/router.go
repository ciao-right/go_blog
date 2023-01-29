package config

import (
	"github.com/gin-gonic/gin"
	middleware "go_blog/middleware/jwt"
	"go_blog/router"
)

func Routers() *gin.Engine {
	r := gin.Default()
	systemRouter := new(router.RoutersGroupApp).System
	PublicGroup := r.Group("")
	{
		systemRouter.InitBaseRouter(PublicGroup)
	}
	PrivateGroup := r.Group("")
	PrivateGroup.Use(middleware.Jwt())
	{
		systemRouter.InitGoodsRouter(PrivateGroup)
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitDeptRouter(PrivateGroup)
	}

	return r
}
