package system

import (
	"github.com/gin-gonic/gin"
	v1 "go_blog/api"
)

type UserRouter struct{}

func (u UserRouter) InitUserRouter(r *gin.RouterGroup) gin.IRouter {
	userRouter := r.Group("/user")
	userApi := new(v1.ApiGroup).SystemApiGroup.UserApi
	userRouter.POST("/getList", userApi.GetUserList)
	userRouter.GET("/delete", userApi.DelUser)
	return userRouter
}
