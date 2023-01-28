package response

import (
	"go_blog/model"
	"go_blog/model/request"
)

type RoleRes struct {
	request.Role
	model.OverrideTimeModel
}
