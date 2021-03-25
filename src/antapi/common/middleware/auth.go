package middleware

import (
	"antapi/common/errcode"
	"antapi/common/resp"

	"github.com/gogf/gf/net/ghttp"
)

// Auth 鉴权中间件
func Auth(r *ghttp.Request) {
	userID := r.Session.GetString("id")
	if userID == "" {
		resp.Error(r).SetCode(errcode.AuthorizationError).SetMsg(errcode.AuthorizationErrorMsg).Json()
	}
	r.Middleware.Next()
}
