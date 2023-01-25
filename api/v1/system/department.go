package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go_blog/model/request"
	"go_blog/model/response"
	service "go_blog/service/system"
	"go_blog/utils"
	"net/http"
)

type DepartmentApi struct{}

type deptList []response.Department

func (l deptList) ToINode() (nodes []utils.INode) {
	for _, v := range l {
		nodes = append(nodes, v)
	}
	return
}
func (d DepartmentApi) GetDept(c *gin.Context) {
	logic := service.DeptService{}
	condition := make(map[string]string)
	//todo
	// 还需要加 一个state 用来其他地方筛选出可以使用的dept
	// 还需要考虑修改或者删除上级部门时 会发生的情况
	condition["name"] = c.Query("name")
	list := logic.GetDept(condition)
	reps := utils.GenerateTree(deptList.ToINode(list), nil)
	c.JSON(http.StatusOK, utils.ResStruct(200, "success", reps))
}
func (d DepartmentApi) AddDept(c *gin.Context) {
	var newDept request.Department
	err := c.ShouldBindJSON(&newDept)
	if err != nil {
		c.JSON(http.StatusOK, utils.ResStruct(201, utils.ErrToString(err), nil))
		return
	}
	v := validator.New()
	VErr := v.Struct(newDept)
	if VErr != nil {
		c.JSON(http.StatusOK, utils.ResStruct(201, utils.ErrToString(VErr), nil))
		return
	}
	logic := service.DeptService{}
	lErr := logic.AddDept(newDept)
	if lErr != nil {
		c.JSON(http.StatusOK, utils.ResStruct(201, utils.ErrToString(lErr), nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResStruct(200, "success", 1))

}

func (d DepartmentApi) UpdateDept(c *gin.Context) {
	var updateDept response.Department
	err := c.ShouldBindJSON(&updateDept)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}
	logic := service.DeptService{}
	lErr := logic.UpdateDept(updateDept)
	if lErr != nil {
		utils.ErrorRes(lErr, c)
		return
	}
	c.JSON(http.StatusOK, utils.ResStruct(200, "success", 1))
}

func (d DepartmentApi) DelDept(c *gin.Context) {
	id := c.Query("id")
	logic := service.DeptService{}
	err := logic.DelDept(id)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}
	c.JSON(http.StatusOK, utils.ResStruct(200, "success", 1))

}
func (d DepartmentApi) ChangeDeptState(c *gin.Context) {
	id := c.Query("id")
	state := c.Query("state")
	logic := service.DeptService{}
	err := logic.ChangeState(id, state)
	if err != nil {
		utils.ErrorRes(err, c)
		return
	}
	c.JSON(http.StatusOK, utils.ResStruct(200, "success", 1))

}
