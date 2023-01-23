package request

import "go_blog/model"

type User struct {
	model.BaseModel
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
	model.PageListModel
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (u *User) TableName() string {
	return "blog_users"
}
