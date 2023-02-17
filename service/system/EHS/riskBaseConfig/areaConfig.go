package riskBaseService

import (
	"go_blog/common/global"
	"go_blog/model/request/riskBaseConfig"
	"go_blog/utils"
)

type RiskService struct{}

func (RiskService) AddAreaConfig(config riskBaseConfig.RiskAreaConfig) error {
	result := global.GLOBAL_DB.Create(&config)
	return result.Error
}

func (RiskService) GetAreaConfigList(page, limit int, condition riskBaseConfig.RiskAreaConfig) (list []riskBaseConfig.RiskAreaConfig, err error) {
	result := global.GLOBAL_DB.Offset(utils.GetPage(page, limit)).Limit(limit).Find(&list, condition)
	return list, result.Error
}

func (RiskService) GetAreaConfigListCount(condition riskBaseConfig.RiskAreaConfig) (count int64, err error) {
	result := global.GLOBAL_DB.Model(&riskBaseConfig.RiskAreaConfig{}).Where(condition).Count(&count)
	return count, result.Error
}

func (RiskService) UpdateAreaConfig(config riskBaseConfig.RiskAreaConfig) error {
	result := global.GLOBAL_DB.Save(&config)
	return result.Error
}

func (RiskService) DeleteAreaConfig(id int) error {
	result := global.GLOBAL_DB.Delete(&riskBaseConfig.RiskAreaConfig{}, id)
	return result.Error
}
