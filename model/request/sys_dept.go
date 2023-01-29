package request

import "go_blog/model"

type Department struct {
	model.BaseModel
	Name       string       `json:"name" validate:"required" `
	ParentDept uint         `json:"parentDept" gorm:"column:parentDept"`
	State      int          `json:"state"`
	Remark     string       `json:"remark"`
	Children   []Department `gorm:"-"`
}

func (d Department) TableName() string {
	return "sys_dept"
}
