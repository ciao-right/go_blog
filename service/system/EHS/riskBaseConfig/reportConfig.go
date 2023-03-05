package riskBaseService

import (
	"go_blog/common/global"
	"go_blog/model/request/riskBaseConfig"
)

type RiskReportConfigService struct{}

func (r RiskReportConfigService) AddRiskReportConfig(riskReport riskBaseConfig.RiskReportConfig) error {
	res := global.GLOBAL_DB.Create(&riskReport)
	return res.Error
}
