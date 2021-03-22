package controller

import (
	"antapi/api/resp"
	"antapi/core/crud"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
)

var Biz = new(bizControl)

type bizControl struct{}

type getListReq struct {
	Page int `d:"1"  v:"min:1#分页号码错误"`     // 分页号码
	Size int `d:"10" v:"max:50#分页数量最大50条"` // 分页数量，最大50
	Sort string
}

// Get : 查询详情
func (bizControl) Get(r *ghttp.Request) {
	collectionName := r.GetString("collection")
	id := r.GetString("id")

	if res, err := crud.Get(collectionName, g.Map{"id": id}); err != nil {
		resp.Error(r).SetError(err).Json()
	} else {
		resp.Success(r).SetData(res.MustToJsonString())
	}
}

// GetList : 查询列表数据
func (bizControl) GetList(r *ghttp.Request) {
	var reqArgs *getListReq
	collectionName := r.GetString("collection")
	if err := r.ParseQuery(&reqArgs); err != nil {
		resp.Error(r).SetError(err).Json()
	}

	if res, total, err := crud.GetList(collectionName, reqArgs.Page, reqArgs.Size, nil); err != nil {
		resp.Error(r).SetError(err).Json()
	} else {
		resp.Success(r).SetData(resp.ListsData{List: res.MustToJsonString(), Total: total}).Json()
	}
}

// Create : 新建数据
func (bizControl) Create(r *ghttp.Request) {
	collectionName := r.GetString("collection")

	if id, err := crud.Insert(collectionName, r.GetBodyString()); err != nil {
		resp.Error(r).SetError(err).Json()
	} else {
		resp.Success(r).SetData(g.Map{"id": id}).Json()
	}
}

// Update : 更新数据
func (bizControl) Update(r *ghttp.Request) {
	collectionName := r.GetString("collection")
	id := r.GetString("id")

	if err := crud.Update(collectionName, id, r.GetBodyString()); err != nil {
		resp.Error(r).SetError(err).Json()
	} else {
		resp.Success(r).Json()
	}
}

// Delete : 删除数据
func (bizControl) Delete(r *ghttp.Request) {
	collectionName := r.GetString("collection")
	id := r.GetString("id")

	if err := crud.Delete(collectionName, "id", gstr.SplitAndTrimSpace(id, ",")); err != nil {
		resp.Error(r).SetError(err).Json()
	} else {
		resp.Success(r).Json()
	}
}
