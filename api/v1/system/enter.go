package system

import (
	"go_blog/api/v1/system/ehs/riskBaseConfig"
)

type ApiGroup struct {
	BaseApi
	UserApi
	GoodsClassApi
	DepartmentApi
	RoleApi
	DictApi
	riskBaseConfig.EHSApi
}
