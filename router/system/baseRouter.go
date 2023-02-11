package system

import (
	"github.com/gin-gonic/gin"
	v1 "go_blog/api"
)

type BaseRouter struct{}

func (b BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	baseRouter := Router.Group("base") //都是以/base 为开头
	baseApi := new(v1.ApiGroup).SystemApiGroup.BaseApi
	{
		baseRouter.POST("/login", baseApi.Login)
		baseRouter.POST("/register", baseApi.Register)
	}
	return baseRouter
}
