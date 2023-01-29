package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_blog/common/global"
	"go_blog/model/request"
	service "go_blog/service/system"
	"net/http"
)

type UserApi struct{}

// GetUserList 获取用户列表
func (u *UserApi) GetUserList(c *gin.Context) {
	var userCond request.UserCond
	logic := service.UserService{}
	err := c.ShouldBindJSON(&userCond)
	fmt.Println(userCond.Page)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    201,
			"message": err,
			"data":    nil,
		})
		return
	}

	userList, userListErr := logic.GetUserList(userCond)
	if userListErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    global.ERROR,
			"message": userListErr,
			"data":    nil,
			"total":   0,
			"page":    userCond.Page,
			"limit":   userCond.Limit,
		})
	}
	maps := make(map[string]string)
	maps["name"] = userCond.Name
	maps["phone"] = userCond.Phone
	total := logic.GetUserTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code":    global.SUCCESS,
		"message": "success",
		"data":    userList,
		"total":   total,
		"page":    userCond.Page,
		"limit":   userCond.Limit,
	})
}

// DelUser 删除用户
func (u *UserApi) DelUser(c *gin.Context) {
	id := c.Query("id")
	fmt.Println(id)
	logic := service.UserService{}
	if logic.DelUser(id) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": 1,
			"data":    1,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "服务器错误",
			"data":    0,
		})
	}

}

func (u *UserApi) UpdateUser(c *gin.Context) {
	var bindUser request.User
	err := c.ShouldBindJSON(&bindUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    201,
			"message": err,
			"data":    0,
		})
		return
	}
	logic := service.UserService{}
	dbErr := logic.UpdateUser(bindUser)
	if dbErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    201,
			"message": dbErr,
			"data":    0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    1,
		})
	}

}
