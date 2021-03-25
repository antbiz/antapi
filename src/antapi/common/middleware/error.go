package middleware

import (
	"antapi/common/errcode"
	"antapi/common/resp"
	"fmt"
	"net/http"

	"github.com/gogf/gf/net/ghttp"
)

// ErrorHandler 顶层的handler
func ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	switch r.Response.Status {
	case http.StatusNotFound:
		r.Response.ClearBuffer()
		resp.Error(r).SetCode(errcode.APINotFound).SetMsg(fmt.Sprintf(errcode.APINotFoundMsg, r.URL.String())).Json()
	case http.StatusInternalServerError:
		r.Response.ClearBuffer()
		resp.Error(r).SetCode(errcode.ServerError).SetMsg(errcode.ServerErrorMsg).Json()
	}
}
