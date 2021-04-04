package req

import (
	"antapi/pkg/rqp"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
)

// SessionUser .
type SessionUser struct {
	ID        string   `json:"id"`
	Username  string   `json:"username"`
	Phone     string   `json:"phone"`
	Email     string   `json:"email"`
	Blocked   bool     `json:"blocked"`
	IsSysuser bool     `json:"is_sysuser"`
	Roles     []string `json:"roles"`
}

// GetSessionUserInfo 获取当前会话用户信息
func GetSessionUserInfo(r *ghttp.Request) *SessionUser {
	sess := r.Session
	return &SessionUser{
		ID:        sess.GetString("id"),
		Username:  sess.GetString("username"),
		Phone:     sess.GetString("phone"),
		Email:     sess.GetString("email"),
		Blocked:   sess.GetBool("blocked"),
		IsSysuser: sess.GetBool("is_sysuser"),
		Roles:     sess.GetStrings("roles"),
	}
}

// GetFilter 获取过滤器
// TODO: 按照指定的CollectionName以及请求中的参数和用户角色，
// TODO: 将解析后的sql conditions和数据权限设定中的规则合并喂给每个接口对应的dao层
func GetFilter(r *ghttp.Request) (p *rqp.Parse, err error) {
	p, err = rqp.New(r.GetUrl(), &rqp.Config{
		SkipWrongQuery:        true,
		TransformQueryKeyFunc: gstr.SnakeCase,
	})
	if err != nil {
		return
	}

	return
}
