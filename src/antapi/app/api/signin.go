package api

import (
	"antapi/app/logic"
	"antapi/app/model"
	"antapi/common/errcode"
	"antapi/common/resp"

	"github.com/gogf/gf/errors/gerror"
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
	res, err := logic.User.SignIn(data, r)
	if err != nil {
		resp.Error(r).SetError(gerror.Current(err)).SetCode(gerror.Code(err)).Json()
	}
	resp.Success(r).SetData(res).Json()
}
