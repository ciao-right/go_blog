package service

import (
	"go_blog/common/global"
	"go_blog/model"
)

type UserService struct {
}

func (u UserService) Register(user model.User) (userInter model.User, err error) {
	var _user model.User
	global.GLOBAL_DB.Where("account = ?", user.Account).First(&_user)
}
