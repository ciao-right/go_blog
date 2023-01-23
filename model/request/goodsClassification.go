package request

import (
	"go_blog/model"
)

type GoodsClassification struct {
	model.BaseModel
	TypeName string                `json:"typeName" gorm:"column:typeName" validate:"required"`
	Sort     int                   `json:"sort" `
	ParentId uint                  `json:"parentId" gorm:"column:parentId"`
	Children []GoodsClassification `json:"children" gorm:"-"`
}

func (g GoodsClassification) TableName() string {
	return "blog_goods_class"
}

func (g GoodsClassification) GetTitle() string {
	return g.TypeName
}

func (g GoodsClassification) GetId() uint {
	return g.ID
}

func (g GoodsClassification) GetParentId() uint {
	return g.ParentId
}

func (g GoodsClassification) GetData() any {
	return g
}

func (g GoodsClassification) IsRoot() bool {
	return g.ParentId == 0 || g.ParentId == g.ID
}
