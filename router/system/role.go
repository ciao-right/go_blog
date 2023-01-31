package system

import (
	"github.com/gin-gonic/gin"
	v1 "go_blog/api"
)

type RoleRouter struct{}

func (r *RoleRouter) InitRoleRouter(router *gin.RouterGroup) gin.IRouter {
	roleRouter := router.Group("/role")
	roleApi := new(v1.ApiGroup).SystemApiGroup.RoleApi
	roleRouter.GET("/getList", roleApi.GetList)
	roleRouter.POST("/addRole", roleApi.AddRole)
	return roleRouter
}
