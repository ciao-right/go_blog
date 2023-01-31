package request

import "go_blog/model"

type Role struct {
	model.BaseModel
	RoleName string `json:"roleName" validate:"required" gorm:"column:roleName"`
	RoleMark string `json:"roleMark" validate:"required" gorm:"column:roleMark"`
	State    int    `json:"state"`
	Remark   string `json:"remark"`
	//Permission []int  `json:"permission"`
}

func (r Role) TableName() string {
	return "sys_role"
}
