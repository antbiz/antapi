package api

import (
	"antapi/common/errcode"
	"antapi/app/model"
	"antapi/common/resp"

	"github.com/gogf/gf/net/ghttp"
)

var SignIn = new(signInApi)

type signInApi struct{}

// SignInByUser 用户账号（用户名/手机号/邮箱）登录
func (signInApi) SignInByUser(r *ghttp.Request) {
	var data *model.UserSignInReq
	if err := r.Parse(&data); err != nil {
		resp.Error(r).SetError(err).SetCode(errcode.ParameterBindError).Json()
	}
}
