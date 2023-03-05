package riskBaseConfig

import (
	"go_blog/model"
	"go_blog/model/request"
)

type ReportTimeList struct {
	StartLevel int `json:"startLevel" `
	EndLevel   int `json:"endLevel" `
	Value      int `json:"value" `
}

type ReportPostList struct {
	StartLevel int                `json:"startLevel" `
	EndLevel   int                `json:"endLevel" `
	DeptId     int                `json:"deptId" `
	Dept       request.Department `json:"dept" gorm:"foreignKey:DeptId;references:DeptId"`
}

type RiskReportConfig struct {
	model.BaseModel
	AreaId   int              `json:"areaId" gorm:"column:area_id;comment:区域id"`
	Area     RiskAreaConfig   `json:"area" gorm:"foreignKey:AreaId;references:AreaId"`
	Remark   string           `json:"remark" gorm:"column:remark;comment:备注"`
	TimeList []ReportTimeList `json:"timeList"`
	PostList []ReportPostList `json:"postList"`
}

func (RiskReportConfig) TableName() string {
	return "risk_report_config"
}
