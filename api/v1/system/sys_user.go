package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go_blog/common/global"
	"go_blog/model"
	service "go_blog/service/system"
	"go_blog/utils"
	"net/http"
)

type BaseApi struct {
}

func (b *BaseApi) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    201,
			"message": err,
			"data":    0,
		})
		return
	}
	v := validator.New()
	err := v.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    200,
			"data":    0,
			"message": "参数错误",
		})
		return
	}
	_, err = service.UserService{}.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    201,
			"data":    0,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    1,
		"message": "注册成功",
	})

}

func (b *BaseApi) Login(c *gin.Context) {
	var loginObj utils.Login
	if err := c.ShouldBindJSON(&loginObj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    201,
			"message": "参数错误",
			"data":    0,
		})
		return
	}

	v := validator.New()
	err := v.Struct(loginObj)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    201,
			"message": "请输入正确的参数",
			"data":    0,
		})
		return
	}
	findUser, err := service.UserService{}.Login(loginObj)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    201,
			"message": err,
			"data":    0,
		})
		return
	}
	fmt.Println(findUser)
	if token, tokenErr := utils.GenerateToken(findUser.Account, findUser.Password); tokenErr != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"code":    201,
			"message": tokenErr,
			"data":    nil,
			"token":   "",
		})
	} else {
		// todo 应该在数据库中更新
		findUser.IsLogin = 1
		findUser.Password = ""
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data": struct {
				model.User
				CreatedOn string `json:"created_on"`
				Token     string `json:"token"`
			}{
				findUser,
				utils.FormatTime(*findUser.CreatedOn, utils.DateTime),
				token,
			},
		})
	}

}

type UserApi struct{}

// GetUserList 获取用户列表
func (u *UserApi) GetUserList(c *gin.Context) {
	var userCond model.UserCond
	err := c.ShouldBindJSON(&userCond)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    201,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}
	userList, userListErr := service.UserService{}.GetUserList(userCond)
	c.JSON(http.StatusOK, gin.H{
		"code":    global.SUCCESS,
		"message": "success",
		"data":    userList,
	})
	fmt.Println(userList, userListErr)
}
