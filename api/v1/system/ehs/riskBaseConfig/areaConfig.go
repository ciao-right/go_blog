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

	c.JSON(200, gin.H{
		"message": "GetAreaConfigList",
	})
}
