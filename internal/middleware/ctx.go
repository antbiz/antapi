package middleware

import (
	"github.com/BeanWei/apikit/ctxsrv"
	"github.com/antbiz/antapi/internal/common/types"
	"github.com/gogf/gf/net/ghttp"
)

// Ctx 自定义上下文变量
func Ctx(r *ghttp.Request) {
	if sessionUserID := r.Session.GetString("_id"); sessionUserID != "" {
		ctxsrv.Init(r, &ctxsrv.Ctx{
			Value: &types.Context{
				Session: r.Session,
				User: &types.ContextUser{
					ID:        sessionUserID,
					Username:  r.Session.GetString("username"),
					Phone:     r.Session.GetString("phone"),
					Email:     r.Session.GetString("email"),
					Avatar:    r.Session.GetString("avatar"),
					Language:  r.Session.GetString("language"),
					IsSysUser: r.Session.GetBool("isSysUser"),
					Sid:       r.Session.GetString("sid"),
				},
			},
		})
	}

	r.Middleware.Next()
}
