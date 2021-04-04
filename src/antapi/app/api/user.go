package api

import (
	"antapi/app/logic"
	"antapi/common/errcode"
	"antapi/common/resp"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
)

// User 用户接口
var User = userApi{}

type userApi struct{}

// SignOut 退出登录
func (userApi) SignOut(r *ghttp.Request) {
	if err := r.Session.Remove(r.GetSessionId()); err != nil {
		resp.Error(r).SetMsg(errcode.ServerErrorMsg).SetCode(errcode.ServerError).Json()
	}
	resp.Success(r).Json()
}

// MyProfile 当前回话用户的信息
func (userApi) MyProfile(r *ghttp.Request) {
	userID := r.Session.GetString("id")
	dataGJson, err := logic.User.GetProfileByID(userID)
	if err != nil {
		resp.Error(r).SetError(gerror.Current(err)).SetCode(gerror.Code(err)).Json()
	}
	resp.Success(r).SetData(dataGJson).Json()
}
