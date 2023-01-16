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
	if res, _ := isExist(user.Account); !res {
		user.Password = utils.BcryptHash(user.Password)
		user.State = 1
		user.IsLogin = 0
		//user.ID = uuid.NewV4().String()
		fmt.Println(user.CreatedOn)
		err = global.GLOBAL_DB.Create(&user).Error
		return user, err
	} else {
		return _user, errors.New("该账号已经注册")
	}
}

func (u UserService) Login(user utils.Login) (findUser model.User, err error) {
	res, findUser := isExist(user.Account)
	if res {
		//存在
		if utils.BcryptCheck(user.Password, findUser.Password) {
			global.GLOBAL_DB.Where("account = ?", user.Account).Update("isLogin", 1)
			return findUser, err
		} else {
			return findUser, errors.New("密码错误")
		}
	}
	return findUser, errors.New("请先注册")
}

func isExist(account string) (result bool, user model.User) {
	global.GLOBAL_DB.Where("account = ?", account).First(&user)
	if user.ID == 0 {
		return false, user
	}
	return true, user

}
