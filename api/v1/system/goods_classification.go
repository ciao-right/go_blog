package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go_blog/common/global"
	"go_blog/model/request"
	service "go_blog/service/system"
	"go_blog/utils"
	"net/http"
)

type GoodsClassApi struct{}

type goodsClassList []request.GoodsClassification

// ConvertToINodeArray 将当前数组转换成父类 INode 接口 数组
func (gs goodsClassList) ConvertToINodeArray() (nodes []utils.INode) {
	for _, v := range gs {
		nodes = append(nodes, v)
	}
	return
}

func (g GoodsClassApi) AddGoodsClass(ctx *gin.Context) {
	var goodsClass request.GoodsClassification
	err := ctx.ShouldBindJSON(&goodsClass)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ResStruct(201, "错误参数", nil))
		return
	}
	v := validator.New()
	sErr := v.Struct(goodsClass)
	if sErr != nil {
		ctx.JSON(http.StatusBadRequest, utils.ResStruct(201, fmt.Sprint("/s", sErr), nil))
		return
	}
	logic := service.GoodsClassService{}
	createErr := logic.CreateGoodsClass(goodsClass)
	if createErr != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ResStruct(500, utils.ErrToString(createErr), nil))
		return
	}
	ctx.JSON(http.StatusOK, utils.ResStruct(200, "添加成功", 1))
}

func (g GoodsClassApi) UpdateGoodsClass(c *gin.Context) {
	var gc request.GoodsClassification
	bindJsonErr := c.ShouldBindJSON(&gc)
	if bindJsonErr != nil {
		c.JSON(http.StatusBadRequest, utils.ResStruct(201, "参数错误", nil))
		return
	}
	logic := service.GoodsClassService{}
	v := validator.New()
	vErr := v.Struct(gc)
	if vErr != nil {
		c.JSON(http.StatusBadRequest, utils.ResStruct(201, utils.ErrToString(vErr), nil))
		return
	}
	updateErr := logic.UpdateGoodsClass(gc)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, utils.ResStruct(500, utils.ErrToString(updateErr), nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResStruct(200, "success", 1))
}

func (g GoodsClassApi) DelGoodsClass(c *gin.Context) {
	id := c.Query("id")
	logic := service.GoodsClassService{}
	delErr := logic.DelGoodsClass(id)
	if delErr != nil {
		c.JSON(http.StatusOK, utils.ResStruct(201, utils.ErrToString(delErr), nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResStruct(200, global.SUCCESS, 1))

}

func (g GoodsClassApi) GetGoodsClassList(c *gin.Context) {
	logic := service.GoodsClassService{}
	list := logic.GetGoodsClass()

	resp := utils.GenerateTree(goodsClassList.ConvertToINodeArray(list), nil)
	c.JSON(http.StatusOK, utils.ResStruct(200, "success", resp))

}
