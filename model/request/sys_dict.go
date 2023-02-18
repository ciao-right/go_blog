package request

import (
	"go_blog/model"
)

type Dict struct {
	model.BaseModel
	ParentId  int    `json:"parentId" gorm:"column:parentId;comment:父字典id" `
	DictName  string `json:"dictName" gorm:"column:dictName;comment:字典名称" Validate:"required"`
	DictValue string `json:"dictValue" gorm:"column:dictValue;comment:字典值"  Validate:"required"`
	Status    int    `json:"status" gorm:"column:status;comment:状态（0正常 1停用）"`
	Remark    string `json:"remark" gorm:"column:remark;comment:备注"`
}

func (d *Dict) TableName() string {
	return "sys_dict"
}
