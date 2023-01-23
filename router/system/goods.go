package system

import (
	"github.com/gin-gonic/gin"
	v1 "go_blog/api"
)

type GoodsRouter struct{}

func (g GoodsRouter) InitGoodsRouter(r *gin.RouterGroup) gin.IRouter {
	goodsRouter := r.Group("/goods")
	goodsApi := new(v1.ApiGroup).SystemApiGroup.GoodsClassApi
	goodsRouter.GET("/getGoodClass", goodsApi.GetGoodsClassList)
	goodsRouter.POST("/addGoodsClass", goodsApi.AddGoodsClass)
	goodsRouter.GET("/delete", goodsApi.DelGoodsClass)
	goodsRouter.POST("/updateGoodsClass", goodsApi.UpdateGoodsClass)
	return goodsRouter
}
