package api

import (
	"antapi/app/logic"
	"antapi/app/model"
	"antapi/common/errcode"
	"antapi/common/resp"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
)

var SignUp = new(signUpApi)

type signUpApi struct{}

// SignUpWithEmail 用户邮箱注册
func (signUpApi) SignUpWithEmail(r *ghttp.Request) {
	var data *model.UserSignUpWithEmailReq
	if err := r.Parse(&data); err != nil {
		resp.Error(r).SetError(err).SetCode(errcode.ParameterBindError).Json()
	}
	if err := logic.User.SignUpWithEmail(data); err != nil {
		resp.Error(r).SetError(gerror.Current(err)).SetCode(gerror.Code(err)).Json()
	}
	resp.Success(r).Json()
}

// SignUpWithPhone 用户手机号注册
func (signUpApi) SignUpWithPhone(r *ghttp.Request) {
	var data *model.UserSignUpWithPhoneReq
	if err := r.Parse(&data); err != nil {
		resp.Error(r).SetError(err).SetCode(errcode.ParameterBindError).Json()
	}
	if err := logic.User.SignUpWithPhone(data); err != nil {
		resp.Error(r).SetError(gerror.Current(err)).SetCode(gerror.Code(err)).Json()
	}
	resp.Success(r).Json()
}
