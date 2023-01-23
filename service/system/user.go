package service

import (
	"errors"
	"fmt"
	"go_blog/common/global"
	"go_blog/model/request"
	"go_blog/utils"
)

type UserService struct {
}

func (u UserService) Register(user request.User) (userInter request.User, err error) {
	var _user request.User
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

func (u UserService) Login(user utils.Login) (findUser request.User, err error) {
	res, findUser := isExist(user.Account)
	if res {
		//存在
		if utils.BcryptCheck(user.Password, findUser.Password) {
			global.GLOBAL_DB.Model(&request.User{}).Where("account = ?", user.Account).Update("is_login", 1)
			return findUser, err
		} else {
			return findUser, errors.New("密码错误")
		}
	}
	return findUser, errors.New("请先注册")
}

func isExist(account string) (result bool, user request.User) {
	global.GLOBAL_DB.Where("account = ?", account).First(&user)
	if user.ID == 0 {
		return false, user
	}
	return true, user

}

type userList struct {
	request.User
	CreatedOn  string `json:"created_on"`
	ModifiedOn string `json:"modified_on"`
	Password   string `json:"-"`
}

func (u UserService) GetUserList(userCond request.UserCond) (userList []userList, err error) {

	global.GLOBAL_DB.Limit(userCond.Limit).Offset(utils.GetPage(userCond.Page, userCond.Limit)).Where(&request.UserCond{Name: userCond.Name, Phone: userCond.Phone}).Find(&userList)
	//for _, list := range userList {
	//list.CreatedOn = utils.FormatTime(list.CreatedOn, utils.DateTime)
	//}
	return userList, err
}

func (u UserService) GetUserTotal(maps interface{}) (count int64) {
	global.GLOBAL_DB.Model(&request.User{}).Where(maps).Count(&count)
	return
}
func (u UserService) DelUser(id string) bool {
	err := global.GLOBAL_DB.Where("id = ?", id).Delete(&request.User{})
	return err != nil
}

func (u UserService) UpdateUser(bindUser request.User) error {
	if bindUser.ID != 0 {
		global.GLOBAL_DB.Model(&bindUser).Updates(bindUser)
		return nil
	} else {
		return errors.New("无id")
	}
}
