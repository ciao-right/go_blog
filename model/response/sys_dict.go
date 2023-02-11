package response

import (
	"go_blog/model"
	"go_blog/model/request"
)

type Dict struct {
	request.Dict
	model.OverrideTimeModel
}
