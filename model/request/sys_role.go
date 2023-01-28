package request

import "go_blog/model"

type Role struct {
	model.BaseModel
	RoleName   string `json:"roleName" validate:"required"`
	RoleMark   string `json:"roleMark" validate:"required"`
	State      int    `json:"state"`
	Remark     string `json:"remark"`
	Permission []int  `json:"permission"`
}

func (r Role) TableName() string {
	return "sys_role"
}
