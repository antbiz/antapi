package req

import (
	"antapi/pkg/rqp"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
)

// GetSessionUsername 获取当前会话用户名
func GetSessionUsername(r *ghttp.Request) string {
	return r.Session.GetString("username")
}

// GetSessionRoles 获取当前会话用户角色
func GetSessionUserRoles(r *ghttp.Request) []string {
	roles := r.Session.GetStrings("roles")
	if len(roles) == 0 {
		return []string{"Guest"}
	}
	return roles
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
