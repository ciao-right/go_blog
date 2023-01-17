package model

import (
	"go_blog/utils"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Name        string `json:"name" validate:"required"`
	Account     string `json:"account" validate:"required"`
	Password    string `json:"password,omitempty" validate:"required,min=6"`
	Sex         uint8  `json:"sex" `
	Address     string `json:"address" `
	Phone       string `json:"phone" validate:"required,min=13"`
	AuthorityId uint   `json:"authorityId" gorm:"default:888;commit:用户角色"`
	Birth       string `json:"birth" gorm:"default:'';commit:生日"`
	Avatar      string `json:"avatar" gorm:"default:'';commit:头像"`
	State       int    `json:"state"`
	IsLogin     int    `json:"isLogin"`
}

type UserCond struct {
	PageListModel
	Name  string
	Phone string
}

func (u *User) TableName() string {
	return "blog_users"
}

func (u *User) BeforeCreate(DB *gorm.DB) (err error) {
	var time = utils.GetNow()
	u.CreatedOn = &time
	return err
}

func (u *User) AfterUpdate(DB *gorm.DB) (err error) {
	var time = utils.GetNow()
	u.ModifiedOn = &time
	return err
}
