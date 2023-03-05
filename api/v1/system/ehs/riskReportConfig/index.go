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

}

func (r RiskReportConfigApi) Update(c *gin.Context) {

}

func (r RiskReportConfigApi) Delete(c *gin.Context) {

}
