package EHSApi

import (
	"github.com/gin-gonic/gin"
	"go_blog/model/request/riskBaseConfig"
	riskBaseService "go_blog/service/system/EHS/riskBaseConfig"
	"go_blog/utils"
)

type RiskReportConfigApi struct {
}

func (r RiskReportConfigApi) Add(c *gin.Context) {
	var riskReportConfig riskBaseConfig.RiskReportConfig
	err := c.ShouldBindJSON(&riskReportConfig)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}
	logic := new(riskBaseService.RiskReportConfigService)
	err = logic.AddRiskReportConfig(riskReportConfig)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": "1",
	})

}

func (r RiskReportConfigApi) GetList(c *gin.Context) {
	var riskReportConfig riskBaseConfig.RiskReportConfig
	err := c.ShouldBindJSON(&riskReportConfig)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}
	page := utils.StringToInt(c.PostForm("page"))
	limit := utils.StringToInt(c.PostForm("limit"))
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	logic := new(riskBaseService.RiskReportConfigService)
	list, err := logic.GetRiskReportConfigList(page, limit, riskReportConfig)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": list,
	})
}

func (r RiskReportConfigApi) Update(c *gin.Context) {

}

func (r RiskReportConfigApi) Delete(c *gin.Context) {

}
