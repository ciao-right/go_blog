package riskBaseConfig

import "go_blog/model"

type RiskAreaConfig struct {
	model.BaseModel
	AreaId                string `json:"area_id" gorm:"column:area_id;comment:区域ID"`
	RiskLevel             string `json:"risk_level" gorm:"column:risk_level;comment:风险等级"`
	ResponsibleDepartment string `json:"responsible_department" gorm:"column:responsible_department;comment:责任部门"`
	ResponsiblePerson     string `json:"responsible_person" gorm:"column:responsible_person;comment:责任人"`
	Remark                string `json:"remark" gorm:"column:remark;comment:备注"`
}

func (RiskAreaConfig) TableName() string {
	return "risk_area_config"
}
