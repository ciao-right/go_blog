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

type BaseApi struct {
}

func (b *BaseApi) Register(c *gin.Context) {
	var user request.User
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
	if token, tokenErr := utils.GenerateToken(findUser.Account, findUser.Password); tokenErr != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"code":    202,
			"message": tokenErr,
			"data":    nil,
			"token":   "",
		})
	} else {
		findUser.Password = ""
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data": response.LoginRes{
				User:  findUser,
				Token: token,
			},
		})
	}

}
