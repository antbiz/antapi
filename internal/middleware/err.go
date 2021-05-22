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
	case http.StatusNotFound:
		r.Response.ClearBuffer()
		resp.Error(r, errors.NotFound("api_not_found", "api_not_found"))
	case http.StatusInternalServerError:
		r.Response.ClearBuffer()
		resp.Error(r, errors.UnknownError(g.I18n().T(errmsg.UnknownError)))
	}
}
