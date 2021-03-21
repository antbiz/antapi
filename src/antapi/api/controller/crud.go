package controller

import "github.com/gogf/gf/net/ghttp"

var Biz = new(bizControl)

type bizControl struct{}

type getListReq struct {
	Page int `d:"1"  v:"min:0#分页号码错误"`     // 分页号码
	Size int `d:"10" v:"max:50#分页数量最大50条"` // 分页数量，最大50
	Sort string
}

// Get : 查询详情
func (bizControl) Get(r *ghttp.Request) {

}

// GetList : 查询列表数据
func (bizControl) GetList(r *ghttp.Request) {

}

// Create : 新建数据
func (bizControl) Create(r *ghttp.Request) {

}

// Update : 更新数据
func (bizControl) Update(r *ghttp.Request) {

}

// Delete : 删除数据
func (bizControl) Delete(r *ghttp.Request) {

}
