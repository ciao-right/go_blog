package response

import (
	"go_blog/model"
	"go_blog/model/request"
)

type Department struct {
	request.Department
	model.OverrideTimeModel
}

func (d Department) GetTitle() string {
	return d.Name
}

func (d Department) GetId() uint {
	return d.ID
}

func (d Department) GetParentId() uint {
	return d.ParentDept
}

func (d Department) GetData() any {
	return d
}

func (d Department) IsRoot() bool {
	return d.ParentDept == 0 || d.ParentDept == d.ID
}
