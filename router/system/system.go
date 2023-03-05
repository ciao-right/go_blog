package system

import "go_blog/router/system/EHSRouter"

type SysRouterGroup struct {
	BaseRouter
	UserRouter
	GoodsRouter
	RoleRouter
	DictRouter
	EHSRouter.EHSRouter
	EHSRouter.RiskReportConfig
}
