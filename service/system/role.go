package service

import (
	"go_blog/common/global"
	"go_blog/model/request"
	"go_blog/model/response"
	"go_blog/utils"
)

type RoleService struct {
}

func (r RoleService) AddRole(newRole request.Role) error {
	result := global.GLOBAL_DB.Create(&newRole)
	return result.Error
}

func (r RoleService) GetRoleList(page, limit int, condition request.Role) (roleList []response.RoleRes, err error) {
	result := global.GLOBAL_DB.Where(condition).Offset(utils.GetPage(page, limit)).Limit(limit).Find(&roleList)

	return roleList, result.Error
}

func (r RoleService) UpdateRole(role request.Role) error {
	result := global.GLOBAL_DB.Updates(&role)
	return result.Error
}

func (r RoleService) DeleteRole(id int) error {
	var role request.Role
	result := global.GLOBAL_DB.Where("id=?", id).Delete(&role)
	return result.Error
}
