package api

import (
	"antapi/app/dao"
	"antapi/common/errcode"
	"antapi/common/req"
	"antapi/common/resp"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
)

var Biz = new(bizApi)

type bizApi struct{}

// Get : 查询详情
func (bizApi) Get(r *ghttp.Request) {
	collectionName := r.GetString("collection")
	id := r.GetString("id")
	arg := &dao.GetFuncArg{
		Where:           "id",
		WhereArgs:       id,
		RaiseNotFound:   true,
		SessionUsername: req.GetSessionUsername(r),
	}

	if res, err := dao.Get(collectionName, arg); err != nil {
		resp.Error(r).SetError(gerror.Current(err)).SetCode(gerror.Code(err)).Json()
	} else {
		resp.Success(r).SetData(res.Map()).Json()
	}
}

// GetList : 查询列表数据
func (bizApi) GetList(r *ghttp.Request) {
	collectionName := r.GetString("collection")
	query, err := req.GetFilter(r)
	if err != nil {
		resp.Error(r).SetError(err).SetCode(errcode.ParameterBindError).Json()
	}
	arg := &dao.GetListFuncArg{
		Limit:           query.Limit,
		Offset:          query.Offset,
		Order:           query.GetOrderBy(),
		SessionUsername: req.GetSessionUsername(r),
	}

	if res, total, err := dao.GetList(collectionName, arg); err != nil {
		resp.Error(r).SetError(gerror.Current(err)).SetCode(gerror.Code(err)).Json()
	} else {
		if res == nil {
			resp.Success(r).SetData(resp.ListsData{List: g.SliceStr{}, Total: total}).Json()
		}
		resp.Success(r).SetData(resp.ListsData{List: res.Array(), Total: total}).Json()
	}
}

// Create : 新建数据
func (bizApi) Create(r *ghttp.Request) {
	collectionName := r.GetString("collection")
	arg := &dao.InsertFuncArg{
		SessionUsername: req.GetSessionUsername(r),
	}

	if id, err := dao.Insert(collectionName, arg, r.GetFormMap()); err != nil {
		resp.Error(r).SetError(gerror.Current(err)).SetCode(gerror.Code(err)).Json()
	} else {
		resp.Success(r).SetData(g.Map{"id": id}).Json()
	}
}

// Update : 更新数据
func (bizApi) Update(r *ghttp.Request) {
	collectionName := r.GetString("collection")
	id := r.GetString("id")
	arg := &dao.UpdateFuncArg{
		RaiseNotFound:   true,
		SessionUsername: req.GetSessionUsername(r),
	}

	if err := dao.Update(collectionName, arg, id, r.GetFormMap()); err != nil {
		resp.Error(r).SetError(gerror.Current(err)).SetCode(gerror.Code(err)).Json()
	} else {
		resp.Success(r).Json()
	}
}

// Delete : 删除数据
func (bizApi) Delete(r *ghttp.Request) {
	collectionName := r.GetString("collection")
	id := r.GetString("id")
	arg := &dao.DeleteFuncArg{
		Where:           "id",
		WhereArgs:       gstr.SplitAndTrimSpace(id, ","),
		SessionUsername: req.GetSessionUsername(r),
	}

	if err := dao.Delete(collectionName, arg); err != nil {
		resp.Error(r).SetError(gerror.Current(err)).SetCode(gerror.Code(err)).Json()
	} else {
		resp.Success(r).Json()
	}
}
