package system

import (
	"github.com/gin-gonic/gin"
	v1 "go_blog/api"
)

type DictRouter struct{}

func (d DictRouter) InitDictRouter(router *gin.RouterGroup) gin.IRouter {
	dictRouter := router.Group("/dict")
	dictApi := v1.ApiGroup{}.SystemApiGroup.DictApi
	dictRouter.GET("/getList", dictApi.GetList)
	dictRouter.POST("/addDict", dictApi.AddDict)
	dictRouter.POST("/updateDict", dictApi.UpdateDict)
	dictRouter.GET("/deleteDict", dictApi.DeleteDict)
	return dictRouter
}
