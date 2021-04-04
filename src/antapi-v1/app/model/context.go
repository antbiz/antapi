package model

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	// ContextKey 上下文变量存储键名，前后端系统共享
	ContextKey = "ContextKey"
)

// Context 请求上下文结构
type Context struct {
	Session *ghttp.Session // 当前Session管理对象
	User    *ContextUser   // 上下文用户信息
	Data    g.Map          // 自定KV变量，业务模块根据需要设置，不固定
}

// ContextUser 请求上下文中的用户信息
type ContextUser struct {
	ID       uint   // 用户ID
	Passport string // 用户账号
	Nickname string // 用户名称
	Avatar   string // 用户头像
}
