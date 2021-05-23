package middleware

import (
	"net/http"

	"github.com/BeanWei/apikit/errors"
	"github.com/BeanWei/apikit/resp"
	"github.com/antbiz/antapi/internal/common/errmsg"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// ErrorHandler 顶层的handler
func ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()

	switch r.Response.Status {
	case http.StatusInternalServerError:
		r.Response.ClearBuffer()
		resp.Error(r, errors.InternalServer(errmsg.ServerError, g.I18n().T(errmsg.ServerError)))
	}
}
