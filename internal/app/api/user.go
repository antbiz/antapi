package api

import (
	"github.com/BeanWei/apikit/ctxsrv"
	"github.com/BeanWei/apikit/errors"
	"github.com/BeanWei/apikit/resp"
	"github.com/antbiz/antapi/internal/app/dao"
	"github.com/antbiz/antapi/internal/app/dto"
	"github.com/antbiz/antapi/internal/app/service"
	"github.com/antbiz/antapi/internal/common/errmsg"
	"github.com/antbiz/antapi/internal/common/types"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User 用户接口
var User = &userApi{}

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
	if userpwd != service.User.EncryptPwd(username, req.Password) {
		resp.Error(r, errors.Unauthorized(errmsg.IncorrectPassword, g.I18n().T(errmsg.IncorrectPassword)))
	}

	jsonDoc.Remove("password")
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

// GetInfo 获取个人信息
func (userApi) GetInfo(r *ghttp.Request) {
	userID, err := primitive.ObjectIDFromHex(r.Session.GetString("_id"))
	if err != nil {
		resp.Error(r, errors.NotFound(errmsg.AccountNotFound, g.I18n().T(errmsg.AccountNotFound)))
	}
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

// UpdateInfo 修改个人信息
func (userApi) UpdateInfo(r *ghttp.Request) {
	var (
		req      *dto.UserUpdateInfoReq
		localCtx types.Context
		ctxUser  types.ContextUser
	)
	ctx := r.Context()

	if err := ctxsrv.GetVar(ctx).Struct(&localCtx); err != nil {
		g.Log().Async().Error(gerror.Wrap(err, "error get user info from context"))
	} else if localCtx.User != nil {
		ctxUser = *localCtx.User
	}

	if err := r.Parse(&req); err != nil {
		resp.Error(r, errors.InvalidArgument(err.Error()))
	}
	userID, err := primitive.ObjectIDFromHex(r.Session.GetString("_id"))
	if err != nil {
		resp.Error(r, errors.NotFound(errmsg.AccountNotFound, g.I18n().T(errmsg.AccountNotFound)))
	}
	err = dao.Update(
		r.Context(),
		service.User.CollectionName(),
		g.Map{
			"language": req.Language,
			"avatar":   req.Avatar,
		},
		&dao.UpdateOptions{
			Filter:                bson.M{"_id": userID},
			IgnoreFieldValueCheck: true,
			CtxUser:               ctxUser,
		},
	)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			resp.Error(r, errors.NotFound(errmsg.AccountNotFound, g.I18n().T(errmsg.AccountNotFound)))
		}
		resp.Error(r, errors.DatabaseError(g.I18n().T(errmsg.ErrDBUpdate)).WithOrigErr(err))
	}
	resp.OK(r)
}

// UpdatePassword 修改个人密码
func (userApi) UpdatePassword(r *ghttp.Request) {
	var (
		req      *dto.UserUpdatePasswordReq
		localCtx types.Context
		ctxUser  types.ContextUser
	)
	ctx := r.Context()

	if err := ctxsrv.GetVar(ctx).Struct(&localCtx); err != nil {
		g.Log().Async().Error(gerror.Wrap(err, "error get user info from context"))
	} else if localCtx.User != nil {
		ctxUser = *localCtx.User
	}

	if err := r.Parse(&req); err != nil {
		resp.Error(r, errors.InvalidArgument(err.Error()))
	}
	userID, err := primitive.ObjectIDFromHex(r.Session.GetString("_id"))
	if err != nil {
		resp.Error(r, errors.NotFound(errmsg.AccountNotFound, g.I18n().T(errmsg.AccountNotFound)))
	}

	jsonDoc, err := dao.Get(ctx, service.User.CollectionName(), &dao.GetOptions{
		Filter:              bson.M{"_id": userID},
		IncludePrivateField: true,
	})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			resp.Error(r, errors.NotFound(errmsg.AccountNotFound, g.I18n().T(errmsg.AccountNotFound)))
		}
		resp.Error(r, errors.DatabaseError(g.I18n().T(errmsg.ErrDBGet)).WithOrigErr(err))
	}

	username := jsonDoc.GetString("username")
	userpwd := jsonDoc.GetString("password")
	if userpwd != service.User.EncryptPwd(username, req.OldPassword) {
		resp.Error(r, errors.InvalidArgument(g.I18n().T(errmsg.IncorrectOldPassword)))
	}

	err = dao.Update(
		ctx,
		service.User.CollectionName(),
		g.Map{
			"password": service.User.EncryptPwd(username, req.Password),
		},
		&dao.UpdateOptions{
			IgnoreFieldValueCheck: true,
			IncludePrivateField:   true,
			CtxUser:               ctxUser,
		},
	)
	if err != nil {
		resp.Error(r, errors.DatabaseError(g.I18n().T(errmsg.ErrDBUpdate)).WithOrigErr(err))
	}
	resp.OK(r)
}
