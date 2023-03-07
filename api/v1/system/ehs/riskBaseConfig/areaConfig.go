package riskBaseConfig

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go_blog/model/request/riskBaseConfig"
	riskBaseService "go_blog/service/system/EHS/riskBaseConfig"
	"go_blog/utils"
)

type EHSApi struct{}

func (EHSApi) AddAreaConfig(c *gin.Context) {
	var config riskBaseConfig.RiskAreaConfig
	err := c.ShouldBindJSON(&config)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}
	v := validator.New()
	err = v.Struct(config)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}
	vErr := new(riskBaseService.RiskService).AddAreaConfig(config)
	if vErr != nil {
		utils.ErrorRes(vErr, c)
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    1,
		"code":    200,
	})

}

func (EHSApi) GetAreaConfigList(c *gin.Context) {
	config := riskBaseConfig.RiskAreaConfig{}
	err := c.ShouldBindJSON(&config)
	if err != nil {
		return
	}
	page := utils.StringToInt(c.PostForm("page"))
	limit := utils.StringToInt(c.PostForm("limit"))
	utils.InitPage(&page, &limit)
	logic := new(riskBaseService.RiskService)
	makeExcel := utils.StringToInt(c.PostForm("makeExcel"))
	if makeExcel == 1 {
		list, err := logic.GetAreaConfigList(page, limit, config)
		if err != nil {
			utils.ErrorRes(err, c)
			return
		}
		utils.GetStructLabel(config)
		fileUrl := utils.MakeExcel(list)
		c.JSON(200, gin.H{
			"message": "success",
			"data":    fileUrl,
			"code":    200,
		})
		return
	} else {
		list, err := logic.GetAreaConfigList(page, limit, config)
		if err != nil {
			utils.ErrorRes(err, c)
			return
		}
		count, err := logic.GetAreaConfigListCount(config)
		if err != nil {
			utils.ErrorRes(err, c)
			return
		}
		data := make(map[string]interface{})
		data["list"] = list
		data["total"] = count
		c.JSON(200, gin.H{
			"message": "success",
			"data":    data,
			"code":    200,
		})
	}

}

func (EHSApi) UpdateAreaConfig(c *gin.Context) {
	var config riskBaseConfig.RiskAreaConfig
	err := c.ShouldBindJSON(&config)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}
	v := validator.New()
	err = v.Struct(config)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}
	vErr := new(riskBaseService.RiskService).UpdateAreaConfig(config)
	if vErr != nil {
		utils.ErrorRes(vErr, c)
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    1,
		"code":    200,
	})
}

func (EHSApi) DeleteAreaConfig(c *gin.Context) {
	id := c.Query("id")
	vErr := new(riskBaseService.RiskService).DeleteAreaConfig(utils.StringToInt(id))
	if vErr != nil {
		utils.ErrorRes(vErr, c)
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    1,
		"code":    200,
	})
}
