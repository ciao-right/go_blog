package model

import (
	"fmt"
	"go_blog/utils"
	"gorm.io/gorm"
)

type User struct {
	baseModel
	Name     string `json:"name" validate:"required"`
	Sex      uint8  `json:"sex" `
	Address  string `json:"address" `
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"required"`
	Birth    string `json:"birth"`
	Avatar   string `json:"avatar"`
	State    int    `json:"-"`
	IsLogin  uint   `json:"isLogin"`
}

func (u *User) TableName() string {
	return "blog_users"
}

func (u *User) AfterCreate(DB *gorm.DB) (err error) {
	fmt.Println("before", u)
	DB.Model(u).Where("id = ?", u.ID).Update("created_on", utils.FormatTime(utils.GetNow(), "2006-01-02 15:04:05"))
	return err
}

//func (u *User) AfterUpdate(DB *gorm.DB) (err error) {
//	DB.Model(u).Where("id = ?", u.ID).Update("modified_on", utils.FormatTime(utils.GetNow(), "2006-01-02 15:04:05"))
//
//	return err
//}
