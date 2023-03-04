package riskBaseService

import (
	"fmt"
	"go_blog/common/global"
	"go_blog/model/request"
	"go_blog/model/request/riskBaseConfig"
	"go_blog/utils"
)

type RiskService struct{}

func (RiskService) AddAreaConfig(config riskBaseConfig.RiskAreaConfig) error {
	result := global.GLOBAL_DB.Create(&config)
	return result.Error
}

type RiskAreaConfigRes struct {
	riskBaseConfig.RiskAreaConfig
	Dept request.Department `json:"responsible_department" gorm:"foreignKey:id;"`
	User request.User       `json:"responsible_person" gorm:"foreignKey:id;"`
}

func (RiskService) GetAreaConfigList(page, limit int, condition riskBaseConfig.RiskAreaConfig) (list []RiskAreaConfigRes, err error) {

	result := global.GLOBAL_DB.Preload("Dept", "state NOT IN (?)", 0).Preload("User").Offset(utils.GetPage(page, limit)).Limit(limit).Find(&list, condition)
	fmt.Println(result)
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
