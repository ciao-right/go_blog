package system

import (
	"github.com/gin-gonic/gin"
	v1 "go_blog/api"
)

type DeptRouter struct{}

func (g GoodsRouter) InitDeptRouter(r *gin.RouterGroup) gin.IRouter {
	deptRouter := r.Group("/dept")
	deptApi := new(v1.ApiGroup).SystemApiGroup.DepartmentApi
	deptRouter.GET("/getDept", deptApi.GetDept)
	deptRouter.POST("/addDept", deptApi.AddDept)
	deptRouter.GET("/delete", deptApi.DelDept)
	deptRouter.POST("/updateDept", deptApi.UpdateDept)
	deptRouter.GET("/changeDeptState", deptApi.ChangeDeptState)
	return deptRouter
}
