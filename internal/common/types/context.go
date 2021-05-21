package types

import "github.com/gogf/gf/net/ghttp"

// Context 请求上下文结构
type Context struct {
	Session *ghttp.Session
	User    *ContextUser
}

// ContextUser 请求上下文中的用户信息
type ContextUser struct {
	ID        string `json:"_id"`       // 用户 id
	Username  string `json:"username"`  // 用户名
	Phone     string `json:"phone"`     // 手机号
	Email     string `json:"email"`     // 邮箱
	Avatar    string `json:"avatar"`    // 头像
	Language  string `json:"language"`  // 语言
	IsSysUser bool   `json:"isSysUser"` // 是否为系统用户
	Sid       string `json:"sid"`       // session id
}
