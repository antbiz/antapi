package api

import (
	"github.com/BeanWei/apikit/ctxsrv"
	"github.com/BeanWei/apikit/errors"
	"github.com/BeanWei/apikit/resp"
	"github.com/antbiz/antapi/internal/app/dao"
	"github.com/antbiz/antapi/internal/app/service"
	"github.com/antbiz/antapi/internal/common/types"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"go.mongodb.org/mongo-driver/bson"
)

// Biz 公共接口管理
var Biz = bizApi{}

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
	}
	opt.Offset = opt.Limit * (r.GetInt64("pageNum", 1) - 1)

	if !ctxUser.IsSysUser {
		if perm, err := service.Permission.GetReadPermission(collectionName); err != nil {
			resp.Error(r, errors.UnknownError("err_permission_check"))
		} else if perm.CanDoOnlySysUser() {
			resp.Error(r, errors.Forbidden("permission_denied", "Permission Denied"))
		} else if perm.CanDoOnlyOwner() {
			opt.Filter = bson.M{"createdBy": ctxUser.ID}
		} else if perm.CanDoOnlyLogin() && ctxUser.ID == "" {
			resp.Error(r, errors.Unauthorized("unauthorized", "Please Login"))
		}
	}

	docs, total, err := dao.List(r.Context(), collectionName, opt)
	if err != nil {
		resp.Error(r, errors.DatabaseError(""))
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
	id := r.GetString("id")
	opt := &dao.GetOptions{
		Filter:  bson.M{"_id": id},
		CtxUser: ctxUser,
	}

	if !ctxUser.IsSysUser {
		if perm, err := service.Permission.GetReadPermission(collectionName); err != nil {
			resp.Error(r, errors.UnknownError("err_permission_check"))
		} else if perm.CanDoOnlySysUser() {
			resp.Error(r, errors.Forbidden("permission_denied", "Permission Denied"))
		} else if perm.CanDoOnlyOwner() {
			opt.Filter = bson.M{"createdBy": ctxUser.ID}
		} else if perm.CanDoOnlyLogin() && ctxUser.ID == "" {
			resp.Error(r, errors.Unauthorized("unauthorized", "Please Login"))
		}
	}

	doc, err := dao.Get(r.Context(), collectionName, opt)
	if err != nil {
		resp.Error(r, errors.DatabaseError(""))
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
		if perm, err := service.Permission.GetCreatePermission(collectionName); err != nil {
			resp.Error(r, errors.UnknownError("err_permission_check"))
		} else if perm.CanDoOnlySysUser() {
			resp.Error(r, errors.Forbidden("permission_denied", "Permission Denied"))
		} else if perm.CanDoOnlyLogin() && ctxUser.ID == "" {
			resp.Error(r, errors.Unauthorized("unauthorized", "Please Login"))
		}
	}

	id, err := dao.Insert(r.Context(), collectionName, r.GetFormMap(), opt)
	if err != nil {
		resp.Error(r, errors.DatabaseError(""))
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
	id := r.GetString("id")
	opt := &dao.UpdateOptions{
		Filter:  bson.M{"_id": id},
		CtxUser: ctxUser,
	}

	if !ctxUser.IsSysUser {
		if perm, err := service.Permission.GetUpdatePermission(collectionName); err != nil {
			resp.Error(r, errors.UnknownError("err_permission_check"))
		} else if perm.CanDoOnlySysUser() {
			resp.Error(r, errors.Forbidden("permission_denied", "Permission Denied"))
		} else if perm.CanDoOnlyOwner() {
			opt.Filter = bson.M{"createdBy": ctxUser.ID}
		} else if perm.CanDoOnlyLogin() && ctxUser.ID == "" {
			resp.Error(r, errors.Unauthorized("unauthorized", "Please Login"))
		}
	}

	err := dao.Update(r.Context(), collectionName, r.GetFormMap(), opt)
	if err != nil {
		resp.Error(r, errors.DatabaseError(""))
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
	id := r.GetString("id")
	opt := &dao.DeleteOptions{
		Filter:  bson.M{"_id": id},
		CtxUser: ctxUser,
	}

	if !ctxUser.IsSysUser {
		if perm, err := service.Permission.GetDeletePermission(collectionName); err != nil {
			resp.Error(r, errors.UnknownError("err_permission_check"))
		} else if perm.CanDoOnlySysUser() {
			resp.Error(r, errors.Forbidden("permission_denied", "Permission Denied"))
		} else if perm.CanDoOnlyOwner() {
			opt.Filter = bson.M{"createdBy": ctxUser.ID}
		} else if perm.CanDoOnlyLogin() && ctxUser.ID == "" {
			resp.Error(r, errors.Unauthorized("unauthorized", "Please Login"))
		}
	}

	err := dao.Delete(r.Context(), collectionName, opt)
	if err != nil {
		resp.Error(r, errors.DatabaseError(""))
	} else {
		resp.OK(r)
	}
}