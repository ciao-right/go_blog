package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go_blog/model/request"
	"go_blog/utils"
)

type RoleApi struct {
}

func (r RoleApi) GetList(c *gin.Context) {

}

func (r RoleApi) AddRole(c *gin.Context) {
	var role request.Role
	err := c.ShouldBindJSON(&role)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}
	v := validator.New()
	vErr := v.Struct(role)
	if vErr != nil {
		utils.ErrorRes(vErr, c)
		return
	}
	//todo
}
