package router

import (
	"antapi/api/controller"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/api/business", func(group *ghttp.RouterGroup) {
		// 查询 {collection} 列表
		group.GET("/{collection}", controller.CRUD.GetList)
		// 获取 {collection} 详情
		group.GET("/{collection}/{id}", controller.CRUD.Get)
		// 添加 {collection}
		group.POST("/{collection}/{id}", controller.CRUD.Create)
		// 修改 {collection}
		group.PATCH("/{collection}/{id}", controller.CRUD.Update)
		// 删除 {collection}
		group.DELETE("/{collection}/{id}", controller.CRUD.Delete)
	})
}
