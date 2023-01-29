package service

import (
	"go_blog/common/global"
	"go_blog/model/request"
	"go_blog/model/response"
)

type RoleService struct {
}

func (r RoleService) AddRole(newRole request.Role) error {
	result := global.GLOBAL_DB.Create(&newRole)
	return result.Error
}

func (r RoleService) GetRoleList(page, limit int, condition map[string]interface{}) (roleList []response.RoleRes, err error) {
	result := global.GLOBAL_DB.Where(condition).Offset(page).Limit(limit).Find(&roleList)
	return roleList, result.Error
}
