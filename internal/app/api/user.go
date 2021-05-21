package api

import (
	"github.com/BeanWei/apikit/errors"
	"github.com/BeanWei/apikit/resp"
	"github.com/antbiz/antapi/internal/app/dao"
	"github.com/antbiz/antapi/internal/app/dto"
	"github.com/antbiz/antapi/internal/app/service"
	"github.com/gogf/gf/net/ghttp"
	"go.mongodb.org/mongo-driver/bson"
)

// User 用户接口
var User = userApi{}

type userApi struct{}

// LoginByAccount 用户账号（用户名/手机号/邮箱）登录
func (userApi) LoginByAccount(r *ghttp.Request) {
	var req *dto.UserLoginReq
	if err := r.Parse(&req); err != nil {
		resp.Error(r, errors.InvalidArgument(err.Error()))
	}

	jsonDoc, err := service.User.GetUserByLogin(r.Context(), req.Login, req.Pwd)
	if err != nil {
		resp.Error(r, errors.DatabaseError(""))
	}
	if jsonDoc == nil {
		resp.Error(r, errors.Unauthorized("incorrect username or password", "incorrect username or password"))
	}

	jsonDoc.Set("sid", r.Session.Id())
	r.Session.SetMap(jsonDoc.Map())
	resp.OK(r, jsonDoc)
}

// LogOut 退出登录
func (userApi) LogOut(r *ghttp.Request) {
	if err := r.Session.Remove(r.GetSessionId()); err != nil {
		resp.Error(r, errors.InternalServer("", ""))
	}
	resp.OK(r)
}

// Info 当前回话用户的信息
func (userApi) Info(r *ghttp.Request) {
	userID := r.Session.GetString("_id")
	jsonDoc, err := dao.Get(r.Context(), service.User.CollectionName(), &dao.GetOptions{
		Filter: bson.M{"_id": userID},
	})
	if err != nil {
		resp.Error(r, errors.DatabaseError(""))
	}
	resp.OK(r, jsonDoc)
}
