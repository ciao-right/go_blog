package response

import "go_blog/model/request"

// LoginRes 登陆返回结构
type LoginRes struct {
	request.User
	CreatedOn string `json:"created_on"`
	Token     string `json:"token"`
}
