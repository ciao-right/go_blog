package riskBaseService

import (
	"go_blog/common/global"
	"go_blog/model/request/riskBaseConfig"
)

type RiskService struct{}

func (RiskService) AddAreaConfig(config riskBaseConfig.RiskAreaConfig) error {
	result := global.GLOBAL_DB.Create(&config)
	return result.Error
}
