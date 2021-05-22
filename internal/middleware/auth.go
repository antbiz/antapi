package middleware

import (
	"github.com/BeanWei/apikit/ctxsrv"
	"github.com/BeanWei/apikit/errors"
	"github.com/BeanWei/apikit/resp"
	"github.com/antbiz/antapi/internal/common/errmsg"
	"github.com/antbiz/antapi/internal/common/types"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Auth 鉴权中间件，依赖上下文对象中间件
func Auth(r *ghttp.Request) {
	var localCtx *types.Context

	err := ctxsrv.GetVar(r.Context()).Struct(&localCtx)
	if err != nil {
		g.Log().Async().Error(gerror.Wrap(err, "Auth Middleare: err conv r.Context to types.Context"))
	}
	if err != nil || localCtx == nil || localCtx.User == nil || localCtx.User.ID == "" {
		resp.Error(r, errors.Unauthorized(errmsg.Unauthorized, g.I18n().T(errmsg.Unauthorized)))
	}

	r.Middleware.Next()
}
