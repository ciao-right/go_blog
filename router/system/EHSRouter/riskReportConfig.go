package EHSRouter

import (
	"github.com/gin-gonic/gin"
	v1 "go_blog/api"
)

type RiskReportConfig struct{}

func (r RiskReportConfig) InitRiskReportConfig(Router *gin.RouterGroup) (R gin.IRouter) {
	riskReportConfigRouter := Router.Group("/riskReportConfig")
	baseApi := new(v1.ApiGroup).SystemApiGroup.RiskReportConfigApi
	{
		riskReportConfigRouter.POST("/add", baseApi.Add)
		riskReportConfigRouter.POST("/getList", baseApi.GetList)
		riskReportConfigRouter.POST("/update", baseApi.Update)
		riskReportConfigRouter.GET("/delete", baseApi.Delete)
	}
	return riskReportConfigRouter

}
