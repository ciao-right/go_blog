package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_blog/model/request"
	service "go_blog/service/system"
	"go_blog/utils"
	"net/http"
	"strconv"
)

type DictApi struct{}

func (d DictApi) GetList(c *gin.Context) {
	dict := request.Dict{
		DictName: c.Query("dictName"),
		Status:   utils.StringToInt(c.Query("status")),
		ParentId: utils.StringToInt(c.Query("parentId")),
	}
	// 字符串转int
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	logic := service.DictService{}
	err, list := logic.GetList(page, limit, dict)
	total, err := logic.GetTotal(dict)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": gin.H{
			"list":  list,
			"total": total,
		},
	})
}

func (d DictApi) AddDict(c *gin.Context) {
	var dict request.Dict

	err := c.ShouldBindJSON(&dict)
	fmt.Println(dict)
	fmt.Println(err)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}
	fmt.Println(err)
	logic := service.DictService{}
	err = logic.AddDict(dict)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    "",
	})
}

func (d DictApi) UpdateDict(c *gin.Context) {
	var dist request.Dict
	err := c.ShouldBindJSON(&dist)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}
	logic := service.DictService{}
	err = logic.UpdateDict(dist)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    "",
	})
}

func (d DictApi) DeleteDict(c *gin.Context) {
	var dict request.Dict
	err := c.ShouldBindJSON(&dict)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}
	logic := service.DictService{}
	err = logic.DeleteDict(dict)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    "",
	})
}
