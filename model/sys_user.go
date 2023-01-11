package model

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"gorm.Model"`
	Name       string `json:"name" validate:"required"`
	Sex        uint8  `json:"sex" `
	Address    string `json:"address" `
	Account    string `json:"account" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Birth      string `json:"birth"`
	avatar     string `json:"avatar"`
}

func (User) TableName() string {
	return "blog_users"
}
