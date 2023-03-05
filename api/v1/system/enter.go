package system

import (
	"go_blog/api/v1/system/ehs/riskBaseConfig"
	EHSApi "go_blog/api/v1/system/ehs/riskReportConfig"
)

type ApiGroup struct {
	BaseApi
	UserApi
	GoodsClassApi
	DepartmentApi
	RoleApi
	DictApi
	riskBaseConfig.EHSApi
	EHSApi.RiskReportConfigApi
}
