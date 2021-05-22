package api

import (
	"github.com/BeanWei/apikit/errors"
	"github.com/BeanWei/apikit/resp"
	"github.com/antbiz/antapi/internal/app/dao"
	"github.com/antbiz/antapi/internal/app/dto"
	"github.com/antbiz/antapi/internal/app/service"
	"github.com/antbiz/antapi/internal/common/errmsg"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

	jsonDoc, err := service.User.GetUserByLogin(r.Context(), req.Login)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			resp.Error(r, errors.NotFound(errmsg.AccountNotFound, g.I18n().T(errmsg.AccountNotFound)))
		}
		resp.Error(r, errors.DatabaseError(g.I18n().T(errmsg.ErrDBGet)).WithOrigErr(err))
	}

	username := jsonDoc.GetString("username")
	userpwd := jsonDoc.GetString("password")
	if userpwd != service.User.EncryptPwd(username, req.Pwd) {
		resp.Error(r, errors.Unauthorized(errmsg.IncorrectPassword, g.I18n().T(errmsg.IncorrectPassword)))
	}

	jsonDoc.Set("sid", r.Session.Id())
	r.Session.SetMap(jsonDoc.Map())
	resp.OK(r, jsonDoc)
}

// LogOut 退出登录
func (userApi) LogOut(r *ghttp.Request) {
	if err := r.Session.Remove(r.GetSessionId()); err != nil {
		resp.Error(r, errors.UnknownError(g.I18n().T(errmsg.ErrLogout)).WithOrigErr(err))
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
		if err == mongo.ErrNoDocuments {
			resp.Error(r, errors.NotFound(errmsg.AccountNotFound, g.I18n().T(errmsg.AccountNotFound)))
		}
		resp.Error(r, errors.DatabaseError(g.I18n().T(errmsg.ErrDBGet)).WithOrigErr(err))
	}
	resp.OK(r, jsonDoc)
}
