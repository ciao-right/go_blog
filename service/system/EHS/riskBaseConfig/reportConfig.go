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

func (r RiskReportConfigService) GetRiskReportConfigList(page int, limit int, config riskBaseConfig.RiskReportConfig) (interface{}, interface{}) {
	return nil, nil
}
