package api

import (
	"antapi/common/errcode"
	"antapi/common/resp"

	"github.com/gogf/gf/net/ghttp"
)

var User = new(userApi)

type userApi struct{}

// SignOut 退出登录
func (userApi) SignOut(r *ghttp.Request) {
	if err := r.Session.Remove(r.GetSessionId()); err != nil {
		resp.Error(r).SetMsg(errcode.ServerErrorMsg).SetCode(errcode.ServerError).Json()
	}
	resp.Success(r).Json()
}
