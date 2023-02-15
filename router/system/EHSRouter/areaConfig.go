package EHSRouter

import (
	"github.com/gin-gonic/gin"
	v1 "go_blog/api"
)

type EHSRouter struct{}

func (e EHSRouter) InitEHSRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	EHSRouter := Router.Group("/ehs")
	EhsService := new(v1.ApiGroup).SystemApiGroup.EHSApi
	{
		EHSRouter.GET("/areaConfig", EhsService.GetAreaConfigList)
	}
	return EHSRouter
}
