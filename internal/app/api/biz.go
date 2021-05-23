package api

import (
	"github.com/BeanWei/apikit/ctxsrv"
	"github.com/BeanWei/apikit/errors"
	"github.com/BeanWei/apikit/resp"
	"github.com/antbiz/antapi/internal/app/dao"
	"github.com/antbiz/antapi/internal/app/service"
	"github.com/antbiz/antapi/internal/common/errmsg"
	"github.com/antbiz/antapi/internal/common/types"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Biz 公共接口管理
var Biz = &bizApi{}

type bizApi struct{}

// List 列表查询
func (bizApi) List(r *ghttp.Request) {
	var (
		localCtx types.Context
		ctxUser  types.ContextUser
	)
	ctx := r.Context()
	if err := ctxsrv.GetVar(ctx).Struct(&localCtx); err != nil {
		g.Log().Async().Error(gerror.Wrap(err, "error get user info from context"))
	} else if localCtx.User != nil {
		ctxUser = *localCtx.User
	}
	collectionName := r.GetString("collection")
	opt := &dao.ListOptions{
		Limit:   r.GetInt64("pageSize"),
		Sort:    gstr.SplitAndTrimSpace(r.GetString("sort"), ","),
		CtxUser: ctxUser,
		Filter:  gconv.Map(r.GetString("filter")),
	}
	opt.Offset = opt.Limit * (r.GetInt64("pageNum", 1) - 1)

	if !ctxUser.IsSysUser {
		perm := service.Permission.GetReadPermission(collectionName)
		if !perm.CanDoAll() {
			if perm.CanDoOnlyLogin() && ctxUser.ID == "" {
				resp.Error(r, errors.Unauthorized(errmsg.Unauthorized, g.I18n().T(errmsg.Unauthorized)))
			} else if perm.CanDoOnlyOwner() {
				opt.Filter = bson.M{"createdBy": ctxUser.ID}
			} else if perm.CanDoOnlySysUser() {
				resp.Error(r, errors.Forbidden(errmsg.PermissionDenied, g.I18n().T(errmsg.PermissionDenied)))
			}
		}
	}

	docs, total, err := dao.List(ctx, collectionName, opt)
	if err != nil {
		resp.Error(r, errors.DatabaseError(g.I18n().T(errmsg.ErrDBQuery)).WithOrigErr(err))
	} else {
		resp.PageOK(r, total, docs)
	}
}

// Get 详情查询
func (bizApi) Get(r *ghttp.Request) {
	var (
		localCtx types.Context
		ctxUser  types.ContextUser
	)
	ctx := r.Context()
	if err := ctxsrv.GetVar(ctx).Struct(&localCtx); err != nil {
		g.Log().Async().Error(gerror.Wrap(err, "error get user info from context"))
	} else if localCtx.User != nil {
		ctxUser = *localCtx.User
	}
	collectionName := r.GetString("collection")
	id, err := primitive.ObjectIDFromHex(r.GetString("id"))
	if err != nil {
		resp.Error(r, errors.NotFound(errmsg.ErrDBNotFound, g.I18n().T(errmsg.ErrDBNotFound)))
	}
	opt := &dao.GetOptions{
		Filter:  bson.M{"_id": id},
		CtxUser: ctxUser,
	}

	if !ctxUser.IsSysUser {
		perm := service.Permission.GetReadPermission(collectionName)
		if !perm.CanDoAll() {
			if perm.CanDoOnlyLogin() && ctxUser.ID == "" {
				resp.Error(r, errors.Unauthorized(errmsg.Unauthorized, g.I18n().T(errmsg.Unauthorized)))
			} else if perm.CanDoOnlyOwner() {
				opt.Filter = bson.M{"createdBy": ctxUser.ID}
			} else if perm.CanDoOnlySysUser() {
				resp.Error(r, errors.Forbidden(errmsg.PermissionDenied, g.I18n().T(errmsg.PermissionDenied)))
			}
		}
	}

	doc, err := dao.Get(ctx, collectionName, opt)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			resp.Error(r, errors.NotFound(errmsg.ErrDBNotFound, g.I18n().T(errmsg.ErrDBNotFound)))
		} else {
			resp.Error(r, errors.DatabaseError(g.I18n().T(errmsg.ErrDBGet)).WithOrigErr(err))
		}
	} else {
		resp.OK(r, doc)
	}
}

// Find 详情查询
func (bizApi) Find(r *ghttp.Request) {
	var (
		localCtx types.Context
		ctxUser  types.ContextUser
	)
	ctx := r.Context()
	if err := ctxsrv.GetVar(ctx).Struct(&localCtx); err != nil {
		g.Log().Async().Error(gerror.Wrap(err, "error get user info from context"))
	} else if localCtx.User != nil {
		ctxUser = *localCtx.User
	}
	collectionName := r.GetString("collection")
	filter := r.GetQueryMap()

	if !ctxUser.IsSysUser {
		perm := service.Permission.GetReadPermission(collectionName)
		if !perm.CanDoAll() {
			if perm.CanDoOnlyLogin() && ctxUser.ID == "" {
				resp.Error(r, errors.Unauthorized(errmsg.Unauthorized, g.I18n().T(errmsg.Unauthorized)))
			} else if perm.CanDoOnlyOwner() {
				filter["createdBy"] = ctxUser.ID
			} else if perm.CanDoOnlySysUser() {
				resp.Error(r, errors.Forbidden(errmsg.PermissionDenied, g.I18n().T(errmsg.PermissionDenied)))
			}
		}
	}

	opt := &dao.GetOptions{
		Filter:  filter,
		CtxUser: ctxUser,
	}
	doc, err := dao.Get(ctx, collectionName, opt)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			resp.Error(r, errors.NotFound(errmsg.ErrDBNotFound, g.I18n().T(errmsg.ErrDBNotFound)))
		} else {
			resp.Error(r, errors.DatabaseError(g.I18n().T(errmsg.ErrDBGet)).WithOrigErr(err))
		}
	} else {
		resp.OK(r, doc)
	}
}

// Create 新建数据
func (bizApi) Create(r *ghttp.Request) {
	var (
		localCtx types.Context
		ctxUser  types.ContextUser
	)
	ctx := r.Context()
	if err := ctxsrv.GetVar(ctx).Struct(&localCtx); err != nil {
		g.Log().Async().Error(gerror.Wrap(err, "error get user info from context"))
	} else if localCtx.User != nil {
		ctxUser = *localCtx.User
	}
	collectionName := r.GetString("collection")
	opt := &dao.InsertOptions{
		CtxUser: ctxUser,
	}

	if !ctxUser.IsSysUser {
		perm := service.Permission.GetCreatePermission(collectionName)
		if !perm.CanDoAll() {
			if perm.CanDoOnlyLogin() && ctxUser.ID == "" {
				resp.Error(r, errors.Unauthorized(errmsg.Unauthorized, g.I18n().T(errmsg.Unauthorized)))
			} else if perm.CanDoOnlySysUser() {
				resp.Error(r, errors.Forbidden(errmsg.PermissionDenied, g.I18n().T(errmsg.PermissionDenied)))
			}
		}
	}

	id, err := dao.Insert(ctx, collectionName, r.GetFormMap(), opt)
	if err != nil {
		switch err.(type) {
		case *gvalid.Error:
			resp.Error(r, errors.InvalidArgument(err.Error()))
		default:
			if mongo.IsDuplicateKeyError(err) {
				resp.Error(r, errors.AlreadyExists(g.I18n().T(errmsg.ErrDBDuplicate)))
			} else {
				resp.Error(r, errors.DatabaseError(g.I18n().T(errmsg.ErrDBInsert)).WithOrigErr(err))
			}
		}
	} else {
		resp.OK(r, g.Map{
			"_id": id,
		})
	}
}

// Update 更新数据
func (bizApi) Update(r *ghttp.Request) {
	var (
		localCtx types.Context
		ctxUser  types.ContextUser
	)
	ctx := r.Context()
	if err := ctxsrv.GetVar(ctx).Struct(&localCtx); err != nil {
		g.Log().Async().Error(gerror.Wrap(err, "error get user info from context"))
	} else if localCtx.User != nil {
		ctxUser = *localCtx.User
	}
	collectionName := r.GetString("collection")
	id, err := primitive.ObjectIDFromHex(r.GetString("id"))
	if err != nil {
		resp.Error(r, errors.NotFound(errmsg.ErrDBNotFound, g.I18n().T(errmsg.ErrDBNotFound)))
	}
	opt := &dao.UpdateOptions{
		Filter:  bson.M{"_id": id},
		CtxUser: ctxUser,
	}

	if !ctxUser.IsSysUser {
		perm := service.Permission.GetUpdatePermission(collectionName)
		if !perm.CanDoAll() {
			if perm.CanDoOnlyLogin() && ctxUser.ID == "" {
				resp.Error(r, errors.Unauthorized(errmsg.Unauthorized, g.I18n().T(errmsg.Unauthorized)))
			} else if perm.CanDoOnlyOwner() {
				opt.Filter = bson.M{"createdBy": ctxUser.ID}
			} else if perm.CanDoOnlySysUser() {
				resp.Error(r, errors.Forbidden(errmsg.PermissionDenied, g.I18n().T(errmsg.PermissionDenied)))
			}
		}
	}

	err = dao.Update(ctx, collectionName, r.GetFormMap(), opt)
	if err != nil {
		switch err.(type) {
		case *gvalid.Error:
			resp.Error(r, errors.InvalidArgument(err.Error()))
		default:
			if err == mongo.ErrNoDocuments {
				resp.Error(r, errors.NotFound(errmsg.ErrDBNotFound, g.I18n().T(errmsg.ErrDBNotFound)))
			} else if mongo.IsDuplicateKeyError(err) {
				resp.Error(r, errors.AlreadyExists(g.I18n().T(errmsg.ErrDBDuplicate)))
			} else {
				resp.Error(r, errors.DatabaseError(g.I18n().T(errmsg.ErrDBUpdate)).WithOrigErr(err))
			}
		}
	} else {
		resp.OK(r)
	}
}

// Delete 删除数据
func (bizApi) Delete(r *ghttp.Request) {
	var (
		localCtx types.Context
		ctxUser  types.ContextUser
	)
	ctx := r.Context()
	if err := ctxsrv.GetVar(ctx).Struct(&localCtx); err != nil {
		g.Log().Async().Error(gerror.Wrap(err, "error get user info from context"))
	} else if localCtx.User != nil {
		ctxUser = *localCtx.User
	}
	collectionName := r.GetString("collection")
	id, err := primitive.ObjectIDFromHex(r.GetString("id"))
	if err != nil {
		resp.Error(r, errors.NotFound(errmsg.ErrDBNotFound, g.I18n().T(errmsg.ErrDBNotFound)))
	}
	opt := &dao.DeleteOptions{
		Filter:  bson.M{"_id": id},
		CtxUser: ctxUser,
	}

	if !ctxUser.IsSysUser {
		perm := service.Permission.GetDeletePermission(collectionName)
		if !perm.CanDoAll() {
			if perm.CanDoOnlyLogin() && ctxUser.ID == "" {
				resp.Error(r, errors.Unauthorized(errmsg.Unauthorized, g.I18n().T(errmsg.Unauthorized)))
			} else if perm.CanDoOnlyOwner() {
				opt.Filter = bson.M{"createdBy": ctxUser.ID}
			} else if perm.CanDoOnlySysUser() {
				resp.Error(r, errors.Forbidden(errmsg.PermissionDenied, g.I18n().T(errmsg.PermissionDenied)))
			}
		}
	}

	err = dao.Delete(ctx, collectionName, opt)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			resp.Error(r, errors.NotFound(errmsg.ErrDBNotFound, g.I18n().T(errmsg.ErrDBNotFound)))
		} else {
			resp.Error(r, errors.DatabaseError(g.I18n().T(errmsg.ErrDBDelete)).WithOrigErr(err))
		}
	} else {
		resp.OK(r)
	}
}
