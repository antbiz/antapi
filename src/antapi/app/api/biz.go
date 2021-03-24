package api

import (
	"antapi/app/dao"
	"antapi/common/errcode"
	"antapi/common/resp"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
)

var Biz = new(bizApi)

type bizApi struct{}

type getListReq struct {
	Page int `d:"1"  v:"min:1#分页号码错误"`     // 分页号码
	Size int `d:"10" v:"max:50#分页数量最大50条"` // 分页数量，最大50
	Sort string
}

// Get : 查询详情
func (bizApi) Get(r *ghttp.Request) {
	collectionName := r.GetString("collection")
	id := r.GetString("id")
	arg := &dao.GetFuncArg{
		Where: g.Map{"id": id},
	}

	if res, err := dao.Get(collectionName, arg); err != nil {
		resp.Error(r).SetError(gerror.Current(err)).SetCode(gerror.Code(err)).Json()
	} else {
		resp.Success(r).SetData(res.MustToJsonString()).Json()
	}
}

// GetList : 查询列表数据
func (bizApi) GetList(r *ghttp.Request) {
	var reqArgs *getListReq
	collectionName := r.GetString("collection")
	if err := r.ParseQuery(&reqArgs); err != nil {
		resp.Error(r).SetError(err).SetCode(errcode.ParameterBindError).Json()
	}
	arg := &dao.GetListFuncArg{
		PageNum:  reqArgs.Page,
		PageSize: reqArgs.Size,
		Order:    reqArgs.Sort,
	}

	if res, total, err := dao.GetList(collectionName, arg); err != nil {
		resp.Error(r).SetError(gerror.Current(err)).SetCode(gerror.Code(err)).Json()
	} else {
		resp.Success(r).SetData(resp.ListsData{List: res.MustToJsonString(), Total: total}).Json()
	}
}

// Create : 新建数据
func (bizApi) Create(r *ghttp.Request) {
	collectionName := r.GetString("collection")

	if id, err := dao.Insert(collectionName, &dao.InsertFuncArg{}, r.GetBodyString()); err != nil {
		resp.Error(r).SetError(gerror.Current(err)).SetCode(gerror.Code(err)).Json()
	} else {
		resp.Success(r).SetData(g.Map{"id": id}).Json()
	}
}

// Update : 更新数据
func (bizApi) Update(r *ghttp.Request) {
	collectionName := r.GetString("collection")
	id := r.GetString("id")

	if err := dao.Update(collectionName, &dao.UpdateFuncArg{}, id, r.GetBodyString()); err != nil {
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
		Where:     "id",
		WhereArgs: gstr.SplitAndTrimSpace(id, ","),
	}

	if err := dao.Delete(collectionName, arg); err != nil {
		resp.Error(r).SetError(gerror.Current(err)).SetCode(gerror.Code(err)).Json()
	} else {
		resp.Success(r).Json()
	}
}
