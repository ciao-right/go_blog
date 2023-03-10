package system

import (
	"fmt"
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
	condition := request.Role{
		RoleName: c.Query("RoleName"),
	}
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
	var res = make(map[string]interface{})
	res["list"] = list
	res["total"], _ = logic.GetListTotal(condition)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": 1,
		"data":    res,
	})
}

func (r RoleApi) AddRole(c *gin.Context) {
	var role request.Role
	err := c.ShouldBindJSON(&role)
	fmt.Println(err)
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

func (r RoleApi) UpdateRole(c *gin.Context) {
	var role request.Role
	err := c.ShouldBindJSON(&role)
	if err != nil {
		fmt.Println(err)
		utils.ErrorRes(err, c)
		return
	}
	v := validator.New()
	vErr := v.Struct(role)
	if vErr != nil {
		fmt.Println(vErr)
		utils.ErrorRes(vErr, c)
		return

	}
	logic := service.RoleService{}
	lErr := logic.UpdateRole(role)
	if lErr != nil {
		fmt.Println(lErr)
		utils.ErrorRes(lErr, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    1,
	})

}

func (r RoleApi) DeleteRole(c *gin.Context) {
	logic := service.RoleService{}
	id, _ := strconv.Atoi(c.Query("id"))
	err := logic.DeleteRole(id)
	if err != nil {
		fmt.Println(err)
		utils.ErrorRes(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    1,
	})
}
