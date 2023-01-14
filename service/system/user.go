package service

import (
	"errors"
	"fmt"
	"go_blog/common/global"
	"go_blog/model"
	"go_blog/utils"
)

type UserService struct {
}

func (u UserService) Register(user model.User) (userInter model.User, err error) {
	var _user model.User
	if !isExist(user.Account) {
		user.Password = utils.BcryptHash(user.Password)
		user.State = 1
		user.IsLogin = 0
		fmt.Println(user)
		err = global.GLOBAL_DB.Create(&user).Error
		return user, err
	} else {
		return _user, errors.New("该账号已经注册")
	}
}

func isExist(account string) bool {
	var user model.User
	global.GLOBAL_DB.First(&user, "account = ?", account)
	//global.GLOBAL_DB.Where("account = ?", account).First(&user)
	fmt.Println(user.ID)
	if user.ID == 0 {
		return false
	}
	return true

}
