package service

import (
	"errors"
	"fmt"
	"go_blog/common/global"
	"go_blog/model"
	"go_blog/model/request"
	"go_blog/model/response"
	"go_blog/utils"
)

type DeptService struct {
}

func CommonIsExist[T any](name string, field string) T {
	var hasSame T
	condition := fmt.Sprintf("%s = ?", field)
	global.GLOBAL_DB.Where(condition, name).First(&hasSame)
	return hasSame
}

func (d DeptService) AddDept(dept request.Department) error {
	yn := CommonIsExist[request.Department](dept.Name, "name")
	if yn.ID > 0 && yn.ParentDept == 0 {
		return errors.New("无法添加相同的顶层部门名称")
	} else {
		result := global.GLOBAL_DB.Create(&dept)
		return result.Error
	}

}

func (d DeptService) UpdateDept(dept response.Department) error {
	result := global.GLOBAL_DB.Where("id = ?", dept.ID).Select("name", "remark", "state").Updates(&dept)
	return result.Error
}

func (d DeptService) DelDept(id string) error {
	if id == "" {
		return errors.New("未传id")
	} else {
		result := global.GLOBAL_DB.Where("id = ?", id).Where("parentDept = ?", id).Delete(&request.Department{})
		return result.Error
	}
}

func (d DeptService) GetDept(condition map[string]interface{}) []response.Department {
	var list []request.Department
	if condition["name"] == "" {
		if value, ok := condition["state"]; ok {
			global.GLOBAL_DB.Where("state = ?", value).Find(&list)
		} else {
			global.GLOBAL_DB.Find(&list)
		}
	} else {
		global.GLOBAL_DB.Where("name like ?", "%"+condition["name"].(string)+"%").Where("state = ?", condition["state"]).Find(&list)
	}
	resList := make([]response.Department, len(list))
	for i, v := range list {
		resList[i] = response.Department{
			Department: v,
			OverrideTimeModel: model.OverrideTimeModel{
				CreatedOn:  utils.FormatTime(v.CreatedOn, utils.DateTime),
				ModifiedOn: utils.FormatTime(v.ModifiedOn, utils.DateTime),
			},
		}
	}
	return resList
}

func (d DeptService) ChangeState(id string, state string) error {
	var dept request.Department
	result := global.GLOBAL_DB.Model(&dept).Where("id = ?", id).Update("state", state)
	return result.Error
}
