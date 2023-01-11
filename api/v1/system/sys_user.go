package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go_blog/model"
	"net/http"
)

type BaseApi struct {
}

func (b *BaseApi) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    200,
			"message": "未知错误",
			"data":    0,
		})
		return
	}
	v := validator.New()
	err := v.Struct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    200,
			"data":    0,
			"message": "参数错误",
		})
		return
	}

}

func (b *BaseApi) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}
