package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go_blog/model/request"
	service "go_blog/service/system"
	"go_blog/utils"
	"net/http"
	"strconv"
)

type RoleApi struct {
}

func (r RoleApi) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	condition := make(map[string]interface{})
	condition["roleName"] = c.Query("RoleName")
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	logic := service.RoleService{}
	list, err := logic.GetRoleList(page, limit, condition)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "1",
		"data":    list,
	})
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
	logic := service.RoleService{}
	lErr := logic.AddRole(role)
	if lErr != nil {
		utils.ErrorRes(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    1,
	})
}
