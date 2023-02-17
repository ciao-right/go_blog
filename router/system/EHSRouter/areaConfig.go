package EHSRouter

import (
	"github.com/gin-gonic/gin"
	v1 "go_blog/api"
)

type EHSRouter struct{}

func (e EHSRouter) InitEHSRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	EHSRouter := Router.Group("/ehs/areaConfig")
	EhsService := new(v1.ApiGroup).SystemApiGroup.EHSApi
	{
		EHSRouter.POST("/add", EhsService.AddAreaConfig)
		EHSRouter.POST("/getList", EhsService.GetAreaConfigList)
		EHSRouter.POST("/update", EhsService.UpdateAreaConfig)
		EHSRouter.GET("/delete", EhsService.DeleteAreaConfig)
	}
	return EHSRouter
}
