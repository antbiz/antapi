package api

import (
	"github.com/BeanWei/apikit/errors"
	"github.com/BeanWei/apikit/resp"
	"github.com/antbiz/antapi/internal/app/dao"
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
	collectionName := r.GetString("collection")

	opt := &dao.ListOptions{
		Limit: r.GetInt64("pageSize"),
		Sort:  gstr.SplitAndTrimSpace(r.GetString("sort"), ","),
	}
	opt.Offset = opt.Limit * (r.GetInt64("pageNum", 1) - 1)

	docs, total, err := dao.List(r.Context(), collectionName, opt)
	if err != nil {
		resp.Error(r, errors.DatabaseError(""))
	} else {
		resp.PageOK(r, total, docs)
	}
}

// Get 详情查询
func (bizApi) Get(r *ghttp.Request) {
	collectionName := r.GetString("collection")
	id := r.GetString("id")

	opt := &dao.GetOptions{
		Filter: bson.M{"_id": id},
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
	collectionName := r.GetString("collection")

	id, err := dao.Insert(r.Context(), collectionName, r.GetFormMap())
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
	collectionName := r.GetString("collection")
	id := r.GetString("id")

	opt := &dao.UpdateOptions{
		Filter: bson.M{"_id": id},
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
	collectionName := r.GetString("collection")
	id := r.GetString("id")

	opt := &dao.DeleteOptions{
		Filter: bson.M{"_id": id},
	}

	err := dao.Delete(r.Context(), collectionName, opt)
	if err != nil {
		resp.Error(r, errors.DatabaseError(""))
	} else {
		resp.OK(r)
	}
}
